package routes

import (
	"OnlineStore/order-service/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(router *mux.Router, orderController *controllers.OrderController) {
	ordersRouter := router.PathPrefix("/orders").Subrouter()

	ordersRouter.HandleFunc("", orderController.GetOrdersController).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/{id:[0-9]+}", orderController.GetOrderByIDController).Methods(http.MethodGet)
	ordersRouter.HandleFunc("", orderController.CreateOrderController).Methods(http.MethodPost)
	ordersRouter.HandleFunc("/{id:[0-9]+}", orderController.UpdateOrderController).Methods(http.MethodPut)
	ordersRouter.HandleFunc("/{id:[0-9]+}", orderController.DeleteOrderController).Methods(http.MethodDelete)
	ordersRouter.HandleFunc("/search", orderController.SearchOrderController).Methods(http.MethodGet)
}
