package service

import (
    "goapp/pkg/models"
    "goapp/pkg/repository"
)

type UserService interface {
    GetAllUsers() ([]*models.User, error)
    GetUserByID(id int) (*models.User, error)
    CreateUser(user *models.User) error
    UpdateUser(user *models.User) error
    DeleteUser(id int) error
}

type UserServiceImpl struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
    return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
    return s.repo.GetAll()
}

func (s *UserServiceImpl) GetUserByID(id int) (*models.User, error) {
    return s.repo.GetByID(id)
}

func (s *UserServiceImpl) CreateUser(user *models.User) error {
    return s.repo.Create(user)
}

func (s *UserServiceImpl) UpdateUser(user *models.User) error {
    return s.repo.Update(user)
}

func (s *UserServiceImpl) DeleteUser(id int) error {
    return s.repo.Delete(id)
}