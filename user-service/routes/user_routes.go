package routes

import (
	"OnlineStore/user-service/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(router *mux.Router, userController *controllers.UserController) {
	usersRouter := router.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("", userController.GetUsersController).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id:[0-9]+}", userController.GetUserByIDController).Methods(http.MethodGet)
	usersRouter.HandleFunc("", userController.CreateUserController).Methods(http.MethodPost)
	usersRouter.HandleFunc("/{id:[0-9]+}", userController.UpdateUserController).Methods(http.MethodPut)
	usersRouter.HandleFunc("/{id:[0-9]+}", userController.DeleteUserController).Methods(http.MethodDelete)
	usersRouter.HandleFunc("/search", userController.SearchUserController).Methods(http.MethodGet)
}
