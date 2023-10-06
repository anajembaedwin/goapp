package service_test

import (
	"goapp/pkg/models"
	"goapp/pkg/service"
	"testing"
)

type MockUserRepository struct{}

// Create implements repository.UserRepository.
func (*MockUserRepository) Create(user *models.User) error {
	panic("unimplemented")
}

// Delete implements repository.UserRepository.
func (*MockUserRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements repository.UserRepository.
func (*MockUserRepository) GetAll() ([]*models.User, error) {
	panic("unimplemented")
}

// GetByID implements repository.UserRepository.
func (*MockUserRepository) GetByID(id int) (*models.User, error) {
	panic("unimplemented")
}

// Update implements repository.UserRepository.
func (*MockUserRepository) Update(user *models.User) error {
	panic("unimplemented")
}

func (m *MockUserRepository) GetUserByID(id int) (*models.User, error) {
	// For simplicity, return a static user.
	return &models.User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}

func TestGetUserByID(t *testing.T) {
	userService := service.NewUserService(&MockUserRepository{})

	user, err := userService.GetUserByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected ID to be 1, got %v", user.ID)
	}
}
