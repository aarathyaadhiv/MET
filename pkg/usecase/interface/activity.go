package useCaseInterface

import "github.com/aarathyaadhiv/met/pkg/utils/response"


type ActivityUseCase interface{
	Like(likedId, userId uint) (response.Like, error)
	Unlike(likeId, userId uint) (response.Like, error)
	GetLike(page, count int, userId uint) (response.ShowLike, error) 
	UnMatch(userId,matchId uint)(response.UnMatch,error)
	GetMatch(userId uint,page,count int)(response.ShowMatch,error)
	Report(reportId,userId uint,message string)(response.Report,error)
	BlockUser(userId,blockedId uint)(response.BlockUser,error)
}