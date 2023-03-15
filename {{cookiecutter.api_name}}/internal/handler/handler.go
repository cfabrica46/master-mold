package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"{{cookiecutter.api_name}}/internal/entity"
)

func DecodeRequest(_ context.Context, r *http.Request) (any, error) {
	var request entity.Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(response)
}
