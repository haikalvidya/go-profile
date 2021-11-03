package requests

import validation "github.com/go-ozzo/ozzo-validation"

type BasicProfile struct {
	Alamat      string `json:"alamat,omitempty" validate:"required" example:"Jl. Echo"`
	Pekerjaan   string `json:"pekerjaan,omitempty" validate:"required" example:"DevOps"`
	NamaLengkap string `json:"nama_lengkap,omitempty" validate:"required" example:"Echo Parecho"`
	Pendidikan  string `json:"pendidikan_terakhir,omitempty" validate:"required" example:"Mahasiswa"`
	NomorTelpon string `json:"nomer_telpon,omitempty" validate:"required" example:"+6281234123421"`
}

func (bp BasicProfile) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.Alamat, validation.Required),
		validation.Field(&bp.Pekerjaan, validation.Required),
		validation.Field(&bp.NamaLengkap, validation.Required),
		validation.Field(&bp.Pendidikan, validation.Required),
		validation.Field(&bp.NomorTelpon, validation.Required),
	)
}

type CreateProfileRequest struct {
	BasicProfile
}

type UpdateProfileRequest struct {
	BasicProfile
}
