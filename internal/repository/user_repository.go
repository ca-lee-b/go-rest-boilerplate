package repository

import (
	"github.com/ca-lee-b/go-rest-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserByEmail(email string) *models.User {
	var user models.User
	u.db.Find(&user, "email = ?", email)

	return &user
}

func (u *UserRepository) Create(username string, email string, password string) error {
	result := u.db.Create(&models.User{
		Id:       uuid.NewString(),
		Email:    email,
		Username: username,
		Password: password,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
