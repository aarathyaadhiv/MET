package repository

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	DB *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) interfaces.SubscriptionRepository {
	return &SubscriptionRepository{db}
}

func (s *SubscriptionRepository) IsExist(name string) (bool, error) {
	var count int
	if err := s.DB.Raw(`SELECT COUNT(*) FROM subscriptions WHERE name=?`, name).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *SubscriptionRepository) Add(sub models.Subscription) (uint, error) {
	var id uint
	if err := s.DB.Raw(`INSERT INTO subscriptions(name,amount,days,likes,see_like) VALUES(?,?,?,?,?) RETURNING id`, sub.Name, sub.Amount, sub.Days, sub.Likes, sub.SeeLike).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SubscriptionRepository) Update(sub models.Subscription, sID uint) (uint, error) {
	var id uint
	if err := s.DB.Raw(`UPDATE subscriptions SET name=?,amount=?,days=?,likes=?,see_like=? WHERE id=? RETURNING id`, sub.Name, sub.Amount, sub.Days, sub.Likes, sub.SeeLike, sID).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SubscriptionRepository) Activate(sID uint) (uint, error) {
	var id uint
	if err := s.DB.Raw(`UPDATE subscriptions SET is_active=? WHERE id=? RETURNING id`, true, sID).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SubscriptionRepository) Deactivate(sID uint) (uint, error) {
	var id uint
	if err := s.DB.Raw(`UPDATE subscriptions SET is_active=? WHERE id=? RETURNING id`, false, sID).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SubscriptionRepository) IsActive(sID uint) (bool, error) {
	var IsActive bool
	if err := s.DB.Raw(`SELECT is_active FROM subscriptions WHERE id=?`, sID).Scan(&IsActive).Error; err != nil {
		return false, err
	}
	return IsActive, nil
}

func (s *SubscriptionRepository) Get() ([]response.GetSubscription, error) {
	var subscriptions []response.GetSubscription
	if err := s.DB.Raw(`SELECT id,name,amount,is_active FROM subscriptions`).Scan(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (s *SubscriptionRepository) GetById(sID uint) (domain.Subscription, error) {
	var subscription domain.Subscription
	if err := s.DB.Raw(`SELECT * FROM subscriptions WHERE id=?`, sID).Scan(&subscription).Error; err != nil {
		return domain.Subscription{}, err
	}
	return subscription, nil
}

func (s *SubscriptionRepository) GetToUsers() ([]response.BriefSubscription, error) {
	var subscriptions []response.BriefSubscription
	if err := s.DB.Raw(`SELECT id,name,amount FROM subscriptions WHERE is_active=?`, true).Scan(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (s *SubscriptionRepository) GetByIdToUsers(sID uint) (response.ShowSubscription, error) {
	var subscription response.ShowSubscription
	if err := s.DB.Raw(`SELECT id,name,amount,days,likes,see_like FROM subscriptions WHERE is_active=? AND id=?`, true, sID).Scan(&subscription).Error; err != nil {
		return response.ShowSubscription{}, err
	}
	return subscription, nil
}

func (s *SubscriptionRepository) AddOrder(order models.Order) (uint, error) {
	var id uint
	if err := s.DB.Raw(`INSERT INTO subscription_orders(subscription_id,user_id,subscribe_date,status) VALUES(?,?,?,?) RETURNING id`,order.SubscriptionId,order.UserId,order.SubscribeDate,order.Status).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SubscriptionRepository) GetDetailsForPayment(orderId uint)(models.OrderDetails,error){
	var order models.OrderDetails
	if err:=s.DB.Raw(`SELECT u.name AS user_name,s.amount FROM subscription_orders AS so JOIN users AS u ON so.user_id=u.id JOIN subscriptions s ON so.subscription_id=s.id WHERE so.id=?`,orderId).Scan(&order).Error;err!=nil{
		return models.OrderDetails{},err
	}
	return order,nil
}

func (s *SubscriptionRepository) AddRazorId(orderId uint, razorId string) error {
	if err := s.DB.Exec(`UPDATE subscription_orders SET razor_id=? WHERE id=?`, razorId, orderId).Error; err != nil {
		return err
	}
	return nil
}

func (s *SubscriptionRepository) AddPaymentId(orderId uint, paymentId,status string) (models.PaymentRes,error) {
	var res models.PaymentRes
	if err:= s.DB.Raw(`UPDATE subscription_orders SET payment_id=?,status=? WHERE id=? RETURNING user_id,subscription_id`, paymentId,status, orderId).Scan(&res).Error;err!=nil{
		return models.PaymentRes{},err
	}
	return res,nil
}

func (s *SubscriptionRepository) FetchRazorId(orderId uint) (string, error) {
	var razorId string
	if err := s.DB.Raw(`SELECT razor_id FROM subscription_orders WHERE id=?`, orderId).Scan(&razorId).Error; err != nil {
		return "", err
	}
	return razorId, nil
}

func (s *SubscriptionRepository) OrderStatus(orderId uint)(string,error){
	var status string
	if err:=s.DB.Raw(`SELECT status FROM subscription_orders WHERE id=?`,orderId).Scan(&status).Error;err!=nil{
		return "",err
	}
	return status,nil
}

func (s *SubscriptionRepository) MakeUserSubscribed(subUser models.PaymentRes)error{
	return s.DB.Exec(`UPDATE users SET is_subscribed=?,like_count=(SELECT likes FROM subscriptions WHERE id=?),subscription_id=? WHERE id=?`,true,subUser.SubscriptionId,subUser.SubscriptionId,subUser.UserId).Error
}