package repository

import (
	"OnlineStore/user-service/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) GetUsers() ([]*models.User, error) {
	rows, err := ur.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Address, &user.RegistrationDate, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	err := ur.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Address, &user.RegistrationDate, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) CreateUser(user models.User) error {
	_, err := ur.DB.Exec("INSERT INTO users (username, email, address, role) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.Address, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdateUser(user models.User) error {
	_, err := ur.DB.Exec("UPDATE users SET username = $1, email = $2, address = $3, role = $4 WHERE id = $5", user.Username, user.Email, user.Address, user.Role, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByUsername(username string) ([]*models.User, error) {
	var users []*models.User
	rows, err := ur.DB.Query("SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Address, &user.RegistrationDate, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := ur.DB.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Address, &user.RegistrationDate, &user.Role)
	if err != nil {
		return nil, err
	}

	return user, nil
}
