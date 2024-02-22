package repository

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepository{db}
}

func (a *AdminRepository) IsAdminExist(email string) bool {
	var count int
	if err := a.DB.Raw(`SELECT COUNT(*) FROM admins WHERE email=?`, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (a *AdminRepository) Save(admin models.Admin) (uint, error) {
	var id uint
	if err := a.DB.Raw(`INSERT INTO admins(email,password) VALUES(?,?) RETURNING id`, admin.Email, admin.Password).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AdminRepository) FetchAdmin(email string) (domain.Admin, error) {
	var admin domain.Admin
	if err := a.DB.Raw(`SELECT * FROM admins WHERE email=?`, email).Scan(&admin).Error; err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}

func (a *AdminRepository) IsUserExist(id uint) (bool, error) {
	var count int
	if err := a.DB.Raw(`SELECT COUNT(*) FROM users WHERE id=?`, id).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *AdminRepository) BlockUser(id uint) (uint, error) {
	var userId uint
	if err := a.DB.Raw(`UPDATE users SET is_block=true WHERE id=? RETURNING id `, id).Scan(&userId).Error; err != nil {
		return 0, err
	}
	return userId, nil
}

func (a *AdminRepository) UnblockUser(id uint) (uint, error) {
	var userId uint
	if err := a.DB.Raw(`UPDATE users SET is_block=false WHERE id=? RETURNING id`, id).Scan(&userId).Error; err != nil {
		return 0, err
	}
	return userId, nil
}

func (a *AdminRepository) GetUsers(page, count int) ([]response.User, error) {
	offset := (page - 1) * count
	var users []response.User
	if err := a.DB.Raw(`SELECT u.id,u.name,u.age,u.ph_no,g.name as gender,u.city,u.country,u.is_block FROM users as u JOIN genders as g ON u.gender_id=g.id limit ? offset ?`, count, offset).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (a *AdminRepository) IsUserBlocked(id uint) (bool, error) {
	var isBlock bool
	if err := a.DB.Raw(`SELECT is_block FROM users WHERE id=?`, id).Scan(&isBlock).Error; err != nil {
		return false, err
	}
	return isBlock, nil
}

func (a *AdminRepository) GetSingleUser(id uint) (response.UserDetailsToAdmin, error) {
	var user response.UserDetailsToAdmin
	if err := a.DB.Raw(`SELECT
    u.id,
    u.name,
    TO_CHAR(DATE(u.dob),'YYYY FMMonth DD') AS dob,
    u.age,
    u.ph_no,
    g.name AS gender,
    u.city,
    u.country,
    u.longitude,
    u.lattitude,
    u.bio,
    (
        SELECT STRING_AGG(i.image, ', ' ORDER BY i.image)
        FROM images AS i
        WHERE i.user_id = u.id
    ) AS images,
    (
        SELECT STRING_AGG(t.interest, ', ' ORDER BY t.interest)
        FROM user_interests AS ut
        JOIN interests AS t ON ut.interest_id = t.id
        WHERE ut.user_id = u.id
    ) AS interests
FROM
    users AS u
JOIN
    genders AS g ON g.id = u.gender_id
WHERE
    u.id = ?
`, id).Scan(&user).Error; err != nil {
		return response.UserDetailsToAdmin{}, err
	}
	return user, nil
}

func (a *AdminRepository) ReportedUsers() ([]response.ReportedUsers, error) {
	var reported []response.ReportedUsers
	if err := a.DB.Raw(`SELECT r.reported_id as id,u.name,COUNT(r.reported_id) as report_count FROM reported_users AS r JOIN users AS u ON u.id=r.reported_id GROUP BY r.reported_id,u.name`).Scan(&reported).Error; err != nil {
		return nil, err
	}
	return reported, nil
}

func (a *AdminRepository) IsReportedUser(id uint) (bool, error) {
	var count int
	if err := a.DB.Raw(`SELECT COUNT(*) FROM reported_users WHERE reported_id=?`, id).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (a *AdminRepository) ReportedUser(reportId uint) (response.ReportedUser, error) {
	var report response.ReportedUser
	if err := a.DB.Raw(`SELECT r.reported_id as id,u1.name,r.user_id as reported_by_id,u2.name as reported_by_name,r.message,r.time FROM reported_users as r JOIN users AS u1 ON r.reported_id=u1.id JOIN users u2 ON r.user_id=u2.id WHERE r.reported_id=?`, reportId).Scan(&report).Error; err != nil {
		return response.ReportedUser{}, err
	}
	return report, nil
}
