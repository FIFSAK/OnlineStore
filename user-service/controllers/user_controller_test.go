package controllers

import (
	"OnlineStore/user-service/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// MockUserModel is a mock implementation of the UserModel interface
type MockUserModel struct {
	Users []*models.User
}

func (m *MockUserModel) GetUsers() ([]*models.User, error) {
	return m.Users, nil
}

func (m *MockUserModel) CreateUser(user models.User) error {
	m.Users = append(m.Users, &user)
	return nil
}

func (m *MockUserModel) GetUserByID(id int) (*models.User, error) {
	for _, user := range m.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (m *MockUserModel) UpdateUser(user models.User) error {
	for i, u := range m.Users {
		if u.ID == user.ID {
			m.Users[i] = &user
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockUserModel) DeleteUser(id int) error {
	for i, user := range m.Users {
		if user.ID == id {
			m.Users = append(m.Users[:i], m.Users[i+1:]...)
			return nil
		}
	}
	return sql.ErrNoRows
}

func (m *MockUserModel) GetUserByEmail(email string) (*models.User, error) {
	for _, user := range m.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (m *MockUserModel) GetUserByUsername(username string) ([]*models.User, error) {
	var users []*models.User
	for _, user := range m.Users {
		if user.Username == username {
			users = append(users, user)
		}
	}
	return users, nil
}

func TestGetUsersController(t *testing.T) {
	mockModel := &MockUserModel{
		Users: []*models.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
			{ID: 2, Username: "user2", Email: "user2@example.com"},
		},
	}
	controller := NewUserController(mockModel)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetUsersController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var users []*models.User
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(users))
}

func TestCreateUserController(t *testing.T) {
	mockModel := &MockUserModel{}
	controller := NewUserController(mockModel)

	newUser := models.User{Username: "newuser", Email: "newuser@example.com", Address: "123 Street", Role: "user"}
	userJson, _ := json.Marshal(newUser)

	req, err := http.NewRequest("POST", "/users", strings.NewReader(string(userJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateUserController)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, 1, len(mockModel.Users))
	assert.Equal(t, newUser.Username, mockModel.Users[0].Username)
}

func TestGetUserByIDController(t *testing.T) {
	mockModel := &MockUserModel{
		Users: []*models.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
		},
	}
	controller := NewUserController(mockModel)

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", controller.GetUserByIDController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var user models.User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "user1", user.Username)
}

func TestUpdateUserController(t *testing.T) {
	mockModel := &MockUserModel{
		Users: []*models.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
		},
	}
	controller := NewUserController(mockModel)

	updatedUser := models.User{ID: 1, Username: "updateduser", Email: "user1@example.com", Address: "123 Street", Role: "admin"}
	userJson, _ := json.Marshal(updatedUser)

	req, err := http.NewRequest("PUT", "/users/1", strings.NewReader(string(userJson)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", controller.UpdateUserController).Methods("PUT")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "updateduser", mockModel.Users[0].Username)
}

func TestDeleteUserController(t *testing.T) {
	mockModel := &MockUserModel{
		Users: []*models.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
		},
	}
	controller := NewUserController(mockModel)

	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", controller.DeleteUserController).Methods("DELETE")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, 0, len(mockModel.Users))
}

func TestSearchUserController(t *testing.T) {
	mockModel := &MockUserModel{
		Users: []*models.User{
			{ID: 1, Username: "user1", Email: "user1@example.com"},
			{ID: 2, Username: "user2", Email: "user2@example.com"},
		},
	}
	controller := NewUserController(mockModel)

	// Test search by email
	req, err := http.NewRequest("GET", "/users/search?email=user1@example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/search", controller.SearchUserController).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var user models.User
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "user1", user.Username)

	// Test search by username
	req, err = http.NewRequest("GET", "/users/search?name=user2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var users []*models.User
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(users))
	assert.Equal(t, "user2", users[0].Username)
}
