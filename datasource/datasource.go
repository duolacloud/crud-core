package datasource

import(
	"context"
)

type DataSource interface {
	GetDB(ctx ctx.Context) (any, error)
}

type dataSource struct {
	db any
}

type NewSingleDatasource(db any) DataSource {
	return &dataSource{
		db,
	}
}

func (s *dataSource) GetDB(ctx context.Context) (any, error) {
	return s.db, nil
}
