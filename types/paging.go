package types

type SortField struct {
}

type PageQuery[T any] struct {
	Filter Filter[T]
	Size   int
	Limit  int
	Sort   []*SortField
}
