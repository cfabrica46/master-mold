package main

import (
	"net"
	"net/http"
	"os"

	commons "gitlab.falabella.tech/fif/integracion/forthehorde/commons/go-microservices-commons"
	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	otr "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	tr "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"{{cookiecutter.api_name}}/cmd/config"
	"{{cookiecutter.api_name}}/internal/endpoint"
	h "{{cookiecutter.api_name}}/internal/handler"
	"{{cookiecutter.api_name}}/internal/service"
)

func main() {
	cfg := config.GetAPIConfig()

	logger := commons.ConfigureLogger(cfg.LoggingLevel)

	var metricsConf *commons.MetricsConfig
	if cfg.EnabledMetrics {
		metricsConf = commons.MakeDefaultEndpointMetrics()
	}

	var tracer opentracing.Tracer

	if cfg.DDTraceEnabled {
		// New
		addr := net.JoinHostPort(
			os.Getenv("DD_AGENT_HOST"),
			"8126",
		)

		tracer = otr.New(
			tr.WithAgentAddr(addr),
		)

		defer tr.Stop()

	} else {
		tracer = opentracing.NoopTracer{}
	}

	opentracing.SetGlobalTracer(tracer)

	if cfg.DDProfileEnabled {
		// Mover a commons
		err := profiler.Start(
			profiler.WithAgentAddr(os.Getenv("DD_AGENT_HOST")+":8126"),
			profiler.WithProfileTypes(
				profiler.CPUProfile,
				profiler.HeapProfile,
			),
		)
		if err != nil {
			logger.Log(err)
			os.Exit(1)
		}
		defer profiler.Stop()
	}

	var handler http.Handler
	{
		svc := service.MakeService()
		serviceEndpoint := endpoint.MakeServiceEndpoint(svc)
		handler = h.NewHTTPHandler(logger, serviceEndpoint, tracer, metricsConf, cfg.URIPrefix)
	}

	g := commons.CreateServer(handler, cfg.Port, logger)

	logger.Log("exit", g.Run())
}
