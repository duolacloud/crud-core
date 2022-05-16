package types

type CreateOptions struct {
}

type CreateOption func(*CreateOptions)

type UpdateOptions struct {
	Upsert bool
}

type UpdateOption func(*UpdateOptions)

func WithUpsert(v bool) UpdateOption {
	return func(o *UpdateOptions) {
		o.Upsert = v
	}
}
