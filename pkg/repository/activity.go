package repository

import (
	"time"

	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	DB *gorm.DB
}

func NewActivityRepository(db *gorm.DB) interfaces.ActivityRepository {
	return &ActivityRepository{db}
}

func (l *ActivityRepository) IsUserExist(id uint) (bool, error) {
	var count int
	if err := l.DB.Raw(`SELECT COUNT(*) FROM users WHERE id=?`, id).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (l *ActivityRepository) Like(likedId, userId uint) (response.Like, error) {
	var like response.Like
	if err := l.DB.Raw(`INSERT INTO likes(user_id,liked_id,time) VALUES(?,?,?) RETURNING user_id,liked_id`, userId, likedId, time.Now()).Scan(&like).Error; err != nil {
		return response.Like{}, err
	}
	return like, nil
}
func (l *ActivityRepository) Unlike(likeId, userId uint) (response.Like, error) {
	var like response.Like
	if err := l.DB.Raw(`DELETE FROM likes WHERE user_id=? AND liked_id=? RETURNING user_id,liked_id`, userId, likeId).Scan(&like).Error; err != nil {
		return response.Like{}, err
	}
	return response.Like{}, nil
}

func (l *ActivityRepository) GetLike(page, count int, userId uint) ([]response.ShowUserDetails, error) {
	var like []response.ShowUserDetails
	offset := (page - 1) * count
	if err := l.DB.Raw(`SELECT l.user_id as id,u.name,TO_CHAR(DATE(u.dob),'YYYY FMMonth DD') AS dob,u.age,g.name as gender,u.city,u.country,u.bio,(
        SELECT STRING_AGG(i.image, ', ' ORDER BY i.image)
        FROM images AS i
        WHERE i.user_id = l.user_id
    ) AS image  FROM likes as l JOIN users as u ON l.user_id=u.id JOIN genders as g ON u.gender_id=g.id  WHERE l.liked_id=? limit ? offset ? `, userId, count, offset).Scan(&like).Error; err != nil {
		return nil, err
	}
	return like, nil
}

func (l *ActivityRepository) GetLikeCount(userId uint)(int,error){
	var count int
	if err:=l.DB.Raw(`SELECT COUNT(*) FROM likes WHERE liked_id=?`,userId).Scan(&count).Error;err!=nil{
		return 0,err
	}
	return count,nil
}

func (a *ActivityRepository) FetchInterests(id uint) ([]string, error) {
	var interests []string
	if err := a.DB.Raw(`SELECT i.interest FROM user_interests as u JOIN interests as i ON u.interest_id=i.id WHERE u.user_id=?`, id).Scan(&interests).Error; err != nil {
		return nil, err
	}
	return interests, nil
}

func (l *ActivityRepository) IsLikeExist(userId, likedId uint) (bool, error) {
	var count int
	if err := l.DB.Raw(`SELECT COUNT(*) FROM likes WHERE user_id=? AND liked_id=?`, userId, likedId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (l *ActivityRepository) LikeCount(userId uint) (int, error) {
	var count int
	if err := l.DB.Raw(`SELECT like_count FROM users WHERE id=?`, userId).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (l *ActivityRepository) UpdateLikeCount(userId uint, count int) error {
	return l.DB.Exec(`UPDATE users SET like_count=? WHERE id=?`, count, userId).Error
}

func (l *ActivityRepository) IsSubscribed(userId uint) (bool, error) {
	var isSubscribed bool
	if err := l.DB.Raw(`SELECT is_subscribed FROM users WHERE id=?`, userId).Scan(&isSubscribed).Error; err != nil {
		return false, err
	}
	return isSubscribed, nil
}

func (l *ActivityRepository) SeeLike(userId uint) (bool, error) {
	var seeLike bool
	if err := l.DB.Raw(`SELECT s.see_like FROM users AS u JOIN subscriptions AS s ON u.subscription_id=s.id WHERE u.id=?`, userId).Scan(&seeLike).Error; err != nil {
		return false, err
	}
	return seeLike, nil
}

func (l *ActivityRepository) Match(userId, matchId uint) error {
	if err := l.DB.Exec(`INSERT INTO matches(user_id,match_id,time) VALUES(?,?,?)`, userId, matchId, time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (l *ActivityRepository) GetMatch(page, count int, userId uint) ([]response.ShowUserDetails, error) {
	var match []response.ShowUserDetails
	offset := (page - 1) * count
	if err := l.DB.Raw(`SELECT m.match_id as id,u.name,TO_CHAR(DATE(u.dob),'YYYY FMMonth DD') AS dob,u.age,g.name as gender,u.city,u.country,u.bio,(
        SELECT STRING_AGG(i.image, ', ' ORDER BY i.image)
        FROM images AS i
        WHERE i.user_id = m.match_id
    ) AS image FROM matches as m JOIN users as u ON m.match_id=u.id JOIN genders as g ON u.gender_id=g.id   WHERE m.user_id=? limit ? offset ?`, userId, count, offset).Scan(&match).Error; err != nil {
		return nil, err
	}
	return match, nil
}

func (l *ActivityRepository) UnMatch(userId, matchId uint) (response.UnMatch, error) {
	var unMatch response.UnMatch
	if err := l.DB.Exec(`DELETE FROM matches WHERE user_id=? AND match_id=? `, matchId, userId).Error; err != nil {
		return response.UnMatch{}, err
	}
	if err := l.DB.Raw(`DELETE FROM matches WHERE user_id=? AND match_id=? RETURNING user_id,match_id`, userId, matchId).Scan(&unMatch).Error; err != nil {
		return response.UnMatch{}, err
	}
	return unMatch, nil
}

func (l *ActivityRepository) IsMatchExist(userId, matchId uint) (bool, error) {
	var count int
	if err := l.DB.Raw(`SELECT COUNT(*) FROM matches WHERE user_id=? AND match_id=?`, userId, matchId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (l *ActivityRepository) IsReported(userId, reportId uint) (bool, error) {
	var count int
	if err := l.DB.Raw(`SELECT COUNT(*) FROM reported_users WHERE user_id=? AND reported_id=? `, userId, reportId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (l *ActivityRepository) Report(userId, reportId uint, message string) (response.Report, error) {
	var report response.Report
	if err := l.DB.Raw(`INSERT INTO reported_users(user_id,reported_id,message,time) VALUES(?,?,?,?) RETURNING user_id,reported_id`, userId, reportId, message, time.Now()).Scan(&report).Error; err != nil {
		return response.Report{}, err
	}
	return report, nil
}

func (l *ActivityRepository) IsBlocked(userId, blockedId uint) (bool, error) {
	var count int
	if err := l.DB.Raw(`SELECT COUNT(*) FROM blocked_users WHERE user_id=? AND blocked_id=?`, userId, blockedId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (l *ActivityRepository) BlockUser(userId, blockedId uint) (response.BlockUser, error) {
	var block response.BlockUser
	if err := l.DB.Raw(`INSERT INTO blocked_users(user_id,blocked_id) VALUES(?,?) RETURNING user_id,blocked_id`, userId, blockedId).Scan(&block).Error; err != nil {
		return response.BlockUser{}, err
	}
	return block, nil
}
