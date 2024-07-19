package routes

import (
	"OnlineStore/payment-service/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(router *mux.Router, paymentController *controllers.PaymentController) {
	paymentsRouter := router.PathPrefix("/payments").Subrouter()

	paymentsRouter.HandleFunc("", paymentController.GetPaymentsController).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("/{id:[0-9]+}", paymentController.GetPaymentByIDController).Methods(http.MethodGet)
	paymentsRouter.HandleFunc("", paymentController.CreatePaymentController).Methods(http.MethodPost)
	paymentsRouter.HandleFunc("/{id:[0-9]+}", paymentController.UpdatePaymentController).Methods(http.MethodPut)
	paymentsRouter.HandleFunc("/{id:[0-9]+}", paymentController.DeletePaymentController).Methods(http.MethodDelete)
	paymentsRouter.HandleFunc("/search", paymentController.SearchPaymentController).Methods(http.MethodGet)
}
