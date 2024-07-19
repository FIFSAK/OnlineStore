package models

type Payment struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	OrderID       int     `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentDate   string  `json:"payment_date"`
	PaymentStatus string  `json:"payment_status"`
}

type PaymentModel interface {
	GetPayments() ([]*Payment, error)
	CreatePayment(payment Payment) error
	GetPaymentByID(id int) (*Payment, error)
	UpdatePayment(payment Payment) error
	DeletePayment(id int) error
	GetPaymentByOrderID(orderID int) ([]*Payment, error)
	GetPaymentByUserID(userID int) ([]*Payment, error)
	GetPaymentByStatus(status string) ([]*Payment, error)
}
