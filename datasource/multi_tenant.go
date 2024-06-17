package datasource

import(
	"context"
)

type DBGetter[DB any] interface {
	Get(ctx context.Context, tenantId string) (*DB, error)
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
	tenantId, ok := ctx.Value(s.tenantKey).(string)
	if !ok {
		return nil, fmt.Errorf("invalid value, value for tenant key not string", key)
	}

	db, err := s.dbGetter.Get(ctx, tenantId)
	if err != nil {
		return nil, err
	}
}
