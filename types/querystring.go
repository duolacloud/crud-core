package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	bracketRE = regexp.MustCompile(`(?P<typ>filter|sort|page)=(?P<field>.+?)(\&|\z)`)
	commaRE   = regexp.MustCompile(`\s?\,\s?`)
	fieldsRE  = regexp.MustCompile(`fields=(?P<field>.+?)(\&|\z)`)
	sortRE    = regexp.MustCompile(`sort=(?P<field>.+?)(\&|\z)`)
	valueRE   = regexp.MustCompile(`\=(.*?)(\&|\z)`)
)

// FromQuerystring parses an Options object from the provided querystring
func FromQuerystring(qs string) (*PageQuery, error) {
	if qs == "" {
		return &PageQuery{}, nil
	}

	uqs, err := url.QueryUnescape(qs)
	if err != nil {
		return &PageQuery{}, err
	}

	pageQuery := &PageQuery{}

	// parse fields
	pageQuery.Fields = parseFields(&uqs)

	// parse sort
	pageQuery.Sort = parseSort(&uqs)

	// parse filter and page
	if err := parseBracketParams(uqs, pageQuery); err != nil {
		return pageQuery, err
	}

	// attempt to infer pagination strategy
	if _, ok := pageQuery.Page["limit"]; ok {
		pageQuery.SetPaginationStrategy(&OffsetStrategy{})
	}

	if _, ok := pageQuery.Page["size"]; ok {
		pageQuery.SetPaginationStrategy(&PageSizeStrategy{})
	}

	return pageQuery, nil
}

func extract(qs *string, re regexp.Regexp) string {
	coords := re.FindStringIndex(*qs)

	var r string
	if coords != nil {
		if coords[0] == 0 {
			// at beginning of string
			r = (*qs)[coords[0]:coords[1]]
			*qs = (*qs)[coords[1]:]
			return r
		}

		if coords[1] == len(*qs) {
			// at end of string
			r = (*qs)[coords[0]:coords[1]]
			*qs = (*qs)[0:coords[0]]
			return r
		}

		// in middle of string
		r = (*qs)[coords[0]:coords[1]]
		*qs = fmt.Sprintf("%s%s", (*qs)[0:coords[0]], (*qs)[coords[1]:])
		return r
	}

	return r
}

func parseBracketParams(qs string, o *PageQuery) error {
	o.Filter = map[string]any{}
	o.Page = map[string]int{}

	terms := bracketRE.FindAllStringSubmatch(qs, -1)
	values := valueRE.FindAllStringSubmatch(qs, -1)

	if len(terms) > 0 && len(terms) > len(values) {
		// multiple nested bracket params... not sure how to parse
		return errors.New("unable to parse: an object hierarchy has been provided")
	}

	for i, term := range terms {
		switch strings.ToLower(term[1]) {
		case "filter":
			if o.Filter == nil {
				o.Filter = map[string]any{}
			}

			filter := make(map[string]any)
			err := json.Unmarshal([]byte(values[i][1]), &filter)
			if err != nil {
				fmt.Printf("Unmarshal with error: %+v\n", err)
				return err
			}

			fmt.Printf("values: %d, %v\n", i, filter)

			o.Filter = filter
		case "page":
			if o.Page == nil {
				o.Page = map[string]int{}
			}

			page := make(map[string]int)
			err := json.Unmarshal([]byte(values[i][1]), &page)
			if err != nil {
				fmt.Printf("Unmarshal with error: %+v\n", err)
				return err
			}

			o.Page = page
		}
	}

	return nil
}

func parseFields(qs *string) []string {
	fields := []string{}

	fieldNames := fieldsRE.FindAllStringSubmatch(extract(qs, *fieldsRE), -1)
	for fieldNames != nil {
		for _, field := range fieldNames {
			// check if sort value is an array
			if commaRE.MatchString(field[1]) {
				fa := commaRE.Split(field[1], -1)
				fields = append(fields, fa...)
				continue
			}

			fields = append(fields, field[1])
		}

		// look for more fields= occurrences
		fieldNames = fieldsRE.FindAllStringSubmatch(extract(qs, *fieldsRE), -1)
	}

	return fields
}

func parseSort(qs *string) []string {
	sort := []string{}

	fieldNames := sortRE.FindAllStringSubmatch(extract(qs, *sortRE), -1)
	for fieldNames != nil {
		for _, field := range fieldNames {
			// check if sort value is an array
			if commaRE.MatchString(field[1]) {
				fa := commaRE.Split(field[1], -1)
				sort = append(sort, fa...)
				continue
			}

			sort = append(sort, field[1])
		}

		// look for more sort= occurrences
		fieldNames = sortRE.FindAllStringSubmatch(extract(qs, *sortRE), -1)
	}

	return sort
}
