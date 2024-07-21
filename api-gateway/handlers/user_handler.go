package handlers

import (
	_ "OnlineStore/user-service/models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

var urlUsersService string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	urlUsersService = os.Getenv("USER_SERVICE_URL") + "/users"
}

type InputUser struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users [get]
// @Failure 404 {string} string "No users found"
// @Failure 500 {string} string "Internal server error"
func GetUsersHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(urlUsersService)
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

// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /api/users/{id} [get]
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
func GetUserByIDHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
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

// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body InputUser true "User object"
// @Success 201 {string} string "User created"
// @Router /api/users [post]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body InputUser true "User object"
// @Success 200 {string} string "User updated"
// @Router /api/users/{id} [put]
// @Failure 400 {string} string "Missing required fields"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
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

// @Summary Delete user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 204 {string} string "User deleted"
// @Router /api/users/{id} [delete]
// @Failure 500 {string} string "Internal server error"
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

// @Summary Search user
// @Tags users
// @Produce json
// @Param name query string false "User name"
// @Param email query string false "User email"
// @Param role query string false "User role"
// @Success 200 {array} models.User
// @Router /api/users/search [get]
// @Failure 400 {string} string "Missing required fields"
// @Failure 500 {string} string "Internal server error"
func SearchUserHandler(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	resp, err := http.Get(urlUsersService + "/search?" + queryParams.Encode())
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
