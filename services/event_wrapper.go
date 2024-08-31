package services

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/duolacloud/broker-core"
	"github.com/duolacloud/crud-core/types"
)

type Options struct {
	prefix string
}

type Option func(*Options)


type ID interface {
	ID() string
}

func WithPrefix(v string) Option {
	return func(o *Options) {
		o.prefix = v
	}
}

type event[DTO ID, CreateDTO any, UpdateDTO any] struct {
	broker broker.Broker
	CrudService[DTO, CreateDTO, UpdateDTO]
	domain string
	opts *Options
}

func WrapEvent[DTO ID, CreateDTO any, UpdateDTO any](
	svc CrudService[DTO, CreateDTO, UpdateDTO],
	broker broker.Broker,
	o ...Option,
) CrudService[DTO, CreateDTO, UpdateDTO] {
	opts := &Options{}
	for _, opt := range o {
		opt(opts)
	}

	var m DTO
	appType := reflect.TypeOf(m)
	domain := strings.ToLower(appType.Name())

	return &event[DTO, CreateDTO, UpdateDTO]{
		CrudService: svc,
		broker:      broker,
		domain:     domain,
		opts:	    opts,
	}
}

func (s *event[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, dto *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	newDto, err := s.CrudService.Create(c, dto, opts...)
	if err != nil {
		return nil, err
	}

	err = publish(c, s.broker, domainCreated(s.opts.prefix, s.domain), newDto)
	if err != nil {
		return nil, err
	}

	return newDto, err
}

func (s *event[DTO, CreateDTO, UpdateDTO]) CreateMany(c context.Context, items []*CreateDTO, opts ...types.CreateManyOption) ([]*DTO, error) {
	dtos, err := s.CrudService.CreateMany(c, items, opts...)
	if err != nil {
		return nil, err
	}

	for _, dto := range dtos {
		err = publish(c, s.broker, domainCreated(s.opts.prefix, s.domain), dto)
		if err != nil {
			return nil, err
		}
	}

	return dtos, err
}

func (s *event[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	dto, err := s.CrudService.Get(c, id)
	if err != nil {
		return err
	}

	err = s.CrudService.Delete(c, id)
	if err != nil {
		return err
	}

	err = publish(c, s.broker, domainDeleted(s.opts.prefix, s.domain), dto)
	if err != nil {
		return err
	}

	return nil
}

func (s *event[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	dto, err := s.CrudService.Update(c, id, updateDTO, opts...)
	if err != nil {
		return nil, err
	}

	err = publish(c, s.broker, domainUpdated(s.opts.prefix, s.domain), dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (s *event[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID) (*DTO, error) {
	return s.CrudService.Get(c, id)
}

func (s *event[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	return s.CrudService.Query(c, query)
}

func (s *event[DTO, CreateDTO, UpdateDTO]) QueryOne(c context.Context, filter map[string]any) (*DTO, error) {
	return s.CrudService.QueryOne(c, filter)
}

func (s *event[DTO, CreateDTO, UpdateDTO]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	return s.CrudService.Count(c, query)
}

func (s *event[DTO, CreateDTO, UpdateDTO]) Aggregate(
	c context.Context,
	filter map[string]any,
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return s.CrudService.Aggregate(c, filter, aggregateQuery)
}

func publish[DTO ID](c context.Context, b broker.Broker, topic string, data *DTO) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return b.Publish(c, topic, &broker.Message{
		Body: body,
	}, broker.WithShardingKey((*data).ID()))
}

func domainCreated(prefix, domain string) string {
	if prefix != "" {
		prefix += "_"
	}
	return fmt.Sprintf("%s%s_created", prefix, domain)
}

func domainDeleted(prefix, domain string) string {
	if prefix != "" {
		prefix += "_"
	}
	return fmt.Sprintf("%s%s_deleted", domain)
}

func domainUpdated(prefix, domain string) string {
	if prefix != "" {
		prefix += "_"
	}
	return fmt.Sprintf("%s%s_updated", domain)
}

