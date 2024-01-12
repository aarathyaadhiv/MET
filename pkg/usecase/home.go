package usecase

import (
	"errors"

	"github.com/aarathyaadhiv/met/pkg/helper"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type HomeUseCase struct {
	Repo interfaces.HomeRepository
}

func NewHomeUseCase(repo interfaces.HomeRepository) useCaseInterface.HomeUseCase {
	return &HomeUseCase{repo}
}

func (h *HomeUseCase) HomePage(id uint, page, count int) ([]response.Home, error) {
	preference, err := h.Repo.FetchPreference(id)
	if err != nil {
		return nil, errors.New("error in fetching preference")
	}
	user, err := h.Repo.FetchUser(id)
	if err != nil {
		return nil, errors.New("error in fetching user data")
	}
	userInterests, err := h.Repo.FetchInterests(id)
	if err != nil {
		return nil, errors.New("error in fetching user interests")
	}
	end := len(userInterests) - 1
	users, err := h.Repo.FetchUsers(preference.MaxAge, preference.MinAge, preference.Gender, id)
	if err != nil {
		return nil, errors.New("error in fetching users")
	}
	scores := make([]float64, 0)
	matchUsers := make([]response.Home, 0)
	for _, u := range users {
		block, err := h.Repo.IsBlocked(id, u.Id)
		if err != nil {
			return nil, errors.New("error in fetching match data")
		}
		if block {
			continue
		}
		like, err := h.Repo.IsLikeExist(id, u.Id)
		if err != nil {
			return nil, errors.New("error in fetching match data")
		}
		if like {
			continue
		}
		distance := helper.HaversineDistance(u.Lattitude, u.Longitude, user.Lattitude, u.Longitude)
		if distance > float64(preference.MaxDistance) {
			continue
		}
		image, err := h.Repo.FetchImages(u.Id)
		if err != nil {
			return nil, errors.New("error in fetvhing images")
		}
		u.Images = image
		//interest
		interests, err := h.Repo.FetchInterests(u.Id)
		if err != nil {
			return nil, errors.New("error in fetching interests")
		}
		interestScore := 0
		for _, interest := range interests {
			search := helper.SearchForInterest(userInterests, interest, 0, end)
			if !search {
				continue
			}
			interestScore++

		}
		AgeScore := helper.Abs(u.Age - user.Age)
		//score
		score := 5*float64(interestScore) + (-2)*float64(AgeScore) + (-3)*distance
		scores = append(scores, score)
		matchUsers = append(matchUsers, u)
	}
	//sort
	helper.QuickSort(scores, matchUsers, 0, len(scores)-1)
	offset := (page - 1) * count
	if len(matchUsers) < offset+count {
		return matchUsers, nil
	}
	return matchUsers[offset : offset+count], nil
}
