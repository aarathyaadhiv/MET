package usecase

import (
	"errors"
	"time"

	"github.com/aarathyaadhiv/met/pkg/config"
	"github.com/aarathyaadhiv/met/pkg/domain"
	interfaces "github.com/aarathyaadhiv/met/pkg/repository/interface"
	useCaseInterface "github.com/aarathyaadhiv/met/pkg/usecase/interface"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
	"github.com/razorpay/razorpay-go"
	"github.com/razorpay/razorpay-go/utils"
)

type SubscriptionUseCase struct {
	Repo   interfaces.SubscriptionRepository
	Config config.Config
}

func NewSubscriptionUseCase(repo interfaces.SubscriptionRepository, config config.Config) useCaseInterface.SubscriptionUseCase {
	return &SubscriptionUseCase{Repo: repo, Config: config}
}

func (s *SubscriptionUseCase) Add(sub models.Subscription) (response.Subscription, error) {
	isExist, err := s.Repo.IsExist(sub.Name)
	if err != nil {
		return response.Subscription{}, errors.New("error in fetcing subscription details")
	}
	if isExist {
		return response.Subscription{}, errors.New("already existing plan")
	}
	res, err := s.Repo.Add(sub)
	if err != nil {
		return response.Subscription{}, errors.New("error while adding")
	}
	return response.Subscription{SubscriptionId: res}, nil
}

func (s *SubscriptionUseCase) Update(sub models.Subscription, sID uint) (response.Subscription, error) {
	isExist,err:=s.Repo.IsExistById(sID)
	if err != nil {
		return response.Subscription{}, errors.New("error while fetching")
	}
	if !isExist{
		return response.Subscription{},errors.New("this subscription id is not existing")
	}
	res, err := s.Repo.Update(sub, sID)
	if err != nil {
		return response.Subscription{}, errors.New("error while updating")
	}
	return response.Subscription{SubscriptionId: res}, nil
}

func (s *SubscriptionUseCase) Activate(sID uint) (response.Subscription, error) {
	isExist,err:=s.Repo.IsExistById(sID)
	if err != nil {
		return response.Subscription{}, errors.New("error while fetching")
	}
	if !isExist{
		return response.Subscription{},errors.New("this subscription id is not existing")
	}
	isActive, err := s.Repo.IsActive(sID)
	if err != nil {
		return response.Subscription{}, errors.New("error in fetcing subscription details")
	}
	if isActive {
		return response.Subscription{}, errors.New("already plan is active")
	}
	res, err := s.Repo.Activate(sID)
	if err != nil {
		return response.Subscription{}, errors.New("error in activating plan")
	}
	return response.Subscription{SubscriptionId: res}, nil
}

func (s *SubscriptionUseCase) Dectivate(sID uint) (response.Subscription, error) {
	isExist,err:=s.Repo.IsExistById(sID)
	if err != nil {
		return response.Subscription{}, errors.New("error while fetching")
	}
	if !isExist{
		return response.Subscription{},errors.New("this subscription id is not existing")
	}
	isActive, err := s.Repo.IsActive(sID)
	if err != nil {
		return response.Subscription{}, errors.New("error in fetcing subscription details")
	}
	if !isActive {
		return response.Subscription{}, errors.New("already plan is deactive")
	}
	res, err := s.Repo.Deactivate(sID)
	if err != nil {
		return response.Subscription{}, errors.New("error in activating plan")
	}
	return response.Subscription{SubscriptionId: res}, nil
}

func (s *SubscriptionUseCase) Get() ([]response.GetSubscription, error) {
	res, err := s.Repo.Get()
	if err != nil {
		return nil, errors.New("error in fetching subscription")
	}
	return res, nil
}

func (s *SubscriptionUseCase) GetById(sID uint) (domain.Subscription, error) {
	res, err := s.Repo.GetById(sID)
	if err != nil {
		return domain.Subscription{}, errors.New("error in fetching subscription")
	}
	return res, nil
}

func (s *SubscriptionUseCase) GetToUsers() ([]response.BriefSubscription, error) {
	res, err := s.Repo.GetToUsers()
	if err != nil {
		return nil, errors.New("error in fetching data")
	}
	return res, nil
}

func (s *SubscriptionUseCase) GetByIdToUsers(sID uint) (response.ShowSubscription, error) {
	isExist,err:=s.Repo.IsExistById(sID)
	if err != nil {
		return response.ShowSubscription{}, errors.New("error while fetching")
	}
	if !isExist{
		return response.ShowSubscription{},errors.New("this subscription id is not existing")
	}
	res, err := s.Repo.GetByIdToUsers(sID)
	if err != nil {
		return response.ShowSubscription{}, err
	}
	return res, nil
}

func (s *SubscriptionUseCase) AddOrder(sID, userId uint) (response.Order, error) {
	isExist,err:=s.Repo.IsExistById(sID)
	if err != nil {
		return response.Order{}, errors.New("error while fetching")
	}
	if !isExist{
		return response.Order{},errors.New("this subscription id is not existing")
	}
	order := models.Order{
		SubscriptionId: sID,
		UserId:         userId,
		SubscribeDate:  time.Now(),
		Status:         "pending",
	}
	res, err := s.Repo.AddOrder(order)
	if err != nil {
		return response.Order{}, err
	}
	return response.Order{OrderId: res}, nil
}

func (s *SubscriptionUseCase) MakePayment(orderId uint) (response.OrderDetails, error) {
	isExist,err:=s.Repo.IsOrderExist(orderId)
	if err != nil {
		return response.OrderDetails{}, errors.New("error in fetching data")
	}
	if !isExist{
		return response.OrderDetails{},errors.New("orderId is not existing")
	}
	status, err := s.Repo.OrderStatus(orderId)
	if err != nil {
		return response.OrderDetails{}, err
	}
	if status == "paid" {
		return response.OrderDetails{}, errors.New("already paid")
	}
	client := razorpay.NewClient(s.Config.RAZORPAY_KEY, s.Config.RAZORPAY_SECRET)

	order, err := s.Repo.GetDetailsForPayment(orderId)
	if err != nil {
		return response.OrderDetails{}, err
	}

	data := map[string]interface{}{
		"amount":   int(order.Amount * 100),
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return response.OrderDetails{}, err
	}
	razorId := body["id"].(string)
	err = s.Repo.AddRazorId(orderId, razorId)
	if err != nil {
		return response.OrderDetails{}, err
	}
	return response.OrderDetails{
		UserName:   order.UserName,
		Amount:     order.Amount,
		RazorId:    razorId,
		OrderId:    orderId,
		AmountPisa: int(order.Amount * 100),
	}, nil

}

func (s *SubscriptionUseCase) VerifyPayment(orderId uint, signature string, paymentId string) error {
	isExist,err:=s.Repo.IsOrderExist(orderId)
	if err != nil {
		return  errors.New("error in fetching data")
	}
	if !isExist{
		return errors.New("orderId is not existing")
	}
	status, err := s.Repo.OrderStatus(orderId)
	if err != nil {
		return err
	}
	if status == "paid" {
		return errors.New("already paid")
	}
	razorId, err := s.Repo.FetchRazorId(orderId)
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"razorpay_order_id":   razorId,
		"razorpay_payment_id": paymentId,
	}
	isValid := utils.VerifyPaymentSignature(params, signature, s.Config.RAZORPAY_SECRET)
	if isValid {
		res, err := s.Repo.AddPaymentId(orderId, paymentId, "paid")
		if err != nil {
			return err
		}
		return s.Repo.MakeUserSubscribed(res)

	}
	return errors.New("the payment received is not from the authenticated resource")
}


func (s *SubscriptionUseCase) GetOrders(userId uint)([]response.ShowOrder,error){
	return s.Repo.ShowOrders(userId)
}
