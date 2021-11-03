package routes

import (
	"fmt"
	"go-profile/server"
	"go-profile/server/handlers"
	"go-profile/services/token"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *server.Server) {
	profileHandler := handlers.NewProfileHandlers(server)
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)

	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/api/v1/login", authHandler.Login)
	server.Echo.POST("/api/v1/register", registerHandler.Register)
	server.Echo.POST("/api/v1/refresh", authHandler.RefreshToken)

	fmt.Println(server.Config.Auth.AccessSecret)

	server.Echo.Static("/uploads", "uploads")

	r := server.Echo.Group("/api/v1")
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/profiles", profileHandler.GetProfile)
	r.GET("/profiles/:id", profileHandler.GetAProfile)
	r.POST("/profiles", profileHandler.CreateProfile)
	r.DELETE("/profiles/:id", profileHandler.DeleteProfile)
	r.PUT("/profiles/:id", profileHandler.UpdateProfile)
}
