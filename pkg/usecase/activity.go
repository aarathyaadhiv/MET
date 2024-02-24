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
	isIdExist, err := l.Lik.IsUserExist(likedId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if !isIdExist {
		return response.Like{}, errors.New("you are trying to like a non existing userId")
	}
	if likedId == userId {
		return response.Like{}, errors.New("cannot like the user itself")
	}
	isExist, err := l.Lik.IsLikeExist(userId, likedId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if isExist {
		return response.Like{}, errors.New("liked already")
	}
	isSubscribed, err := l.Lik.IsSubscribed(userId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	var res response.Like
	if !isSubscribed {
		likes, err := l.Lik.LikeCount(userId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
		if likes == 0 {
			return response.Like{}, errors.New("your like limit exceeded")
		}
		res, err = l.Lik.Like(likedId, userId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
		err = l.Lik.UpdateLikeCount(userId, likes-1)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
	} else {
		res, err = l.Lik.Like(likedId, userId)
		if err != nil {
			return response.Like{}, errors.New("error in connecting database")
		}
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
	isIdExist, err := l.Lik.IsUserExist(likeId)
	if err != nil {
		return response.Like{}, errors.New("error in connecting database")
	}
	if !isIdExist {
		return response.Like{}, errors.New("you are trying to unlike a non existing userId")
	}
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
	isSubscribed, err := l.Lik.IsSubscribed(userId)
	if err != nil {
		return response.ShowLike{}, errors.New("error in checking subscribtion data from database")
	}
	seeLike := false
	if isSubscribed {
		seeLike, err = l.Lik.SeeLike(userId)
		if err != nil {
			return response.ShowLike{}, errors.New("error in checking seelike  from database")
		}

	}

	res, err := l.Lik.GetLike(page, count, userId)
	if err != nil {
		return response.ShowLike{}, errors.New("error in fetching  data from database")
	}
	updateResponse := make([]response.ShowUserDetails, 0)
	for _, r := range res {
		r.Interests, err = l.Lik.FetchInterests(r.Id)
		if err != nil {
			return response.ShowLike{}, errors.New("error in fetching interests from database")
		}
		updateResponse = append(updateResponse, r)
	}

	likeCount, err := l.Lik.GetLikeCount(userId)
	if err != nil {
		return response.ShowLike{}, errors.New("error in fetching count from database")
	}

	return response.ShowLike{
		UserId:       userId,
		IsSubscribed: isSubscribed,
		SeeLike:      seeLike,
		LikeCount:    likeCount,
		Likes:        updateResponse,
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
	updatedResponse := make([]response.ShowUserDetails, 0)
	for _, v := range res {
		v.Interests, err = l.Lik.FetchInterests(v.Id)
		if err != nil {
			return response.ShowMatch{}, errors.New("error in fetching interests")
		}
		updatedResponse = append(updatedResponse, v)
	}
	return response.ShowMatch{
		UserId:  userId,
		Matches: updatedResponse,
	}, nil
}

func (l *ActivityUseCase) Report(reportId, userId uint, message string) (response.Report, error) {
	isIdExist, err := l.Lik.IsUserExist(reportId)
	if err != nil {
		return response.Report{}, errors.New("error in connecting database")
	}
	if !isIdExist {
		return response.Report{}, errors.New("you are trying to report a non existing userId")
	}
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
	isIdExist, err := l.Lik.IsUserExist(blockedId)
	if err != nil {
		return response.BlockUser{}, errors.New("error in connecting database")
	}
	if !isIdExist {
		return response.BlockUser{}, errors.New("you are trying to block a non existing userId")
	}
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
