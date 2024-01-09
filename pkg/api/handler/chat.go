package handler

import (
	handlerInterface "github.com/aarathyaadhiv/met/pkg/api/handler/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
)


type ChatHandler struct{
	UseCase useCaseInterface.ChatUseCase
}

func NewChatHandler(usecase useCaseInterface.ChatUseCase)handlerInterface.ChatHandler{
	return &ChatHandler{usecase}
}