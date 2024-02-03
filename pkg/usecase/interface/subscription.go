package useCaseInterface

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type SubscriptionUseCase interface {
	Add(sub models.Subscription) (response.Subscription, error)
	Update(sub models.Subscription, sID uint) (response.Subscription, error)
	Activate(sID uint) (response.Subscription, error)
	Dectivate(sID uint) (response.Subscription, error)
	Get() ([]response.GetSubscription, error)
	GetById(sID uint) (domain.Subscription, error)
	GetToUsers() ([]response.BriefSubscription, error)
	GetByIdToUsers(sID uint) (response.ShowSubscription, error)
	AddOrder(sID, userId uint) (response.Order, error)
	MakePayment(orderId uint) (response.OrderDetails,error)
	VerifyPayment(orderId uint,signature string,paymentId string)error
	GetOrders(userId uint)([]response.ShowOrder,error)
}
