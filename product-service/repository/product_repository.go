package repository

import (
	"OnlineStore/product-service/models"
	"database/sql"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) GetProducts() ([]*models.Product, error) {
	rows, err := pr.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*models.Product{}
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Quantity, &product.DateAdded)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	product := &models.Product{}
	err := pr.DB.QueryRow("SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Quantity, &product.DateAdded)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepository) CreateProduct(product models.Product) error {
	_, err := pr.DB.Exec("INSERT INTO products (name, description, price, category, quantity) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Price, product.Category, product.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) UpdateProduct(product models.Product) error {
	_, err := pr.DB.Exec("UPDATE products SET name = $1, description = $2, price = $3, category = $4, quantity = $5 WHERE id = $6", product.Name, product.Description, product.Price, product.Category, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	_, err := pr.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) GetProductByName(name string) ([]*models.Product, error) {
	var products []*models.Product
	rows, err := pr.DB.Query("SELECT * FROM products WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Quantity, &product.DateAdded)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) GetProductByCategory(category string) ([]*models.Product, error) {
	var products []*models.Product
	rows, err := pr.DB.Query("SELECT * FROM products WHERE category = $1", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Quantity, &product.DateAdded)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
