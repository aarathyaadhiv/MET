//go:build wireinject
// +build wireinject

package di

import (
	server "github.com/aarathyaadhiv/met/pkg/api"
	"github.com/aarathyaadhiv/met/pkg/api/handler"
	"github.com/aarathyaadhiv/met/pkg/api/middleware"
	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/db"
	"github.com/aarathyaadhiv/met/pkg/repository"
	"github.com/aarathyaadhiv/met/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(c config.Config) (*server.ServerHTTP, error) {
	wire.Build(
		db.ConnectDB,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
		repository.NewActivityRepository,
		usecase.NewActivityUseCase,
		handler.NewActivityHandler,
		repository.NewHomeRepository,
		usecase.NewHomeUseCase,
		handler.NewHomeHandler,
		db.ConnectMongo,
		repository.NewChatRepository,
		usecase.NewChatUseCase,
		handler.NewChatHandler,
		repository.NewSubscriptionRepository,
		usecase.NewSubscriptionUseCase,
		handler.NewSubscriptionHandler,
		middleware.NewAuthMiddleware,
		server.NewServerHTTP,
	)
	return &server.ServerHTTP{}, nil
}
