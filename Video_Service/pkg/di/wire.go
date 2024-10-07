package di

import (
	"video-microservice/pkg/api"
	"video-microservice/pkg/api/service"
	"video-microservice/pkg/config"
	"video-microservice/pkg/db"
	"video-microservice/pkg/repository"

	"github.com/google/wire"
)

func initializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb, repository.NewVideoRepo, service.NewVideoServer, api.NewgrpcServe)
	return &api.Server{}, nil
}
