package handlers

import (
	_ "OnlineStore/payment-service/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

var urlPaymentService string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	urlPaymentService = os.Getenv("PAYMENT_SERVICE_URL") + "/payments"
}

type InputPayment struct {
	UserID  int     `json:"user_id"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
}

// @Summary Get all payments
// @Tags payments
// @Produce json
// @Success 200 {array} models.Payment
// @Router /api/payments [get]
// @Failure 404 {string} string "No payments found"
// @Failure 500 {string} string "Internal server error"
func GetPaymentsHandler(writer http.ResponseWriter, request *http.Request) {
	resp, err := http.Get(urlPaymentService)
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

// @Summary Get payment by ID
// @Tags payments
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} models.Payment
// @Router /api/payments/{id} [get]
// @Failure 404 {string} string "Payment not found"
// @Failure 500 {string} string "Internal server error"
func GetPaymentByIDHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	resp, err := http.Get(urlPaymentService + "/" + id)
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

// @Summary Create a new payment
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body InputPayment true "Payment object"
// @Success 201 {string} string "Payment created"
// @Router /api/payments [post]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func CreatePaymentHandler(writer http.ResponseWriter, request *http.Request) {
	resp, err := http.Post(urlPaymentService, "application/json", request.Body)
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

// @Summary Update payment by ID
// @Tags payments
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Param payment body InputPayment true "Payment object"
// @Success 200 {string} string "Payment updated"
// @Router /api/payments/{id} [put]
// @Failure 400 {string} string "Missing required fields"
// @Failure 404 {string} string "Payment not found"
// @Failure 500 {string} string "Internal server error"
func UpdatePaymentHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, urlPaymentService+"/"+id, request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

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

// @Summary Delete payment by ID
// @Tags payments
// @Param id path int true "Payment ID"
// @Success 204 {string} string "Payment deleted"
// @Router /api/payments/{id} [delete]
// @Failure 500 {string} string "Internal server error"
func DeletePaymentHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, urlPaymentService+"/"+id, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

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

// @Summary Search payments
// @Tags payments
// @Produce json
// @Param order_id query int false "Order ID"
// @Param user_id query int false "User ID"
// @Param status query string false "Payment status"
// @Success 200 {array} models.Payment
// @Router /api/payments/search [get]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func SearchPaymentHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	resp, err := http.Get(urlPaymentService + "/search?" + query.Encode())
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
