package profile

import "go-profile/models"

func (profileService *Service) Create(profile *models.Profile) {
	profileService.DB.Create(profile)
}
