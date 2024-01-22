package interfaces

import "github.com/aarathyaadhiv/met/pkg/utils/response"



type ActivityRepository interface{
	Like(likedId,userId uint)(response.Like,error)
	Unlike(likeId,userId uint)(response.Like,error)
	GetLike(page,count int,userId uint)([]response.ShowUserDetails,error)
	IsLikeExist(userId,likedId uint)(bool,error)
	LikeCount(userId uint)(int,error)
	UpdateLikeCount(userId uint,count int)error
	IsSubscribed(userId uint)(bool,error)
	SeeLike(userId uint)(bool,error)
	Match(userId,matchId uint)error
	GetMatch(page,count int,userId uint)([]response.ShowUserDetails,error)
	UnMatch(userId,matchId uint)(response.UnMatch,error)
	IsMatchExist(userId,matchId uint)(bool,error)
	IsReported(userId,reportId uint)(bool,error)
	Report(userId,reportId uint,message string) (response.Report, error) 
	IsBlocked(userId,blockedId uint)(bool,error)
	BlockUser(userId, blockedId uint) (response.BlockUser, error)
}