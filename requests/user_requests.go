package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	minPathLength = 8
)

type BasicAuth struct {
	Email    string `json:"email,omitempty" validate:"required" example:"john.doe@example.com"`
	Password string `json:"password,omitempty" validate:"required" example:"password123456"`
}

func (ba BasicAuth) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(&ba.Email, is.Email),
		validation.Field(&ba.Password, validation.Length(minPathLength, 0)),
	)
}

type LoginRequest struct {
	BasicAuth
}

type RegisterRequest struct {
	BasicAuth
	Name string `json:"name,omitempty" validate:"required" example:"JohnDoe"`
}

func (rr RegisterRequest) Validate() error {
	err := rr.BasicAuth.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
	)
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}
