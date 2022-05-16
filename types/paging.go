package types

type SortField struct {
}

type PageQuery[T any] struct {
	Filter Filter[T]
	Page   int64
	Limit  int64
	Sort   []*SortField
}
