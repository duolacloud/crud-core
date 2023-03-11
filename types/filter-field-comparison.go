package types

type FilterComparisonOperators string

const (
	FilterComparisonOperatorsEq    FilterComparisonOperators = "eq"
	FilterComparisonOperatorsNeq   FilterComparisonOperators = "neq"
	FilterComparisonOperatorsGt    FilterComparisonOperators = "gt"
	FilterComparisonOperatorsGte   FilterComparisonOperators = "gte"
	FilterComparisonOperatorsLt    FilterComparisonOperators = "lt"
	FilterComparisonOperatorsLte   FilterComparisonOperators = "lte"
	FilterComparisonOperatorsIn    FilterComparisonOperators = "in"
	FilterComparisonOperatorsNotIn FilterComparisonOperators = "notin"

	FilterComparisonOperatorsIs    FilterComparisonOperators = "is"
	FilterComparisonOperatorsIsNot FilterComparisonOperators = "isnot"
	FilterComparisonOperatorsLike  FilterComparisonOperators = "like"
	FilterComparisonOperatorsIlike FilterComparisonOperators = "iLike"
)
