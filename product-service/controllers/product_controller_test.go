package controllers

import (
	"OnlineStore/product-service/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"database/sql"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// MockProductModel is a mock implementation of the ProductModel interface
type MockProductModel struct {
	Products []*models.Product
}

func (m *MockProductModel) GetProducts() ([]*models.Product, error) {
	return m.Products, nil
}

func (m *MockProductModel) CreateProduct(product models.Product) error {
	m.Products = append(m.Products, &product)
	return nil
}

func (m *MockProductModel) GetProductByID(id int) (*models.Product, error) {
	for _, product := range m.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (m *MockProductModel) UpdateProduct(product models.Product) error {
	for i, p := range m.Products {
		if p.ID == product.ID {
			m.Products[i] = &product
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockProductModel) DeleteProduct(id int) error {
	for i, product := range m.Products {
		if product.ID == id {
			m.Products = append(m.Products[:i], m.Products[i+1:]...)
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockProductModel) GetProductByName(name string) ([]*models.Product, error) {
	var products []*models.Product
	for _, product := range m.Products {
		if product.Name == name {
			products = append(products, product)
		}
	}
	return products, nil
}

func (m *MockProductModel) GetProductByCategory(category string) ([]*models.Product, error) {
	var products []*models.Product
	for _, product := range m.Products {
		if product.Category == category {
			products = append(products, product)
		}
	}
	return products, nil
}

func TestGetProductsController(t *testing.T) {
	mockModel := &MockProductModel{
		Products: []*models.Product{
			{ID: 1, Name: "Product1", Description: "Description1", Price: 10.0, Category: "Category1", Quantity: 100},
			{ID: 2, Name: "Product2", Description: "Description2", Price: 20.0, Category: "Category2", Quantity: 200},
		},
	}
	controller := NewProductController(mockModel)

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetProductsController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var products []*models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &products)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(products))
}

func TestCreateProductController(t *testing.T) {
	mockModel := &MockProductModel{}
	controller := NewProductController(mockModel)

	newProduct := models.Product{Name: "NewProduct", Description: "NewDescription", Price: 30.0, Category: "NewCategory", Quantity: 300}
	productJson, _ := json.Marshal(newProduct)

	req, err := http.NewRequest("POST", "/products", strings.NewReader(string(productJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateProductController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, 1, len(mockModel.Products))
	assert.Equal(t, newProduct.Name, mockModel.Products[0].Name)
}

func TestGetProductByIDController(t *testing.T) {
	mockModel := &MockProductModel{
		Products: []*models.Product{
			{ID: 1, Name: "Product1", Description: "Description1", Price: 10.0, Category: "Category1", Quantity: 100},
		},
	}
	controller := NewProductController(mockModel)

	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", controller.GetProductByIDController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var product models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &product)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Product1", product.Name)
}

func TestUpdateProductController(t *testing.T) {
	mockModel := &MockProductModel{
		Products: []*models.Product{
			{ID: 1, Name: "Product1", Description: "Description1", Price: 10.0, Category: "Category1", Quantity: 100},
		},
	}
	controller := NewProductController(mockModel)

	updatedProduct := models.Product{ID: 1, Name: "UpdatedProduct", Description: "UpdatedDescription", Price: 15.0, Category: "UpdatedCategory", Quantity: 150}
	productJson, _ := json.Marshal(updatedProduct)

	req, err := http.NewRequest("PUT", "/products/1", strings.NewReader(string(productJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", controller.UpdateProductController).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "UpdatedProduct", mockModel.Products[0].Name)
}

func TestDeleteProductController(t *testing.T) {
	mockModel := &MockProductModel{
		Products: []*models.Product{
			{ID: 1, Name: "Product1", Description: "Description1", Price: 10.0, Category: "Category1", Quantity: 100},
		},
	}
	controller := NewProductController(mockModel)

	req, err := http.NewRequest("DELETE", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", controller.DeleteProductController).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, 0, len(mockModel.Products))
}

func TestSearchProductController(t *testing.T) {
	mockModel := &MockProductModel{
		Products: []*models.Product{
			{ID: 1, Name: "Product1", Description: "Description1", Price: 10.0, Category: "Category1", Quantity: 100},
			{ID: 2, Name: "Product2", Description: "Description2", Price: 20.0, Category: "Category2", Quantity: 200},
		},
	}
	controller := NewProductController(mockModel)

	// Test search by name
	req, err := http.NewRequest("GET", "/products/search?name=Product1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/search", controller.SearchProductController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var products []*models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &products)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(products))
	assert.Equal(t, "Product1", products[0].Name)

	// Test search by category
	req, err = http.NewRequest("GET", "/products/search?category=Category2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	err = json.Unmarshal(rr.Body.Bytes(), &products)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(products))
	assert.Equal(t, "Product2", products[0].Name)
}
