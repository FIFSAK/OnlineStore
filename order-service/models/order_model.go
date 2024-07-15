package models

type Order struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	OrderDate  string  `json:"order_date"`
	Status     string  `json:"status"`
	ProductIDs []int   `json:"product_ids"`
}

type OrderModel interface {
	GetOrders() ([]*Order, error)
	CreateOrder(order Order) error
	GetOrderByID(id int) (*Order, error)
	UpdateOrder(order Order) error
	DeleteOrder(id int) error
	GetOrderByUserID(userID int) ([]*Order, error)
	GetOrderByStatus(status string) ([]*Order, error)
}
