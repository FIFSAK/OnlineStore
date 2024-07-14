package models

type User struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Address          string `json:"address"`
	RegistrationDate string `json:"registration_date"`
	Role             string `json:"role"`
}

type UserModel interface {
	GetUsers() ([]*User, error)
	CreateUser(user User) error
	GetUserByID(id int) (*User, error)
	UpdateUser(user User) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) ([]*User, error)
}
