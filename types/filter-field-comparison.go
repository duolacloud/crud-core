package types

type BooleanFieldComparisons struct {
	Is bool
	IsNot bool
}

type CommonFieldComparisonBetweenType[FieldType any] struct {
	Lower FieldType
	Upper FieldType
}

type CommonFieldComparisonType[FieldType any] struct {
	BooleanFieldComparisons
	Eq FieldType
	Neq FieldType
	Gt FieldType
	Gte FieldType
	Lt FieldType
	Lte FieldType
	In []FieldType
	NotIn []FieldType
	Between CommonFieldComparisonBetweenType[FieldType]
	NotBetween CommonFieldComparisonBetweenType[FieldType]
}

type StringFieldComparisons struct {
	CommonFieldComparisonType[string]
	Like string
	NotLike string
	Ilike string
	NotILike string
}

type FilterFieldComparison [FieldType any] struct {
	BooleanFieldComparisons
	StringFieldComparisons
	CommonFieldComparisonType[FieldType]
}