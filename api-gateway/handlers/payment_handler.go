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

var urlPaymentService = os.Getenv("PAYMENT_SERVICE_URL") + "/payments"

func GetPaymentsHandler(writer http.ResponseWriter, request *http.Request) {
	resp, err := http.Get(urlPaymentService)
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

func GetPaymentByIDHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	resp, err := http.Get(urlPaymentService + "/" + id)
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

func CreatePaymentHandler(writer http.ResponseWriter, request *http.Request) {
	resp, err := http.Post(urlPaymentService, "application/json", request.Body)
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

func UpdatePaymentHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, urlPaymentService+"/"+id, request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

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

func DeletePaymentHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, urlPaymentService+"/"+id, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

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

func SearchPaymentHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	resp, err := http.Get(urlPaymentService + "/search?" + query.Encode())
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
