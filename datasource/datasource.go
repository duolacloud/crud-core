package datasource

import(
	"context"
)

type DataSource[DB any] interface {
	GetDB(ctx ctx.Context) (*DB, error)
}

type dataSource[DB any] struct {
	db *DB
}

type NewSingleDatasource[DB any](db DB) DataSource[DB] {
	return &dataSource[DB]{
		db,
	}
}

func (s *dataSource[DB]) GetDB(ctx context.Context) (*DB, error) {
	return s.db, nil
}
