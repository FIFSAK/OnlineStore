package controllers

import (
	"OnlineStore/product-service/models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductModel models.ProductModel
}

func NewProductController(userModel models.ProductModel) *ProductController {
	return &ProductController{ProductModel: userModel}
}

func (pc *ProductController) GetProductsController(writer http.ResponseWriter, request *http.Request) {
	products, err := pc.ProductModel.GetProducts()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(products) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	jsonProducts, err := json.Marshal(products)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonProducts)
}

func (pc *ProductController) GetProductByIDController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := pc.ProductModel.GetProductByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonProduct, err := json.Marshal(product)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonProduct)
}

func (pc *ProductController) CreateProductController(writer http.ResponseWriter, request *http.Request) {
	var product models.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.ProductModel.CreateProduct(product)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (pc *ProductController) UpdateProductController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var product models.Product
	err = json.NewDecoder(request.Body).Decode(&product)
	product.ID = id
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.ProductModel.UpdateProduct(product)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (pc *ProductController) DeleteProductController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.ProductModel.DeleteProduct(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (pc *ProductController) SearchProductController(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	category := request.URL.Query().Get("category")
	if name != "" {
		products, err := pc.ProductModel.GetProductByName(name)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(products) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonProducts, err := json.Marshal(products)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonProducts)
	} else if category != "" {
		products, err := pc.ProductModel.GetProductByCategory(category)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(products) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonProducts, err := json.Marshal(products)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonProducts)
	} else {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}
}
