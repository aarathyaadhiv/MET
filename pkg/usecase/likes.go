package usecase

import (
	"errors"

	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type LikeUseCase struct {
	Lik  interfaces.LikeRepository
	User interfaces.UserRepository
}

func NewLikeUseCase(like interfaces.LikeRepository, user interfaces.UserRepository) useCaseInterface.LikeUseCase {
	return &LikeUseCase{Lik: like,
		User: user}
}

func (l *LikeUseCase) Like(likedId, userId uint) (response.Like, error) {
	isExist, err := l.Lik.IsLikeExist(userId, likedId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if isExist {
		return response.Like{}, errors.New("liked already")
	}
	res, err := l.Lik.Like(likedId, userId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	exist,err:=l.Lik.IsLikeExist(likedId,userId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if exist{
		err:=l.Lik.Match(userId,likedId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
		err=l.Lik.Match(likedId,userId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
	}
	return res, nil
}

func (l *LikeUseCase) Unlike(likeId, userId uint) (response.Like, error) {
	isExist, err := l.Lik.IsLikeExist(userId, likeId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if !isExist {
		return response.Like{}, errors.New("it is not liked previously")
	}
	match,err:=l.Lik.IsMatchExist(userId,likeId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if match{
		return response.Like{},errors.New("matched user , cannot unlike. please unmatch")
	}
	res, err := l.Lik.Unlike(likeId, userId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	return res, nil
}

func (l *LikeUseCase) GetLike(page, count int, userId uint) (response.ShowLike, error) {
	resId, err := l.Lik.GetLike(page, count, userId)
	if err != nil {
		return response.ShowLike{}, errors.New("error in fetching data from database")
	}
	likes := make([]response.ShowProfile, 0)
	for _, id := range resId {
		var like response.ShowProfile
		userDetails, err := l.User.ShowProfile(id)
		if err != nil {
			return response.ShowLike{}, errors.New("error in fetching data from database")
		}
		images, err := l.User.FetchImages(id)
		if err != nil {
			return response.ShowLike{}, errors.New("error in fetching data from database")
		}
		like.UserDetails = userDetails
		like.Image = images
		likes = append(likes, like)
	}
	return response.ShowLike{
		UserId: userId,
		Likes:  likes,
	}, nil

}

func (l *LikeUseCase) UnMatch(userId,matchId uint)(response.UnMatch,error){
	match,err:=l.Lik.IsMatchExist(userId,matchId)
	if err!=nil{
		return response.UnMatch{},errors.New("error in fetching data ")
	}
	if !match{
		return response.UnMatch{},errors.New("cannot unmatch as it is not matched user")
	}
	res,err:=l.Lik.UnMatch(userId,matchId)
	if err!=nil{
		return response.UnMatch{},errors.New("error in fetching data ")
	}
	return res,nil
}


func (l *LikeUseCase) GetMatch(userId uint,page,count int)(response.ShowMatch,error){
	resId,err:=l.Lik.GetMatch(page,count,userId)
	if err!=nil{
		return response.ShowMatch{},errors.New("error in fetching data")
	}
	matches := make([]response.ShowProfile, 0)
	for _, id := range resId {
		var match response.ShowProfile
		userDetails, err := l.User.ShowProfile(id)
		if err != nil {
			return response.ShowMatch{}, errors.New("error in fetching data from database")
		}
		images, err := l.User.FetchImages(id)
		if err != nil {
			return response.ShowMatch{}, errors.New("error in fetching data from database")
		}
		match.UserDetails = userDetails
		match.Image = images
		matches = append(matches, match)
	}
	return response.ShowMatch{
		UserId: userId,
		Matches: matches,
	},nil
}

