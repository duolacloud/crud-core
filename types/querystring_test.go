package types

import (
	"testing"
)

func TestQueryString(t *testing.T) {
	q, err := FromQuerystring("filter={\"name\":\"marry\", \"and\": [{\"age\": { \"eq\": 13}}]}&fields=name,age&page={\"offset\":0, \"limit\": 10}")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("q: filter: %v, page: %v\n", q.Filter, q.Page)
}
