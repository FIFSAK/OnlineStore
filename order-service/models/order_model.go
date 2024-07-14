package models

type Order struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	ProductID  int    `json:"product_id"`
	TotalPrice int    `json:"total_price"`
	OrderDate  string `json:"order_date"`
	Status     string `json:"status"`
}
