package main

import (
	"log"
	"net/http"

	"example/cmd/config"
	"example/internal/endpoint"
	"example/internal/service"
	"example/internal/transport"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.GetAPIConfig()
	if err != nil {
		log.Fatal(err)
	}

	svc := service.NewService()

	serviceHandler := httptransport.NewServer(
		endpoint.MakeServiceEndpoint(svc),
		transport.DecodeRequest,
		transport.EncodeResponse,
	)

	router := mux.NewRouter().PathPrefix(cfg.URIPrefix).Subrouter()

	router.Methods(http.MethodPost).Path("/").Handler(serviceHandler)

	server := &http.Server{
		Addr:              cfg.Port,
		ReadHeaderTimeout: cfg.Timeout,
		Handler:           router,
	}

	log.Println("ListenAndServe on localhost" + cfg.Port)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
