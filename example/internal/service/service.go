package service

import "context"

type Service interface {
	Hello(ctx context.Context, message string) (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (service) Hello(_ context.Context, message string) (string, error) {
	return message, nil
}
