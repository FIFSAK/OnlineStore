package controllers

import (
	"OnlineStore/payment-service/models"
	"OnlineStore/payment-service/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type PaymentController struct {
	PaymentModel models.PaymentModel
}

func NewPaymentController(paymentModel models.PaymentModel) *PaymentController {
	return &PaymentController{PaymentModel: paymentModel}
}

func (pc *PaymentController) GetPaymentsController(writer http.ResponseWriter, request *http.Request) {
	payments, err := pc.PaymentModel.GetPayments()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(payments) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	jsonPayments, err := json.Marshal(payments)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonPayments)
	return
}

func (pc *PaymentController) GetPaymentByIDController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	payment, err := pc.PaymentModel.GetPaymentByID(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if payment == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	jsonPayment, err := json.Marshal(payment)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonPayment)
	return
}

func (pc *PaymentController) CreatePaymentController(writer http.ResponseWriter, request *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(request.Body).Decode(&payment)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	paymentResponse, err := services.MakePayment()
	if err != nil {
		payment.PaymentStatus = "failed"
		fmt.Printf("Payment failed: %v", err)
	} else {
		log.Println("Payment is successful")
		payment.PaymentStatus = paymentResponse.Status
	}
	err = pc.PaymentModel.CreatePayment(payment)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	return
}

func (pc *PaymentController) UpdatePaymentController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var payment models.Payment
	err = json.NewDecoder(request.Body).Decode(&payment)
	payment.ID = id
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = pc.PaymentModel.UpdatePayment(payment)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	return
}

func (pc *PaymentController) DeletePaymentController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = pc.PaymentModel.DeletePayment(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
	return
}

func (pc *PaymentController) SearchPaymentController(writer http.ResponseWriter, request *http.Request) {
	orderID, err := strconv.Atoi(request.URL.Query().Get("order_id"))
	if err != nil && orderID != 0 {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(request.URL.Query().Get("user_id"))
	if err != nil && userID != 0 {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	status := request.URL.Query().Get("status")

	if orderID != 0 {
		payments, err := pc.PaymentModel.GetPaymentByOrderID(orderID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(payments) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonPayments, err := json.Marshal(payments)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonPayments)
		return
	}
	if userID != 0 {
		payments, err := pc.PaymentModel.GetPaymentByUserID(userID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(payments) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonPayments, err := json.Marshal(payments)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonPayments)
		return

	}
	if status != "" {
		payments, err := pc.PaymentModel.GetPaymentByStatus(status)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(payments) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonPayments, err := json.Marshal(payments)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonPayments)
		return

	}
}
