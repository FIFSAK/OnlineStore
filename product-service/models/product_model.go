package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
	DateAdded   string  `json:"date_added"`
}

type ProductModel interface {
	GetProducts() ([]*Product, error)
	CreateProduct(product Product) error
	GetProductByID(id int) (*Product, error)
	UpdateProduct(product Product) error
	DeleteProduct(id int) error
	GetProductByName(name string) ([]*Product, error)
	GetProductByCategory(category string) ([]*Product, error)
}
