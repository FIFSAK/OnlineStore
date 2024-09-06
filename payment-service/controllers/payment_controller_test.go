package controllers

import (
	"OnlineStore/payment-service/models"
	"OnlineStore/payment-service/services"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"database/sql"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// MockPaymentModel is a mock implementation of the PaymentModel interface
type MockPaymentModel struct {
	Payments []*models.Payment
}

func (m *MockPaymentModel) GetPayments() ([]*models.Payment, error) {
	return m.Payments, nil
}

func (m *MockPaymentModel) CreatePayment(payment models.Payment) error {
	m.Payments = append(m.Payments, &payment)
	return nil
}

func (m *MockPaymentModel) GetPaymentByID(id int) (*models.Payment, error) {
	for _, payment := range m.Payments {
		if payment.ID == id {
			return payment, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (m *MockPaymentModel) UpdatePayment(payment models.Payment) error {
	for i, p := range m.Payments {
		if p.ID == payment.ID {
			m.Payments[i] = &payment
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockPaymentModel) DeletePayment(id int) error {
	for i, payment := range m.Payments {
		if payment.ID == id {
			m.Payments = append(m.Payments[:i], m.Payments[i+1:]...)
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockPaymentModel) GetPaymentByOrderID(orderID int) ([]*models.Payment, error) {
	var payments []*models.Payment
	for _, payment := range m.Payments {
		if payment.OrderID == orderID {
			payments = append(payments, payment)
		}
	}
	return payments, nil
}

func (m *MockPaymentModel) GetPaymentByUserID(userID int) ([]*models.Payment, error) {
	var payments []*models.Payment
	for _, payment := range m.Payments {
		if payment.UserID == userID {
			payments = append(payments, payment)
		}
	}
	return payments, nil
}

func (m *MockPaymentModel) GetPaymentByStatus(status string) ([]*models.Payment, error) {
	var payments []*models.Payment
	for _, payment := range m.Payments {
		if payment.PaymentStatus == status {
			payments = append(payments, payment)
		}
	}
	return payments, nil
}

func TestGetPaymentsController(t *testing.T) {
	mockModel := &MockPaymentModel{
		Payments: []*models.Payment{
			{ID: 1, UserID: 1, OrderID: 1, Amount: 100.0, PaymentDate: "2023-01-01", PaymentStatus: "Completed"},
			{ID: 2, UserID: 2, OrderID: 2, Amount: 200.0, PaymentDate: "2023-01-02", PaymentStatus: "Pending"},
		},
	}
	controller := NewPaymentController(mockModel)

	req, err := http.NewRequest("GET", "/payments", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetPaymentsController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var payments []*models.Payment
	err = json.Unmarshal(rr.Body.Bytes(), &payments)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(payments))
}

func TestCreatePaymentController(t *testing.T) {
	mockModel := &MockPaymentModel{}
	controller := NewPaymentController(mockModel)

	newPayment := models.Payment{UserID: 1, OrderID: 1, Amount: 150.0, PaymentDate: "2023-02-01", PaymentStatus: "Pending"}
	paymentJson, _ := json.Marshal(newPayment)

	req, err := http.NewRequest("POST", "/payments", strings.NewReader(string(paymentJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreatePaymentController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, 1, len(mockModel.Payments))
	assert.Equal(t, newPayment.UserID, mockModel.Payments[0].UserID)
}

func TestGetPaymentByIDController(t *testing.T) {
	mockModel := &MockPaymentModel{
		Payments: []*models.Payment{
			{ID: 1, UserID: 1, OrderID: 1, Amount: 100.0, PaymentDate: "2023-01-01", PaymentStatus: "Completed"},
		},
	}
	controller := NewPaymentController(mockModel)

	req, err := http.NewRequest("GET", "/payments/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/payments/{id}", controller.GetPaymentByIDController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var payment models.Payment
	err = json.Unmarshal(rr.Body.Bytes(), &payment)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Completed", payment.PaymentStatus)
}

func TestUpdatePaymentController(t *testing.T) {
	mockModel := &MockPaymentModel{
		Payments: []*models.Payment{
			{ID: 1, UserID: 1, OrderID: 1, Amount: 100.0, PaymentDate: "2023-01-01", PaymentStatus: "Completed"},
		},
	}
	controller := NewPaymentController(mockModel)

	updatedPayment := models.Payment{ID: 1, UserID: 1, OrderID: 1, Amount: 150.0, PaymentDate: "2023-02-01", PaymentStatus: "Pending"}
	paymentJson, _ := json.Marshal(updatedPayment)

	req, err := http.NewRequest("PUT", "/payments/1", strings.NewReader(string(paymentJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/payments/{id}", controller.UpdatePaymentController).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Pending", mockModel.Payments[0].PaymentStatus)
}

func TestDeletePaymentController(t *testing.T) {
	mockModel := &MockPaymentModel{
		Payments: []*models.Payment{
			{ID: 1, UserID: 1, OrderID: 1, Amount: 100.0, PaymentDate: "2023-01-01", PaymentStatus: "Completed"},
		},
	}
	controller := NewPaymentController(mockModel)

	req, err := http.NewRequest("DELETE", "/payments/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/payments/{id}", controller.DeletePaymentController).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Equal(t, 0, len(mockModel.Payments))
}

func TestSearchPaymentController(t *testing.T) {
	mockModel := &MockPaymentModel{
		Payments: []*models.Payment{
			{ID: 1, UserID: 1, OrderID: 1, Amount: 100.0, PaymentDate: "2023-01-01", PaymentStatus: "Completed"},
			{ID: 2, UserID: 2, OrderID: 2, Amount: 200.0, PaymentDate: "2023-01-02", PaymentStatus: "Pending"},
		},
	}
	controller := NewPaymentController(mockModel)

	// Test search by order ID
	req, err := http.NewRequest("GET", "/payments/search?order_id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/payments/search", controller.SearchPaymentController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var payments []*models.Payment
	err = json.Unmarshal(rr.Body.Bytes(), &payments)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(payments))
	assert.Equal(t, 1, payments[0].OrderID)

	// Test search by user ID
	req, err = http.NewRequest("GET", "/payments/search?user_id=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	err = json.Unmarshal(rr.Body.Bytes(), &payments)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(payments))
	assert.Equal(t, 2, payments[0].UserID)

	// Test search by status
	req, err = http.NewRequest("GET", "/payments/search?status=Pending", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	err = json.Unmarshal(rr.Body.Bytes(), &payments)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(payments))
	assert.Equal(t, "Pending", payments[0].PaymentStatus)
}

//	func TestGetToken(t *testing.T) {
//		token, err := services.GetToken()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(token)
//
// }
const (
	DefaultPaymentData = `{
 "hpan":"4405639704015096","expDate":"0125","cvc":"815","terminalId":"67e34d63-102f-4bd1-898e-370781d0074d"
}`
)

func TestMakePayment(t *testing.T) {
	token, err := services.GetPaymentToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	if token == "" {
		fmt.Println("token is empty")
		return
	}
	fmt.Println(token)

	encryptedData, err := services.EncryptData(DefaultPaymentData)
	if err != nil {
		fmt.Println(err)
		return
	}
	if encryptedData == "" {
		fmt.Println("encrypted data is empty")
		return
	}
	fmt.Println(encryptedData)

	paymentResp, err := services.MakePayment(token, encryptedData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(paymentResp)
}
