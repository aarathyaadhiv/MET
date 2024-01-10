package usecase

import (
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
)


type ChatUseCase struct{
	Repo interfaces.ChatRepository
}

func NewChatUseCase(repo interfaces.ChatRepository)useCaseInterface.ChatUseCase{
	return &ChatUseCase{repo}
}

