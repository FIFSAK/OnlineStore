package handlers

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

var urlOrdersService = os.Getenv("ORDER_SERVICE_URL") + "/orders"

func GetOrdersHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodGet, urlOrdersService, nil)
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

func GetOrderByIDHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodGet, urlOrdersService+"/"+id, nil)
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

func CreateOrderHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodPost, urlOrdersService, request.Body)
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

func UpdateOrderHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodPut, urlOrdersService+"/"+id, request.Body)
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

func DeleteOrderHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodDelete, urlOrdersService+"/"+id, nil)
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

func SearchOrderHandler(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	resp, err := http.Get(urlOrdersService + "/search?" + queryParams.Encode())
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
