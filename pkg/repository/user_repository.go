package repository

import (
    "database/sql"
    "goapp/pkg/models"
)

type UserRepository interface {
    GetAll() ([]*models.User, error)
    GetByID(id int) (*models.User, error)
    Create(user *models.User) error
    Update(user *models.User) error
    Delete(id int) error
}

type UserRepositoryImpl struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetAll() ([]*models.User, error) {
    rows, err := r.db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*models.User
    for rows.Next() {
        user := &models.User{}
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (r *UserRepositoryImpl) GetByID(id int) (*models.User, error) {
    user := &models.User{}
    if err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
        return nil, err
    }

    return user, nil
}

func (r *UserRepositoryImpl) Create(user *models.User) error {
    return r.db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
    _, err := r.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
    return err
}

func (r *UserRepositoryImpl) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
    return err
}