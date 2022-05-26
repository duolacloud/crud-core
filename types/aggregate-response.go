package types

type NumberAggregate = map[string]any

type TypeAggregate = map[string]any

type AggregateResponse struct {
	Count   NumberAggregate `json:"count,omitempty"`
	Sum     NumberAggregate `json:"sum,omitempty"`
	Avg     NumberAggregate `json:"avg,omitempty"`
	Max     TypeAggregate   `json:"max,omitempty"`
	Min     TypeAggregate   `json:"min,omitempty"`
	GroupBy map[string]any  `json:"group_by,omitempty"`
}

func (r *AggregateResponse) Append(aggFunc string, field string, value any) {
	var arr map[string]any

	switch aggFunc {
	case "count":
		if r.Count == nil {
			r.Count = map[string]any{}
		}
		arr = r.Count
	case "sum":
		if r.Sum == nil {
			r.Sum = map[string]any{}
		}
		arr = r.Sum
	case "avg":
		if r.Avg == nil {
			r.Avg = map[string]any{}
		}
		arr = r.Avg
	case "max":
		if r.Max == nil {
			r.Max = map[string]any{}
		}
		arr = r.Max
	case "min":
		if r.Min == nil {
			r.Min = map[string]any{}
		}
		arr = r.Min
	case "group_by":
		if r.GroupBy == nil {
			r.GroupBy = map[string]any{}
		}
		arr = r.GroupBy
	}

	arr[field] = value
}

func (r *AggregateResponse) Merge(o *AggregateResponse) {
	if o.Count != nil {
		if r.Count == nil {
			r.Count = map[string]any{}
		}

		for k, v := range o.Count {
			r.Count[k] = v
		}
	}

	if o.Sum != nil {
		if r.Sum == nil {
			r.Sum = map[string]any{}
		}

		for k, v := range o.Sum {
			r.Sum[k] = v
		}
	}
	if o.Avg != nil {
		if r.Avg == nil {
			r.Avg = map[string]any{}
		}

		for k, v := range o.Avg {
			r.Avg[k] = v
		}
	}

	if o.Max != nil {
		if r.Max == nil {
			r.Max = map[string]any{}
		}

		for k, v := range o.Max {
			r.Max[k] = v
		}
	}

	if o.Min != nil {
		if r.Min == nil {
			r.Min = map[string]any{}
		}

		for k, v := range o.Min {
			r.Min[k] = v
		}
	}

	if o.GroupBy != nil {
		if r.GroupBy == nil {
			r.GroupBy = map[string]any{}
		}

		for k, v := range o.GroupBy {
			r.GroupBy[k] = v
		}
	}
}
