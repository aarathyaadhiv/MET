package response

import "github.com/aarathyaadhiv/met/pkg/utils/models"



type ChatResponse struct{
	Chat models.Chat
	User models.UserShortDetail
}