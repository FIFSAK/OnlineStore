package routes

import (
	"OnlineStore/api-gateway/handlers"
	_ "OnlineStore/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router = router.PathPrefix("/api").Subrouter()

	usersRouter := router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", handlers.GetUsersHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id:[0-9]+}", handlers.GetUserByIDHandler).Methods(http.MethodGet)
	usersRouter.HandleFunc("", handlers.CreateUserHandler).Methods(http.MethodPost)
	usersRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateUserHandler).Methods(http.MethodPut)
	usersRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteUserHandler).Methods(http.MethodDelete)
	usersRouter.HandleFunc("/search", handlers.SearchUserHandler).Methods(http.MethodGet)

	productsRouter := router.PathPrefix("/products").Subrouter()
	productsRouter.HandleFunc("", handlers.GetProductsHandler).Methods(http.MethodGet)
	productsRouter.HandleFunc("/{id:[0-9]+}", handlers.GetProductByIDHandler).Methods(http.MethodGet)
	productsRouter.HandleFunc("", handlers.CreateProductHandler).Methods(http.MethodPost)
	productsRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateProductHandler).Methods(http.MethodPut)
	productsRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteProductHandler).Methods(http.MethodDelete)
	productsRouter.HandleFunc("/search", handlers.SearchProductHandler).Methods(http.MethodGet)

	ordersRouter := router.PathPrefix("/orders").Subrouter()
	ordersRouter.HandleFunc("", handlers.GetOrdersHandler).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/{id:[0-9]+}", handlers.GetOrderByIDHandler).Methods(http.MethodGet)
	ordersRouter.HandleFunc("", handlers.CreateOrderHandler).Methods(http.MethodPost)
	ordersRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdateOrderHandler).Methods(http.MethodPut)
	ordersRouter.HandleFunc("/{id:[0-9]+}", handlers.DeleteOrderHandler).Methods(http.MethodDelete)
	ordersRouter.HandleFunc("/search", handlers.SearchOrderHandler).Methods(http.MethodGet)

	paymentRouter := router.PathPrefix("/payments").Subrouter()
	paymentRouter.HandleFunc("", handlers.GetPaymentsHandler).Methods(http.MethodGet)
	paymentRouter.HandleFunc("/{id:[0-9]+}", handlers.GetPaymentByIDHandler).Methods(http.MethodGet)
	paymentRouter.HandleFunc("", handlers.CreatePaymentHandler).Methods(http.MethodPost)
	paymentRouter.HandleFunc("/{id:[0-9]+}", handlers.UpdatePaymentHandler).Methods(http.MethodPut)
	paymentRouter.HandleFunc("/{id:[0-9]+}", handlers.DeletePaymentHandler).Methods(http.MethodDelete)
	paymentRouter.HandleFunc("/search", handlers.SearchPaymentHandler).Methods(http.MethodGet)
}
