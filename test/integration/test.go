package handler_test

import (
	"goapp/api/handler"
	"goapp/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserService struct{}

// CreateUser implements service.UserService.
func (*MockUserService) CreateUser(user *models.User) error {
	panic("unimplemented")
}

// DeleteUser implements service.UserService.
func (*MockUserService) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetAllUsers implements service.UserService.
func (*MockUserService) GetAllUsers() ([]*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements service.UserService.
func (*MockUserService) UpdateUser(user *models.User) error {
	panic("unimplemented")
}

func (m *MockUserService) GetUserByID(id int) (*models.User, error) {
	// For simplicity, return a static user.
	return &models.User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}

func TestGetUserByIDHandler(t *testing.T) {
	userHandler := handler.NewUserHandler(&MockUserService{})

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetUserByID)

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", rec.Code)
	}
}
