package routes

import (
	"OnlineStore/product-service/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(router *mux.Router, userController *controllers.ProductController) {
	productsRouter := router.PathPrefix("/products").Subrouter()

	productsRouter.HandleFunc("", userController.GetProductsController).Methods(http.MethodGet)
	productsRouter.HandleFunc("/{id:[0-9]+}", userController.GetProductByIDController).Methods(http.MethodGet)
	productsRouter.HandleFunc("", userController.CreateProductController).Methods(http.MethodPost)
	productsRouter.HandleFunc("/{id:[0-9]+}", userController.UpdateProductController).Methods(http.MethodPut)
	productsRouter.HandleFunc("/{id:[0-9]+}", userController.DeleteProductController).Methods(http.MethodDelete)
	productsRouter.HandleFunc("/search", userController.SearchProductController).Methods(http.MethodGet)
}
