package repository

import (
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"gorm.io/gorm"
)



type LikeRepository struct{
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB)interfaces.LikeRepository{
	return &LikeRepository{db}
}

func(l *LikeRepository)Like(likedId,userId uint)(response.Like,error){
	var like response.Like
	if err:=l.DB.Raw(`INSERT INTO likes(user_id,liked_id) VALUES(?,?) RETURNING user_id,liked_id`,userId,likedId).Scan(&like).Error;err!=nil{
		return response.Like{},err
	}
	return like,nil
}
func (l *LikeRepository)Unlike(likeId,userId uint)(response.Like,error){
	var like response.Like
	if err:=l.DB.Raw(`DELETE FROM likes WHERE user_id=? AND liked_id=? RETURNING user_id,liked_id`,userId,likeId).Scan(&like).Error;err!=nil{
		return response.Like{},err
	}
	return response.Like{},nil
}

func (l *LikeRepository) GetLike(page,count int,userId uint)([]uint,error){
	var likedId []uint
	offset:=(page-1)*count
	if err:=l.DB.Raw(`SELECT user_id FROM likes WHERE liked_id=? limit ? offset ? `,userId,count,offset).Scan(&likedId).Error;err!=nil{
		return nil,err
	}
	return likedId,nil
}

func (l *LikeRepository) IsLikeExist(userId,likedId uint)(bool,error){
	var count int
	if err:=l.DB.Raw(`SELECT COUNT(*) FROM likes WHERE user_id=? AND liked_id=?`,userId,likedId).Scan(&count).Error;err!=nil{
		return false,err
	}
	return count>0,nil
}

func (l *LikeRepository) Match(userId,matchId uint)error{
	if err:=l.DB.Exec(`INSERT INTO matches(user_id,match_id) VALUES(?,?)`,userId,matchId).Error;err!=nil{
		return err
	}
	return nil
}

func (l *LikeRepository) GetMatch(page,count int,userId uint)([]uint,error){
	var id []uint
	offset:=(page-1)*count
	if err:=l.DB.Raw(`SELECT match_id from matches WHERE user_id=? limit ? offset ?`,userId,count,offset).Scan(&id).Error;err!=nil{
		return nil,err
	}
	return id,nil
}

func (l *LikeRepository) UnMatch(userId,matchId uint)(response.UnMatch,error){
	var unMatch response.UnMatch
	if err:=l.DB.Exec(`DELETE FROM matches WHERE user_id=? AND match_id=? `,matchId,userId).Error;err!=nil{
		return response.UnMatch{},err
	}
	if err:=l.DB.Raw(`DELETE FROM matches WHERE user_id=? AND match_id=? RETURNING user_id,match_id`,userId,matchId).Scan(&unMatch).Error;err!=nil{
		return response.UnMatch{},err
	}
	return unMatch,nil
}

func (l *LikeRepository) IsMatchExist(userId,matchId uint)(bool,error){
	var count int
	if err:=l.DB.Raw(`SELECT COUNT(*) FROM matches WHERE user_id=? AND match_id=?`,userId,matchId).Scan(&count).Error;err!=nil{
		return false,err
	}
	return count>0,nil
}