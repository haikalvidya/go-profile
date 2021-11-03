package main

import (
	"fmt"
	application "go-profile"
	"go-profile/config"

	"go-profile/docs"
)

// @title Go Profile
// @version 1.0
// @description This is a simple version of user profile using Echo app.

// @contact.name Vidya Haikal
// @contact.email haikalvidya@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	application.Start(cfg)
}
