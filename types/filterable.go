package types

type Filter [T any] struct {
	FilterComparisons
	FilterGrouping[T]
}

type FilterGrouping [T any] struct {
	And []Filter[T]
	Or []Filter[T]
}

type FilterComparisons map[string]FilterFieldComparison[any]

type Filterable[T any] struct {
	Filter Filter[T]
}