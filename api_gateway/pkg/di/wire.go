package di

import (
	"api-gateway/pkg/api"
	"api-gateway/pkg/api/handler"
	"api-gateway/pkg/client"
	"api-gateway/pkg/config"

	"github.com/google/wire"
)

func initializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(client.InitClient, client.NewVideoClient, handler.NewVideoHandler, api.NewServeHTTP)
	return &api.Server{}, nil
}