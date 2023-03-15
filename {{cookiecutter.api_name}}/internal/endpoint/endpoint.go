package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"{{cookiecutter.api_name}}/internal/entity"
	"{{cookiecutter.api_name}}/internal/service"
)

//MakeServiceEndpoint crea el endpoint del service
func MakeServiceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {

		req := in.(*entity.RequestStruct)

		return svc.Call(ctx, req)
	}
}
