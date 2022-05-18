package types

type SortField struct {
}

type PageQuery[T any] struct {
	Fields []string
	Filter map[string]interface{}
	Offset   int64
	Limit  int64
	Sort   []string
}
