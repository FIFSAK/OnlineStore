package handlers

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

var urlOrdersService string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	urlOrdersService = os.Getenv("ORDER_SERVICE_URL") + "/orders"
}

type InputOrder struct {
	UserID     int    `json:"user_id"`
	Status     string `json:"status"`
	ProductIDs []int  `json:"product_ids"`
}

// @Summary Get all orders
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order
// @Router /api/orders [get]
// @Failure 404 {string} string "No orders found"
// @Failure 500 {string} string "Internal server error"
func GetOrdersHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodGet, urlOrdersService, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(resp.StatusCode)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Get order by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /api/orders/{id} [get]
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal server error"
func GetOrderByIDHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodGet, urlOrdersService+"/"+id, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(resp.StatusCode)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body InputOrder true "Order object"
// @Success 201 {string} string "Order created"
// @Router /api/orders [post]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func CreateOrderHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodPost, urlOrdersService, request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(resp.StatusCode)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Update order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body InputOrder true "Order object"
// @Success 200 {string} string "Order updated"
// @Router /api/orders/{id} [put]
// @Failure 400 {string} string "Missing required fields"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal server error"
func UpdateOrderHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodPut, urlOrdersService+"/"+id, request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(resp.StatusCode)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Delete order by ID
// @Tags orders
// @Param id path int true "Order ID"
// @Success 204 {string} string "Order deleted"
// @Router /api/orders/{id} [delete]
// @Failure 500 {string} string "Internal server error"
func DeleteOrderHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodDelete, urlOrdersService+"/"+id, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(resp.StatusCode)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Search orders
// @Tags orders
// @Produce json
// @Param user_id query int false "User ID"
// @Param status query string false "Order status"
// @Success 200 {array} models.Order
// @Router /api/orders/search [get]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func SearchOrderHandler(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	resp, err := http.Get(urlOrdersService + "/search?" + queryParams.Encode())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(resp.StatusCode)
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
