package controllers

import (
	"OnlineStore/user-service/models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController struct {
	UserModel models.UserModel
}

func NewUserController(userModel models.UserModel) *UserController {
	return &UserController{UserModel: userModel}
}

func (uc *UserController) GetUsersController(writer http.ResponseWriter, request *http.Request) {
	users, err := uc.UserModel.GetUsers()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonUsers)
}

func (uc *UserController) GetUserByIDController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := uc.UserModel.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonUser)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (uc *UserController) CreateUserController(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = uc.UserModel.CreateUser(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (uc *UserController) UpdateUserController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = id
	err = uc.UserModel.UpdateUser(user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (uc *UserController) DeleteUserController(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	err = uc.UserModel.DeleteUser(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (uc *UserController) SearchUserController(writer http.ResponseWriter, request *http.Request) {
	email := request.URL.Query().Get("email")
	name := request.URL.Query().Get("name")
	if email != "" {
		user, err := uc.UserModel.GetUserByEmail(email)
		if err != nil {
			if err == sql.ErrNoRows {
				writer.WriteHeader(http.StatusNotFound)
				return
			}
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if user == nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonUser)
	}
	if name != "" {
		username := request.URL.Query().Get("name")
		users, err := uc.UserModel.GetUserByUsername(username)
		if err != nil {
			if err == sql.ErrNoRows {
				writer.WriteHeader(http.StatusNotFound)
				return
			}
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(users) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		jsonUsers, err := json.Marshal(users)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err = writer.Write(jsonUsers)
	}

}
