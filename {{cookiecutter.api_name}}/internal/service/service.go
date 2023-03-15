package service

import (
	"context"

	"{{cookiecutter.api_name}}/internal/entity"
)

type Service interface {
	Call(context.Context, *entity.RequestStruct) (*entity.ResponseStruct, error)
}

type service struct {
}

func MakeService() Service {
	return new(service)
}

func (s *service) Call(context.Context, *entity.RequestStruct) (*entity.ResponseStruct, error){
	//DO NOTHING
	return new(entity.ResponseStruct), nil
}