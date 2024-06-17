package datasource

type multiTenantDataSource struct {
	dbRouter,
}

func NewMultiTenantDataSource(key string) DataSource {
	return &multiTenantDataSource{
		dbRouter,
	}
}

func (s *multiTenantDataSource) GetDB(ctx context.Context) (any, error) {
	db := ctx.Value(s.key)
	if db == nil {
		return nil, fmt.Errorf("db not found for %v", key)
	}
}
