package repository

import (
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"gorm.io/gorm"
)

type HomeRepository struct {
	DB *gorm.DB
}

func NewHomeRepository(db *gorm.DB) interfaces.HomeRepository {
	return &HomeRepository{db}
}

func (h *HomeRepository) FetchUser(id uint) (models.FetchUser, error) {
	var user models.FetchUser
	if err := h.DB.Raw(`SELECT longitude,lattitude,age FROM users WHERE id=?`, id).Scan(&user).Error; err != nil {
		return models.FetchUser{}, err
	}
	return user, nil
}

func (h *HomeRepository) FetchUsers(maxAge, minAge int, gender, id uint) ([]response.Home, error) {
	var users []response.Home
	if err := h.DB.Raw(`SELECT u.id,u.name,TO_CHAR(DATE(u.dob), 'YYYY FMMonth DD') AS dob,u.age,g.name as gender,u.city,u.country,u.longitude,u.lattitude,u.bio FROM users AS u 
	JOIN 
	genders AS g ON g.id=u.gender_id
	WHERE u.age>? AND u.age<? AND u.gender_id=? AND u.id<>?`, minAge, maxAge, gender, id).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (h *HomeRepository) FetchPreference(id uint) (models.Preference, error) {
	var preference models.Preference
	if err := h.DB.Raw(`SELECT min_age,max_age,gender,max_distance FROM preferences WHERE user_id=?`, id).Scan(&preference).Error; err != nil {
		return models.Preference{}, err
	}
	return preference, nil
}

func (h *HomeRepository) FetchImages(id uint) ([]string, error) {
	var images []string
	if err := h.DB.Raw(`SELECT image FROM images WHERE user_id=?`, id).Scan(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (h *HomeRepository) FetchInterests(id uint) ([]string, error) {
	var interests []string
	if err := h.DB.Raw(`SELECT i.interest FROM interests i JOIN user_interests u ON u.interest_id=i.id WHERE u.user_id=? ORDER BY i.interest`, id).Scan(&interests).Error; err != nil {
		return nil, err
	}
	return interests, nil
}

func (h *HomeRepository) FetchInterestId(id uint) ([]uint, error) {
	var interests []uint
	if err := h.DB.Raw(`SELECT interest_id FROM user_interests WHERE user_id=?`, id).Scan(&interests).Error; err != nil {
		return nil, err
	}
	return interests, nil
}

func (h *HomeRepository) IsLikeExist(userId, likedId uint) (bool, error) {
	var count int
	if err := h.DB.Raw(`SELECT COUNT(*) FROM likes WHERE user_id=? AND liked_id=?`, userId, likedId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (h *HomeRepository) IsBlocked(userId, blockedId uint) (bool, error) {
	var count int
	if err := h.DB.Raw(`SELECT COUNT(*) FROM blocked_users WHERE user_id=? AND blocked_id=?`, userId, blockedId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (h *HomeRepository) FetchUserWithInterest(id uint, interestId []uint) ([]response.Home, error) {
	var users []response.Home
	if err := h.DB.Raw(`SELECT DISTINCT u.id,u.name,TO_CHAR(DATE(u.dob), 'YYYY FMMonth DD') AS dob,u.age,g.name as gender,u.city,u.country,u.longitude,u.lattitude,u.bio FROM user_interests AS i JOIN users AS u ON i.user_id=u.id JOIN genders AS g ON g.id=u.gender_id WHERE u.id<>? AND i.interest_id IN(?)`, id, interestId).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (h *HomeRepository) IsInterestValid(id uint, interestId uint) (bool, error) {
	var count int
	if err := h.DB.Raw(`SELECT COUNT(*) FROM user_interests WHERE user_id=? AND interest_id=?`, id, interestId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (h *HomeRepository) FetchUserByInterest(id uint, interestId uint) ([]response.Home, error) {
	var users []response.Home
	if err := h.DB.Raw(`SELECT DISTINCT u.id,u.name,TO_CHAR(DATE(u.dob), 'YYYY FMMonth DD') AS dob,u.age,g.name as gender,u.city,u.country,u.longitude,u.lattitude,u.bio FROM user_interests as i JOIN users AS u ON u.id=i.user_id JOIN genders g ON u.gender_id=g.id WHERE u.id<>? AND i.interest_id=?`, id, interestId).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
