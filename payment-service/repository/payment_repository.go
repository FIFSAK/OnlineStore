package repository

import (
	"OnlineStore/payment-service/models"
	"database/sql"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

func (pr *PaymentRepository) GetPayments() ([]*models.Payment, error) {
	rows, err := pr.DB.Query("SELECT * FROM payments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []*models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.ID, &payment.UserID, &payment.OrderID, &payment.Amount, &payment.PaymentDate, &payment.PaymentStatus)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil
}

func (pr *PaymentRepository) CreatePayment(payment models.Payment) error {
	_, err := pr.DB.Exec("INSERT INTO payments (user_id, order_id, amount, payment_status) VALUES ($1, $2, $3, $4)", payment.UserID, payment.OrderID, payment.Amount, payment.PaymentStatus)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PaymentRepository) GetPaymentByID(id int) (*models.Payment, error) {
	var payment models.Payment
	err := pr.DB.QueryRow("SELECT * FROM payments WHERE id = $1", id).Scan(&payment.ID, &payment.UserID, &payment.OrderID, &payment.Amount, &payment.PaymentDate, &payment.PaymentStatus)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (pr *PaymentRepository) UpdatePayment(payment models.Payment) error {
	_, err := pr.DB.Exec("UPDATE payments SET user_id = $1, order_id = $2, amount = $3 WHERE id = $5", payment.UserID, payment.OrderID, payment.Amount, payment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PaymentRepository) DeletePayment(id int) error {
	_, err := pr.DB.Exec("DELETE FROM payments WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PaymentRepository) GetPaymentByOrderID(orderID int) ([]*models.Payment, error) {
	rows, err := pr.DB.Query("SELECT * FROM payments WHERE order_id = $1", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []*models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.ID, &payment.UserID, &payment.OrderID, &payment.Amount, &payment.PaymentDate, &payment.PaymentStatus)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil
}

func (pr *PaymentRepository) GetPaymentByUserID(userID int) ([]*models.Payment, error) {
	rows, err := pr.DB.Query("SELECT * FROM payments WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []*models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.ID, &payment.UserID, &payment.OrderID, &payment.Amount, &payment.PaymentDate, &payment.PaymentStatus)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil
}

func (pr *PaymentRepository) GetPaymentByStatus(status string) ([]*models.Payment, error) {
	rows, err := pr.DB.Query("SELECT * FROM payments WHERE payment_status = $1", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []*models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.ID, &payment.UserID, &payment.OrderID, &payment.Amount, &payment.PaymentDate, &payment.PaymentStatus)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil
}
