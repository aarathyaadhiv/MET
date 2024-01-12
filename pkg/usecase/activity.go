package usecase

import (
	"errors"

	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type ActivityUseCase struct {
	Lik  interfaces.ActivityRepository
	Chat interfaces.ChatRepository
}

func NewActivityUseCase(like interfaces.ActivityRepository, chat interfaces.ChatRepository) useCaseInterface.ActivityUseCase {
	return &ActivityUseCase{Lik: like, Chat: chat}
}

func (l *ActivityUseCase) Like(likedId, userId uint) (response.Like, error) {
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
	exist, err := l.Lik.IsLikeExist(likedId, userId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if exist {
		err := l.Lik.Match(userId, likedId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
		err = l.Lik.Match(likedId, userId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
		
		err = l.Chat.CreateChatRoom(userId, likedId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
	}
	return res, nil
}

func (l *ActivityUseCase) Unlike(likeId, userId uint) (response.Like, error) {
	isExist, err := l.Lik.IsLikeExist(userId, likeId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if !isExist {
		return response.Like{}, errors.New("it is not liked previously")
	}
	match, err := l.Lik.IsMatchExist(userId, likeId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if match {
		return response.Like{}, errors.New("matched user , cannot unlike. please unmatch")
	}
	res, err := l.Lik.Unlike(likeId, userId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	return res, nil
}

func (l *ActivityUseCase) GetLike(page, count int, userId uint) (response.ShowLike, error) {
	res, err := l.Lik.GetLike(page, count, userId)
	if err != nil {
		return response.ShowLike{}, errors.New("error in fetching data from database")
	}

	return response.ShowLike{
		UserId: userId,
		Likes:  res,
	}, nil

}

func (l *ActivityUseCase) UnMatch(userId, matchId uint) (response.UnMatch, error) {
	match, err := l.Lik.IsMatchExist(userId, matchId)
	if err != nil {
		return response.UnMatch{}, errors.New("error in fetching data ")
	}
	if !match {
		return response.UnMatch{}, errors.New("cannot unmatch as it is not matched user")
	}
	res, err := l.Lik.UnMatch(userId, matchId)
	if err != nil {
		return response.UnMatch{}, errors.New("error in fetching data ")
	}
	return res, nil
}

func (l *ActivityUseCase) GetMatch(userId uint, page, count int) (response.ShowMatch, error) {
	res, err := l.Lik.GetMatch(page, count, userId)
	if err != nil {
		return response.ShowMatch{}, errors.New("error in fetching data")
	}

	return response.ShowMatch{
		UserId:  userId,
		Matches: res,
	}, nil
}

func (l *ActivityUseCase) Report(reportId, userId uint, message string) (response.Report, error) {
	isReported, err := l.Lik.IsReported(userId, reportId)
	if err != nil {
		return response.Report{}, errors.New("error in fetching data")
	}
	if isReported {
		return response.Report{}, errors.New("you have already reported the user")
	}
	res, err := l.Lik.Report(userId, reportId, message)
	if err != nil {
		return response.Report{}, errors.New("error in fetching data")
	}
	return res, nil
}

func (l *ActivityUseCase) BlockUser(userId, blockedId uint) (response.BlockUser, error) {
	isBlocked, err := l.Lik.IsBlocked(userId, blockedId)
	if err != nil {
		return response.BlockUser{}, errors.New("error in fetching data")
	}
	if isBlocked {
		return response.BlockUser{}, errors.New("you have already blocked the user")
	}
	res, err := l.Lik.BlockUser(userId, blockedId)
	if err != nil {
		return response.BlockUser{}, errors.New("error in fetching data")
	}
	_, err = l.Lik.BlockUser(blockedId, userId)
	if err != nil {
		return response.BlockUser{}, errors.New("error in fetching data")
	}
	return res, nil
}
