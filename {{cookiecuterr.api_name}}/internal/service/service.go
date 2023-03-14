package service

import (
	"context"

	"example/internal/entity"
)

type Service interface {
	Call(ctx context.Context, req entity.Request) (*entity.Response, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (service) Call(ctx context.Context, req entity.Request) (*entity.Response, error) {
	// DO NOTHING
	return nil, nil
}
