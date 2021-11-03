package responses

import (
	"go-profile/models"
)

type ProfileResponse struct {
	ID          uint   `json:"id" example:"1"`
	Username    string `json:"username" example:"John"`
	Email       string `json:"email" example:"John@doo.com"`
	Alamat      string `json:"alamat" example:"Jl. Echo"`
	Pekerjaan   string `json:"pekerjaan" example:"DevOps"`
	NamaLengkap string `json:"nama_lengkap" example:"Echo Parecho"`
	Pendidikan  string `json:"pendidikan_terakhir" example:"Mahasiswa"`
	NomorTelpon string `json:"nomer_telpon" example:"+6281234123421"`
}

func NewProfileResponse(profiles []models.Profile, users []models.User) *[]ProfileResponse {
	profileResponse := make([]ProfileResponse, 0)
	var username, email string

	for i := range profiles {
		for j := range users {
			if profiles[i].UserID == users[j].ID {
				username = users[j].Name
				email = users[j].Email
			} else {
				username = ""
				email = ""
			}
		}
		profileResponse = append(profileResponse, ProfileResponse{
			ID:          profiles[i].ID,
			Username:    username,
			Email:       email,
			Alamat:      profiles[i].Alamat,
			Pekerjaan:   profiles[i].Pekerjaan,
			NamaLengkap: profiles[i].NamaLengkap,
			Pendidikan:  profiles[i].Pendidikan,
			NomorTelpon: profiles[i].NomorTelpon,
		})
	}

	return &profileResponse
}
