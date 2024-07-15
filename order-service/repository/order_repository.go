package repository

import (
	"OnlineStore/order-service/models"
	"database/sql"
	"fmt"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (or *OrderRepository) GetOrders() ([]*models.Order, error) {
	rows, err := or.DB.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []*models.Order{}
	for rows.Next() {
		order := &models.Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.OrderDate, &order.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	for _, order := range orders {
		rows, err = or.DB.Query(`SELECT product_id FROM orders_products WHERE order_id = $1`, order.ID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var productID int
			if err := rows.Scan(&productID); err != nil {
				return nil, err
			}
			order.ProductIDs = append(order.ProductIDs, productID)
		}
	}

	return orders, nil
}

func (or *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
	order := &models.Order{}
	err := or.DB.QueryRow(`
        SELECT id, user_id, total_price, order_date, status
        FROM orders
        WHERE id = $1`, id).Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.OrderDate, &order.Status)
	if err != nil {
		return nil, err
	}
	rows, err := or.DB.Query(`
        SELECT product_id
        FROM orders_products
        WHERE order_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productID int
		if err := rows.Scan(&productID); err != nil {
			return nil, err
		}
		order.ProductIDs = append(order.ProductIDs, productID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return order, nil
}

func (or *OrderRepository) CreateOrder(order models.Order) error {
	tx, err := or.DB.Begin()
	if err != nil {
		return err
	}
	productsCount := make(map[int]int)
	for _, productID := range order.ProductIDs {
		productsCount[productID]++
	}
	for productID, count := range productsCount {
		var quantity int
		err := tx.QueryRow("SELECT quantity FROM products WHERE id = $1", productID).Scan(&quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
		if quantity < count {
			tx.Rollback()
			return fmt.Errorf("not enough quantity for product %d", productID)
		}
	}

	totalPrice := 0.0

	for _, productID := range order.ProductIDs {
		var price float64
		err := tx.QueryRow("SELECT price FROM products WHERE id = $1", productID).Scan(&price)
		if err != nil {
			tx.Rollback()
			return err
		}
		totalPrice += price
	}

	var orderID int
	err = tx.QueryRow(`
        INSERT INTO orders (user_id, total_price, status)
        VALUES ($1, $2, $3)
        RETURNING id`, order.UserID, totalPrice, order.Status).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, productID := range order.ProductIDs {
		_, err := tx.Exec(`
            INSERT INTO orders_products (order_id, product_id)
            VALUES ($1, $2)`, orderID, productID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) UpdateOrder(order models.Order) error {
	tx, err := or.DB.Begin()
	if err != nil {
		return err
	}
	productsCount := make(map[int]int)
	for _, productID := range order.ProductIDs {
		productsCount[productID]++
	}
	for productID, count := range productsCount {
		var quantity int
		err := tx.QueryRow("SELECT quantity FROM products WHERE id = $1", productID).Scan(&quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
		if quantity < count {
			tx.Rollback()
			return fmt.Errorf("not enough quantity for product %d", productID)
		}
	}
	
	totalPrice := 0.0

	for _, productID := range order.ProductIDs {
		var price float64
		err := tx.QueryRow("SELECT price FROM products WHERE id = $1", productID).Scan(&price)
		if err != nil {
			tx.Rollback()
			return err
		}
		totalPrice += price
	}

	_, err = tx.Exec("UPDATE orders SET user_id = $1, total_price = $2, status = $3 WHERE id = $4",
		order.UserID, totalPrice, order.Status, order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM orders_products WHERE order_id = $1", order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, productID := range order.ProductIDs {
		_, err := tx.Exec("INSERT INTO orders_products (order_id, product_id) VALUES ($1, $2)", order.ID, productID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) DeleteOrder(id int) error {
	_, err := or.DB.Exec("DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (or *OrderRepository) GetOrderByUserID(userID int) ([]*models.Order, error) {
	query := `
        SELECT o.id, o.user_id, o.total_price, o.order_date, o.status, op.product_id
        FROM orders AS o
        JOIN orders_products AS op ON o.id = op.order_id
        WHERE o.user_id = $1`

	rows, err := or.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int]*models.Order)

	for rows.Next() {
		var (
			orderID   int
			productID int
		)
		order := &models.Order{}
		err := rows.Scan(&orderID, &order.UserID, &order.TotalPrice, &order.OrderDate, &order.Status, &productID)
		if err != nil {
			return nil, err
		}

		if existingOrder, found := ordersMap[orderID]; found {
			existingOrder.ProductIDs = append(existingOrder.ProductIDs, productID)
		} else {
			order.ID = orderID
			order.ProductIDs = []int{productID}
			ordersMap[orderID] = order
		}
	}

	var orders []*models.Order
	for _, order := range ordersMap {
		orders = append(orders, order)
	}

	return orders, nil
}

func (or *OrderRepository) GetOrderByStatus(status string) ([]*models.Order, error) {
	query := `
        SELECT o.id, o.user_id, o.total_price, o.order_date, o.status, op.product_id
        FROM orders AS o
        JOIN orders_products AS op ON o.id = op.order_id
        WHERE o.status = $1`

	rows, err := or.DB.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int]*models.Order)

	for rows.Next() {
		var (
			orderID   int
			productID int
		)
		order := &models.Order{}
		err := rows.Scan(&orderID, &order.UserID, &order.TotalPrice, &order.OrderDate, &order.Status, &productID)
		if err != nil {
			return nil, err
		}

		if existingOrder, found := ordersMap[orderID]; found {
			existingOrder.ProductIDs = append(existingOrder.ProductIDs, productID)
		} else {
			order.ID = orderID
			order.ProductIDs = []int{productID}
			ordersMap[orderID] = order
		}
	}

	var orders []*models.Order
	for _, order := range ordersMap {
		orders = append(orders, order)
	}

	return orders, nil
}
