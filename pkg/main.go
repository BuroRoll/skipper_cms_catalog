package main

import (
	"Skipper_cms_catalog/pkg/handlers"
	"Skipper_cms_catalog/pkg/models"
	"Skipper_cms_catalog/pkg/repository"
	service "Skipper_cms_catalog/pkg/servises"
)

// @title Catalog service
// @version 1.0
// @description Catalog moderation methods for skipper cms
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := models.GetDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlerses := handlers.NewHandler(services)
	handlerses.InitRoutes()
}
