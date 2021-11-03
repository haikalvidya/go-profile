package profile

import "go-profile/models"

func (profileService *Service) Delete(profile *models.Profile) {
	profileService.DB.Delete(profile)
}
