package types

type NumberAggregate = map[string]interface{}

type TypeAggregate = map[string]interface{}

type AggregateResponse struct {
	Count   NumberAggregate        `json:"count,omitempty"`
	Sum     NumberAggregate        `json:"sum,omitempty"`
	Avg     NumberAggregate        `json:"avg,omitempty"`
	Max     TypeAggregate          `json:"max,omitempty"`
	Min     TypeAggregate          `json:"min,omitempty"`
	GroupBy map[string]interface{} `json:"group_by,omitempty"`
}

func (r *AggregateResponse) Append(aggFunc string, field string, value interface{}) {
	var arr map[string]interface{}

	switch aggFunc {
	case "count":
		if r.Count == nil {
			r.Count = map[string]interface{}{}
		}
		arr = r.Count
	case "sum":
		if r.Sum == nil {
			r.Sum = map[string]interface{}{}
		}
		arr = r.Sum
	case "avg":
		if r.Avg == nil {
			r.Avg = map[string]interface{}{}
		}
		arr = r.Avg
	case "max":
		if r.Max == nil {
			r.Max = map[string]interface{}{}
		}
		arr = r.Max
	case "min":
		if r.Min == nil {
			r.Min = map[string]interface{}{}
		}
		arr = r.Min
	case "group_by":
		if r.GroupBy == nil {
			r.GroupBy = map[string]interface{}{}
		}
		arr = r.GroupBy
	}

	arr[field] = value
}

func (r *AggregateResponse) Merge(o *AggregateResponse) {
	if o.Count != nil {
		if r.Count == nil {
			r.Count = map[string]interface{}{}
		}

		for k, v := range o.Count {
			r.Count[k] = v
		}
	}

	if o.Sum != nil {
		if r.Sum == nil {
			r.Sum = map[string]interface{}{}
		}

		for k, v := range o.Sum {
			r.Sum[k] = v
		}
	}
	if o.Avg != nil {
		if r.Avg == nil {
			r.Avg = map[string]interface{}{}
		}

		for k, v := range o.Avg {
			r.Avg[k] = v
		}
	}

	if o.Max != nil {
		if r.Max == nil {
			r.Max = map[string]interface{}{}
		}

		for k, v := range o.Max {
			r.Max[k] = v
		}
	}

	if o.Min != nil {
		if r.Min == nil {
			r.Min = map[string]interface{}{}
		}

		for k, v := range o.Min {
			r.Min[k] = v
		}
	}

	if o.GroupBy != nil {
		if r.GroupBy == nil {
			r.GroupBy = map[string]interface{}{}
		}

		for k, v := range o.GroupBy {
			r.GroupBy[k] = v
		}
	}
}
