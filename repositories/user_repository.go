package repositories

import (
	"go-profile/models"

	"github.com/jinzhu/gorm"
)

type UserRepositoryQ interface {
	GetUsers(user *[]models.User)
	GetUserByEmail(user *models.User, email string)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) GetUsers(user *[]models.User) {
	userRepository.DB.Find(user)
}

func (userRepository *UserRepository) GetUserByEmail(user *models.User, email string) {
	userRepository.DB.Where("email = ?", email).Find(user)
}
