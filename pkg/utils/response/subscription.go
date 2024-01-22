package response

type GetSubscription struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	IsActive bool    `json:"is_active"`
}

// to users
type ShowSubscription struct {
	Id      uint    `json:"id"`
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	Days    int     `json:"days"`
	Likes   int     `json:"likes"`
	SeeLike bool    `json:"see_like"`
}

type BriefSubscription struct {
	Id     uint    `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type Subscription struct {
	SubscriptionId uint
}

type Order struct {
	OrderId uint
}

type OrderDetails struct {
	UserName   string
	Amount     float64
	RazorId    string
	OrderId    uint
	AmountPisa int
}
