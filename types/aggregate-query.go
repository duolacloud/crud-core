package types

type AggregateQuery struct {
	Count   []string
	Sum     []string
	Avg     []string
	Max     []string
	Min     []string
	GroupBy []string
}
