package profile

import (
	"go-profile/models"
	"go-profile/requests"
)

func (profileService *Service) Update(profile *models.Profile, updateProfileRequest *requests.UpdateProfileRequest) {
	profile.Alamat = updateProfileRequest.Alamat
	profile.Pekerjaan = updateProfileRequest.Pekerjaan
	profile.NamaLengkap = updateProfileRequest.NamaLengkap
	profile.Pendidikan = updateProfileRequest.Pendidikan
	profile.NomorTelpon = updateProfileRequest.NomorTelpon
	profileService.DB.Save(profile)
}
