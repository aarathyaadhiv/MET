//go:build wireinject
// +build wireinject

package di

import (
	server "github.com/aarathyaadhiv/met/pkg/api"
	"github.com/aarathyaadhiv/met/pkg/api/handler"
	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/db"
	"github.com/aarathyaadhiv/met/pkg/repository"
	"github.com/aarathyaadhiv/met/pkg/usecase"
	"github.com/google/wire"
)


func InitializeAPI(c config.Config)(*server.ServerHTTP,error){
	wire.Build(
		db.ConnectDB,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		server.NewServerHTTP,
	)
	return &server.ServerHTTP{},nil
}