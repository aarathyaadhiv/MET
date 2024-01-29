package repository

import (
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) IsUserExist(phNo string) (bool, error) {
	var count int
	if err := u.DB.Raw(`SELECT COUNT(*) FROM users WHERE ph_no=? `, phNo).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *UserRepository) IsUserBlocked(phNo string) (bool, error) {
	var block bool
	if err := u.DB.Raw(`SELECT is_block FROM users WHERE ph_no=?`, phNo).Scan(&block).Error; err != nil {
		return false, err
	}
	return block, nil
}

func (u *UserRepository) FindByPhone(phNo string) (uint, error) {
	var id uint
	if err := u.DB.Raw(`SELECT id FROM users WHERE ph_no=?`, phNo).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserRepository) CreateUserId(phNo string) (uint, error) {
	var id uint
	if err := u.DB.Raw(`INSERT INTO users(ph_no) VALUES(?) RETURNING id `, phNo).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserRepository) UpdateUser(id uint, profile models.ProfileSave) (uint, error) {
	var userId uint
	if err := u.DB.Raw(`UPDATE users SET name=?,dob=?,age=?,gender_id=?,city=?,country=?,longitude=?,lattitude=?,bio=? WHERE id=? RETURNING id`, profile.Name, profile.Dob, profile.Age, profile.GenderId, profile.City, profile.Country, profile.Longitude, profile.Lattitude, profile.Bio, id).Scan(&userId).Error; err != nil {
		return 0, err
	}
	return userId, nil
}

func (u *UserRepository) UpdateUserDetails(id uint, user models.UpdateUserDetails) error {
	if err := u.DB.Exec(`UPDATE users SET ph_no=?,city=?,country=?,bio=? WHERE id=?`, user.PhNo, user.City, user.Country, user.Bio, id).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdateLocation(id uint, location models.UpdateLocation) (uint, error) {
	var userId uint
	if err := u.DB.Raw(`UPDATE users SET longitude=?,lattitude=? WHERE id=? RETURNING id`, location.Longitude, location.Lattitude, id).Scan(&id).Error; err != nil {
		return 0, err
	}
	return userId, nil
}

func (u *UserRepository) IsInterestExist(id,interest uint)(bool,error){
	var count int
	if err:=u.DB.Raw(`SELECT COUNT(*) FROM user_interests WHERE user_id=? AND interest_id=?`,id,interest).Scan(&count).Error;err!=nil{
		return false,err
	}
	return count>0,nil
}

func (u *UserRepository) AddInterest(id, interest uint) error {
	if err := u.DB.Exec(`INSERT INTO user_interests(user_id,interest_id) values(?,?)`, id, interest).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteInterest(id uint) error {
	if err := u.DB.Exec(`DELETE FROM user_interests WHERE user_id=?`, id).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) IsImageExist(id uint,image string)(bool,error){
	var count int
	if err:=u.DB.Raw(`SELECT COUNT(*) FROM images WHERE user_id=? AND image=?`,id,image).Scan(&count).Error;err!=nil{
		return false,err
	}
	return count>0,nil
}

func (u *UserRepository) AddImage(id uint, image string) error {
	if err := u.DB.Exec(`INSERT INTO images(user_id,image) values(?,?)`, id, image).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteImage(id uint) error {
	if err := u.DB.Exec(`DELETE FROM images WHERE user_id=?`, id).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) ShowProfile(id uint) (response.UserDetails, error) {
	var user response.UserDetails
	if err := u.DB.Raw(`SELECT u.id,u.name,TO_CHAR(DATE(u.dob), 'YYYY-FMMonth-DD') AS dob,u.age,u.ph_no,g.name as gender,u.city,u.country,u.longitude,u.lattitude,u.bio from users as u JOIN genders as g ON u.gender_id=g.id WHERE u.id=?`, id).Scan(&user).Error; err != nil {
		return response.UserDetails{}, err
	}
	return user, nil
}

func (u *UserRepository) FetchImages(id uint) ([]string, error) {
	var images []string
	if err := u.DB.Raw(`SELECT image FROM images WHERE user_id=?`, id).Scan(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (u *UserRepository) FetchInterests(id uint) ([]string, error) {
	var interests []string
	if err := u.DB.Raw(`SELECT i.interest FROM user_interests as u JOIN interests as i ON u.interest_id=i.id WHERE u.user_id=?`, id).Scan(&interests).Error; err != nil {
		return nil, err
	}
	return interests, nil
}

func (u *UserRepository) IsBlocked(id uint) (bool, error) {
	var block bool
	if err := u.DB.Raw(`SELECT is_block FROM users WHERE id=?`, id).Scan(&block).Error; err != nil {
		return false, err
	}
	return block, nil
}

func (u *UserRepository) AddPreference(id uint, preference models.Preference) error {
	if err := u.DB.Exec(`INSERT INTO preferences(user_id,min_age,max_age,gender,max_distance) VALUES(?,?,?,?,?)`, id, preference.MinAge, preference.MaxAge, preference.Gender, preference.MaxDistance).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdatePreference(id uint, preference models.Preference) (uint, error) {
	var userId uint
	if err := u.DB.Raw(`UPDATE preferences SET min_age=?,max_age=?,gender=?,max_distance=? WHERE user_id=? RETURNING user_id`, preference.MinAge, preference.MaxAge, preference.Gender, preference.MaxDistance, id).Scan(&userId).Error; err != nil {
		return 0, err
	}
	return userId, nil

}

func (u *UserRepository) GetPreference(id uint) (models.Preference, error) {
	var preference models.Preference
	if err := u.DB.Raw(`SELECT min_age,max_age,gender,max_distance FROM preferences WHERE user_id=?`, id).Scan(&preference).Error; err != nil {
		return models.Preference{}, err
	}
	return preference, nil
}

func (u *UserRepository) FetchShortDetail(id uint)(models.UserShortDetail,error){
	var user models.UserShortDetail
	if err:=u.DB.Raw(`SELECT u.id,u.name,i.image FROM users AS u JOIN images AS i ON i.user_id=u.id WHERE u.id=? `,id).Scan(&user).Error;err!=nil{
		return models.UserShortDetail{},err
	}
	return user,nil
}
