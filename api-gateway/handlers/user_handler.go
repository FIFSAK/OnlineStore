package handlers

import (
	"fmt"
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

var (
	urlUsersService = os.Getenv("USER_SERVICE_URL") + "/users"
)

func GetUsersHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodGet, urlUsersService, nil)
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
}

func GetUserByIDHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println(id)
	req, err := http.NewRequest(http.MethodGet, urlUsersService+"/"+id, nil)
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
}

func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	req, err := http.NewRequest(http.MethodPost, urlUsersService, request.Body)
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
}

func UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodPut, urlUsersService+"/"+id, request.Body)
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
}

func DeleteUserHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	req, err := http.NewRequest(http.MethodDelete, urlUsersService+"/"+id, nil)
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
}

func SearchUserHandler(writer http.ResponseWriter, request *http.Request) {
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
	}
}
