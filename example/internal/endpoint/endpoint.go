package endpoint

import (
	"context"
	"errors"
	"fmt"

	"example/internal/entity"
	"example/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeServiceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in any) (any, error) {
		req, ok := in.(entity.Request)
		if !ok {
			fmt.Println(ok, "no es ok")
			return nil, errors.New("Error")
		}

		msg, _ := svc.Hello(ctx, req.Body)

		return entity.Response{Message: msg}, nil
	}
}
