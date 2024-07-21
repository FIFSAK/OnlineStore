package handlers

import (
	_ "OnlineStore/product-service/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

var urlProductsService string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	urlProductsService = os.Getenv("PRODUCT_SERVICE_URL") + "/products"
}

type InputProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
}

// @Summary Get all products
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/products [get]
// @Failure 404 {string} string "No products found"
// @Failure 500 {string} string "Internal server error"
func GetProductsHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodGet, urlProductsService, nil)
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

// @Summary Get product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [get]
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal server error"
func GetProductByIDHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodGet, urlProductsService+"/"+id, nil)
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

// @Summary Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body InputProduct true "Product object"
// @Success 201 {string} string "Product created"
// @Router /api/products [post]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func CreateProductHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodPost, urlProductsService, request.Body)
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

// @Summary Update product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body InputProduct true "Product object"
// @Success 200 {string} string "Product updated"
// @Router /api/products/{id} [put]
// @Failure 400 {string} string "Missing required fields"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal server error"
func UpdateProductHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodPut, urlProductsService+"/"+id, request.Body)
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

// @Summary Delete product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Success 204 {string} string "Product deleted"
// @Router /api/products/{id} [delete]
// @Failure 500 {string} string "Internal server error"
func DeleteProductHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodDelete, urlProductsService+"/"+id, nil)
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

// @Summary Search products
// @Tags products
// @Produce json
// @Param name query string false "Product name"
// @Param category query string false "Product category"
// @Success 200 {array} models.Product
// @Router /api/products/search [get]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func SearchProductHandler(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	resp, err := http.Get(urlProductsService + "/search?" + queryParams.Encode())
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
