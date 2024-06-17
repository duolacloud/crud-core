package datasource

import(
	"context"
)

type DBGetter[DB any] interface {
	Get(ctx context.Context, tenantKey string) (*DB, error)
}

type multiTenantDataSource[DB any] struct {
	tenantKey string
	idbGetter DBGetter
}

func NewMultiTenantDataSource[DB any](tenantKey string, dbGetter DBGetter) DataSource[DB] {
	return &multiTenantDataSource[DB]{
		tenantKey,
		dbGetter,
	}
}

func (s *multiTenantDataSource[DB]) GetDB(ctx context.Context) (*DB, error) {
	dbKey := ctx.Value(s.key)
	if db == nil {
		return nil, fmt.Errorf("db not found for %v", key)
	}

	return s.dbGetter.Get(ctx, tenantKey)
}
