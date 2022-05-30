package types

type CreateOptions struct {
}

type CreateOption func(*CreateOptions)

type CreateManyOptions struct {
	CreateBatchSize int
}

type CreateManyOption func(*CreateManyOptions)

func WithCreateBatchSize(v int) CreateManyOption {
	return func(o *CreateManyOptions) {
		o.CreateBatchSize = v
	}
}

type UpdateOptions struct {
	Upsert bool
}

type UpdateOption func(*UpdateOptions)

func WithUpsert(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Upsert = v
	}
}
