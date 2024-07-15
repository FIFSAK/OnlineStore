package controllers

import (
	"OnlineStore/order-service/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type OrderController struct {
	OrderModel models.OrderModel
}

func NewOrderController(orderModel models.OrderModel) *OrderController {
	return &OrderController{OrderModel: orderModel}
}

func (oc *OrderController) GetOrdersController(writer http.ResponseWriter, request *http.Request) {
	orders, err := oc.OrderModel.GetOrders()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(orders) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	jsonOrders, err := json.Marshal(orders)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonOrders)
}

func (oc *OrderController) GetOrderByIDController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := oc.OrderModel.GetOrderByID(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonOrder)
}

func (oc *OrderController) CreateOrderController(writer http.ResponseWriter, request *http.Request) {
	var order models.Order
	err := json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = oc.OrderModel.CreateOrder(order)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (oc *OrderController) UpdateOrderController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return

	}
	var order models.Order
	err = json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	order.ID = id
	err = oc.OrderModel.UpdateOrder(order)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (oc *OrderController) DeleteOrderController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = oc.OrderModel.DeleteOrder(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (oc *OrderController) SearchOrderController(writer http.ResponseWriter, request *http.Request) {
	userID := request.URL.Query().Get("user")
	status := request.URL.Query().Get("status")
	if userID != "" {
		userIdInt, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		orders, err := oc.OrderModel.GetOrderByUserID(userIdInt)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(orders) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonOrders, err := json.Marshal(orders)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonOrders)
	} else if status != "" {
		orders, err := oc.OrderModel.GetOrderByStatus(status)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(orders) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonOrders, err := json.Marshal(orders)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonOrders)
	} else {
		http.Error(writer, "Invalid query parameters", http.StatusBadRequest)
	}

}
