package handler

import (
	"context"
	"{{cookiecutter.api_name}}/internal/entity"
	"encoding/json"
	"net/http"

	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	commons "gitlab.falabella.tech/fif/integracion/forthehorde/commons/go-microservices-commons"
)

// NewHTTPHandler crea un nuevo http handler para un endpoint de service
func NewHTTPHandler(logger log.Logger, serviceEndpoint endpoint.Endpoint, otTracer stdopentracing.Tracer, metricsConf *commons.MetricsConfig, prefix string) http.Handler {
	return NewHTTPHandlerBuilder(logger, serviceEndpoint, otTracer, metricsConf, prefix).Build()
}

// NewHTTPHandlerBuilder crea un nuevo http handler builder
func NewHTTPHandlerBuilder(logger log.Logger, serviceEndpoint endpoint.Endpoint, otTracer stdopentracing.Tracer, metricsConf *commons.MetricsConfig, prefix string) commons.HTTPHandlerBuilder {
	endpointsCfg := []commons.EndpointConfig{
		commons.POST(prefix+"/test/hello", "service-service", serviceEndpoint, DecodeRequest, nil),
	}

	builder := commons.MakeHTTPHandlerBuilder(logger, endpointsCfg).WithTracer(otTracer).WithMetrics(metricsConf).WithHeaderLogger(commons.MakePassThroughAllHeadersMapFunc())
	return builder
}

// DecodeRequest request decode
func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var appRequest entity.RequestStruct

	if err := json.NewDecoder(r.Body).Decode(&appRequest); err != nil {
		return nil, err
	}

	return &appRequest, nil
}
