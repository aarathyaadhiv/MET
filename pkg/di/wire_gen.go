// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/aarathyaadhiv/met/pkg/api"
	"github.com/aarathyaadhiv/met/pkg/api/handler"
	"github.com/aarathyaadhiv/met/pkg/api/middleware"
	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/db"
	"github.com/aarathyaadhiv/met/pkg/repository"
	"github.com/aarathyaadhiv/met/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(c config.Config) (*server.ServerHTTP, error) {
	gormDB, err := db.ConnectDB(c)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, c)
	userHandler := handler.NewUserHandler(userUseCase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	authMiddleware := middleware.NewAuthMiddleware(userRepository)
	activityRepository := repository.NewActivityRepository(gormDB)
	database, err := db.ConnectMongo(c)
	if err != nil {
		return nil, err
	}
	chatRepository := repository.NewChatRepository(database)
	activityUseCase := usecase.NewActivityUseCase(activityRepository, chatRepository)
	activityHandler := handler.NewActivityHandler(activityUseCase)
	homeRepository := repository.NewHomeRepository(gormDB)
	homeUseCase := usecase.NewHomeUseCase(homeRepository)
	homeHandler := handler.NewHomeHandler(homeUseCase)
	chatUseCase := usecase.NewChatUseCase(chatRepository, userRepository)
	chatHandler := handler.NewChatHandler(chatUseCase)
	serverHTTP := server.NewServerHTTP(userHandler, adminHandler, authMiddleware, activityHandler, homeHandler, chatHandler)
	return serverHTTP, nil
}
