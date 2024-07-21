package controllers

import (
	"OnlineStore/order-service/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"database/sql"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// MockOrderModel is a mock implementation of the OrderModel interface
type MockOrderModel struct {
	Orders []*models.Order
}

func (m *MockOrderModel) GetOrders() ([]*models.Order, error) {
	return m.Orders, nil
}

func (m *MockOrderModel) CreateOrder(order models.Order) error {
	m.Orders = append(m.Orders, &order)
	return nil
}

func (m *MockOrderModel) GetOrderByID(id int) (*models.Order, error) {
	for _, order := range m.Orders {
		if order.ID == id {
			return order, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (m *MockOrderModel) UpdateOrder(order models.Order) error {
	for i, o := range m.Orders {
		if o.ID == order.ID {
			m.Orders[i] = &order
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockOrderModel) DeleteOrder(id int) error {
	for i, order := range m.Orders {
		if order.ID == id {
			m.Orders = append(m.Orders[:i], m.Orders[i+1:]...)
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockOrderModel) GetOrderByUserID(userID int) ([]*models.Order, error) {
	var orders []*models.Order
	for _, order := range m.Orders {
		if order.UserID == userID {
			orders = append(orders, order)
		}
	}
	return orders, nil
}

func (m *MockOrderModel) GetOrderByStatus(status string) ([]*models.Order, error) {
	var orders []*models.Order
	for _, order := range m.Orders {
		if order.Status == status {
			orders = append(orders, order)
		}
	}
	return orders, nil
}

func TestGetOrdersController(t *testing.T) {
	mockModel := &MockOrderModel{
		Orders: []*models.Order{
			{ID: 1, UserID: 1, TotalPrice: 100.0, OrderDate: "2023-01-01", Status: "New", ProductIDs: []int{1, 2}},
			{ID: 2, UserID: 2, TotalPrice: 200.0, OrderDate: "2023-01-02", Status: "Shipped", ProductIDs: []int{3, 4}},
		},
	}
	controller := NewOrderController(mockModel)

	req, err := http.NewRequest("GET", "/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetOrdersController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var orders []*models.Order
	err = json.Unmarshal(rr.Body.Bytes(), &orders)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(orders))
}

func TestCreateOrderController(t *testing.T) {
	mockModel := &MockOrderModel{}
	controller := NewOrderController(mockModel)

	newOrder := models.Order{UserID: 1, TotalPrice: 150.0, OrderDate: "2023-02-01", Status: "New", ProductIDs: []int{1, 2, 3}}
	orderJson, _ := json.Marshal(newOrder)

	req, err := http.NewRequest("POST", "/orders", strings.NewReader(string(orderJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateOrderController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, 1, len(mockModel.Orders))
	assert.Equal(t, newOrder.UserID, mockModel.Orders[0].UserID)
}

func TestGetOrderByIDController(t *testing.T) {
	mockModel := &MockOrderModel{
		Orders: []*models.Order{
			{ID: 1, UserID: 1, TotalPrice: 100.0, OrderDate: "2023-01-01", Status: "New", ProductIDs: []int{1, 2}},
		},
	}
	controller := NewOrderController(mockModel)

	req, err := http.NewRequest("GET", "/orders/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/orders/{id}", controller.GetOrderByIDController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var order models.Order
	err = json.Unmarshal(rr.Body.Bytes(), &order)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "New", order.Status)
}

func TestUpdateOrderController(t *testing.T) {
	mockModel := &MockOrderModel{
		Orders: []*models.Order{
			{ID: 1, UserID: 1, TotalPrice: 100.0, OrderDate: "2023-01-01", Status: "New", ProductIDs: []int{1, 2}},
		},
	}
	controller := NewOrderController(mockModel)

	updatedOrder := models.Order{ID: 1, UserID: 1, TotalPrice: 150.0, OrderDate: "2023-02-01", Status: "Shipped", ProductIDs: []int{1, 2, 3}}
	orderJson, _ := json.Marshal(updatedOrder)

	req, err := http.NewRequest("PUT", "/orders/1", strings.NewReader(string(orderJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/orders/{id}", controller.UpdateOrderController).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Shipped", mockModel.Orders[0].Status)
}

func TestDeleteOrderController(t *testing.T) {
	mockModel := &MockOrderModel{
		Orders: []*models.Order{
			{ID: 1, UserID: 1, TotalPrice: 100.0, OrderDate: "2023-01-01", Status: "New", ProductIDs: []int{1, 2}},
		},
	}
	controller := NewOrderController(mockModel)

	req, err := http.NewRequest("DELETE", "/orders/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/orders/{id}", controller.DeleteOrderController).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, 0, len(mockModel.Orders))
}

func TestSearchOrderController(t *testing.T) {
	mockModel := &MockOrderModel{
		Orders: []*models.Order{
			{ID: 1, UserID: 1, TotalPrice: 100.0, OrderDate: "2023-01-01", Status: "New", ProductIDs: []int{1, 2}},
			{ID: 2, UserID: 2, TotalPrice: 200.0, OrderDate: "2023-01-02", Status: "Shipped", ProductIDs: []int{3, 4}},
		},
	}
	controller := NewOrderController(mockModel)

	// Test search by user ID
	req, err := http.NewRequest("GET", "/orders/search?user=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/orders/search", controller.SearchOrderController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var orders []*models.Order
	err = json.Unmarshal(rr.Body.Bytes(), &orders)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(orders))
	assert.Equal(t, 1, orders[0].UserID)

	// Test search by status
	req, err = http.NewRequest("GET", "/orders/search?status=Shipped", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	err = json.Unmarshal(rr.Body.Bytes(), &orders)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(orders))
	assert.Equal(t, "Shipped", orders[0].Status)
}
