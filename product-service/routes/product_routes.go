package routes

import (
	"OnlineStore/product-service/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(router *mux.Router, productController *controllers.ProductController) {
	productsRouter := router.PathPrefix("/products").Subrouter()

	productsRouter.HandleFunc("", productController.GetProductsController).Methods(http.MethodGet)
	productsRouter.HandleFunc("/{id:[0-9]+}", productController.GetProductByIDController).Methods(http.MethodGet)
	productsRouter.HandleFunc("", productController.CreateProductController).Methods(http.MethodPost)
	productsRouter.HandleFunc("/{id:[0-9]+}", productController.UpdateProductController).Methods(http.MethodPut)
	productsRouter.HandleFunc("/{id:[0-9]+}", productController.DeleteProductController).Methods(http.MethodDelete)
	productsRouter.HandleFunc("/search", productController.SearchProductController).Methods(http.MethodGet)
}
