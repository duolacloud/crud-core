package types

type ID interface {
	int | int32 | int64 | string | any
}