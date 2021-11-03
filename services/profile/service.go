package profile

import (
	"go-profile/models"
	"go-profile/requests"

	"github.com/jinzhu/gorm"
)

type ServiceWrapper interface {
	Create(profile *models.Profile)
	Delete(profile *models.Profile)
	Update(profile *models.Profile, updateProfileRequest *requests.UpdateProfileRequest)
}

type Service struct {
	DB *gorm.DB
}

func NewProfileService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
