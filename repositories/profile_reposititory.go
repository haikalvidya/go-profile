package repositories

import (
	"go-profile/models"

	"github.com/jinzhu/gorm"
)

type ProfileRepositoryQ interface {
	GetProfiles(profiles *[]models.Profile)
	GetProfile(profile *models.Profile, id int)
	GetProfileByUserId(profile *models.Profile, id uint)
}

type ProfileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (profileRepository *ProfileRepository) GetProfiles(profiles *[]models.Profile) {
	profileRepository.DB.Find(profiles)
}

func (profileRepository *ProfileRepository) GetProfile(profile *models.Profile, id int) {
	profileRepository.DB.Where("id = ? ", id).Find(profile)
}

func (profileRepository *ProfileRepository) GetProfileByUserId(profile *models.Profile, id uint) {
	profileRepository.DB.Where("user_id = ?", id).Find(profile)
}
