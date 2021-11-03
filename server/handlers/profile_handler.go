package handlers

import (
	"go-profile/models"
	"go-profile/repositories"
	"go-profile/requests"
	"go-profile/responses"
	s "go-profile/server"
	profileservice "go-profile/services/profile"
	"go-profile/services/token"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type ProfileHandlers struct {
	server *s.Server
}

func NewProfileHandlers(server *s.Server) *ProfileHandlers {
	return &ProfileHandlers{server: server}
}

// CreateProfile godoc
// @Summary Create profile
// @Description Create profile
// @ID profiles-create
// @Tags Profiles Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateProfileRequest true "Profile alamat, pekerjaan, nama lengkap, pendidikan, nomor telpon"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Security ApiKeyAuth
// @Router /api/v1/profiles [post]
func (p *ProfileHandlers) CreateProfile(c echo.Context) error {
	createProfileRequest := new(requests.CreateProfileRequest)

	if err := c.Bind(createProfileRequest); err != nil {
		return err
	}

	if err := createProfileRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.JwtCustomClaims)
	id := claims.ID

	aProfile := models.Profile{}

	profileRepository := repositories.NewProfileRepository(p.server.DB)
	profileRepository.GetProfileByUserId(&aProfile, id)

	if aProfile.UserID != 0 {
		return responses.ErrorResponse(c, http.StatusConflict, "Profile found, cant create more profile, use an update")
	}

	profile := models.Profile{
		Alamat:      createProfileRequest.Alamat,
		Pekerjaan:   createProfileRequest.Pekerjaan,
		NamaLengkap: createProfileRequest.NamaLengkap,
		Pendidikan:  createProfileRequest.Pendidikan,
		NomorTelpon: createProfileRequest.NomorTelpon,
		UserID:      id,
	}
	profileService := profileservice.NewProfileService(p.server.DB)
	profileService.Create(&profile)

	return responses.MessageResponse(c, http.StatusCreated, "Profile successfully created")
}

// DeleteProfile godoc
// @Summary Delete profile
// @Description Delete profile
// @ID profiles-delete
// @Tags Profiles Actions
// @Param id path int true "Profile ID"
// @Success 204 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /api/v1/profiles/{id} [delete]
func (p *ProfileHandlers) DeleteProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	profile := models.Profile{}

	profileRepository := repositories.NewProfileRepository(p.server.DB)
	profileRepository.GetProfile(&profile, id)

	if profile.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Profile not found")
	}

	profileService := profileservice.NewProfileService(p.server.DB)
	profileService.Delete(&profile)

	return responses.MessageResponse(c, http.StatusNoContent, "Profile deleted successfully")
}

// GetProfiles godoc
// @Summary Get profiles
// @Description Get the list of all profiles
// @ID profiles-get
// @Tags Profiles Actions
// @Produce json
// @Success 200 {array} responses.ProfileResponse
// @Security ApiKeyAuth
// @Router /api/v1/profiles [get]
func (p *ProfileHandlers) GetProfile(c echo.Context) error {
	var profiles []models.Profile
	var users []models.User

	profileRepository := repositories.NewProfileRepository(p.server.DB)
	profileRepository.GetProfiles(&profiles)

	for i := 0; i < len(profiles); i++ {
		p.server.DB.Model(&profiles[i]).Related(&profiles[i].UserID)
	}

	userRepository := repositories.NewUserRepository(p.server.DB)
	userRepository.GetUsers(&users)

	for i := 0; i < len(users); i++ {
		p.server.DB.Model(&users[i])
	}

	response := responses.NewProfileResponse(profiles, users)
	return responses.Response(c, http.StatusOK, response)
}

// GetProfile godoc
// @Summary Get profile
// @Description Get the profiles by id
// @ID profiles-get-one
// @Tags Profiles Actions
// @Produce json
// @Param id path int true "Profile ID"
// @Success 200 {array} responses.ProfileResponse
// @Security ApiKeyAuth
// @Router /api/v1/profiles/{id} [get]
func (p *ProfileHandlers) GetAProfile(c echo.Context) error {
	var profile models.Profile
	id, _ := strconv.Atoi(c.Param("id"))

	profileRepository := repositories.NewProfileRepository(p.server.DB)
	profileRepository.GetProfile(&profile, id)

	return responses.Response(c, http.StatusOK, profile)
}

// UpdateProfile godoc
// @Summary Update profile
// @Description Update profile
// @ID profiles-update
// @Tags Profiles Actions
// @Accept json
// @Produce json
// @Param id path int true "Profile ID"
// @Param params body requests.UpdateProfileRequest true "Profile alamat, pekerjaan, nama lengkap, pendidikan, nomor telpon"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /api/v1/profiles/{id} [put]
func (p *ProfileHandlers) UpdateProfile(c echo.Context) error {
	updateProfileRequest := new(requests.UpdateProfileRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(updateProfileRequest); err != nil {
		return err
	}

	if err := updateProfileRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	profile := models.Profile{}

	profileRepository := repositories.NewProfileRepository(p.server.DB)
	profileRepository.GetProfile(&profile, id)

	if profile.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Profile not found")
	}

	profileService := profileservice.NewProfileService(p.server.DB)
	profileService.Update(&profile, updateProfileRequest)

	return responses.MessageResponse(c, http.StatusOK, "Profile successfully updated")
}
