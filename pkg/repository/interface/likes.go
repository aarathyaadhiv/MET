package interfaces

import "github.com/aarathyaadhiv/met/pkg/utils/response"



type LikeRepository interface{
	Like(likedId,userId uint)(response.Like,error)
	Unlike(likeId,userId uint)(response.Like,error)
	GetLike(page,count int,userId uint)([]uint,error)
	IsLikeExist(userId,likedId uint)(bool,error)
	Match(userId,matchId uint)error
	GetMatch(page,count int,userId uint)([]uint,error)
	UnMatch(userId,matchId uint)(response.UnMatch,error)
	IsMatchExist(userId,matchId uint)(bool,error)
}