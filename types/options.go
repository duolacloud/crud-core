package types

type GetOptions struct {
}

type GetOption func(*GetOption)

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

type DeleteMode int
const (
	DeleteModeHard DeleteMode = iota
	DeleteModeSoft DeleteMode = iota
)

type DeleteOptions struct {
	DeleteMode DeleteMode
}

type DeleteOption func(*DeleteOptions)

func WithDeleteMode(v DeleteMode) DeleteOption {
	return func(o *DeleteOptions) {
		o.DeleteMode = v
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
