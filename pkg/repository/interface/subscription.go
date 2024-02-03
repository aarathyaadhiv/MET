package interfaces

import (
	"github.com/aarathyaadhiv/met/pkg/domain"
	"github.com/aarathyaadhiv/met/pkg/utils/models"
	"github.com/aarathyaadhiv/met/pkg/utils/response"
)

type SubscriptionRepository interface {
	IsExist(name string)(bool,error)
	IsExistById(sId uint)(bool,error)
	Add(sub models.Subscription) (uint, error)
	Update(sub models.Subscription, sID uint) (uint, error)
	Activate(sID uint) (uint, error)
	Deactivate(sID uint) (uint, error)
	IsActive(sID uint) (bool, error)
	Get() ([]response.GetSubscription, error)
	GetById(sID uint) (domain.Subscription, error)
	GetToUsers()([]response.BriefSubscription,error)
	GetByIdToUsers(sID uint) (response.ShowSubscription, error)  
	AddOrder(order models.Order)(uint,error)
	IsOrderExist(orderId uint)(bool,error)
	GetDetailsForPayment(orderId uint)(models.OrderDetails,error)
	AddRazorId(orderId uint,razorId string)error
	AddPaymentId(orderId uint,paymentId,status string)(models.PaymentRes,error)
	FetchRazorId(orderId uint)(string,error)
	OrderStatus(orderId uint)(string,error)
	MakeUserSubscribed(subUser models.PaymentRes)error
	ShowOrders(userId uint)([]response.ShowOrder,error)
}
