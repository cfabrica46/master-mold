package endpoint

import (
	"context"
	"errors"

	"example/internal/entity"
	"example/internal/service"

	"github.com/go-kit/kit/endpoint"
)

var errTypeAssertion = errors.New("error when convert type")

func MakeServiceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in any) (any, error) {
		req, ok := in.(entity.Request)
		if !ok {
			return nil, errTypeAssertion
		}

		return svc.Call(ctx, req)
	}
}
