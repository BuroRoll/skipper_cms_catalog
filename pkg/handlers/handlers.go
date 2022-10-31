package handlers

import (
	"Skipper_cms_catalog/pkg/docs"
	service "Skipper_cms_catalog/pkg/servises"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() {
	router := gin.Default()
	router.Use(corsMiddleware())

	api_v1 := router.Group("/api/v1", h.userIdentity)
	{
		catalog := api_v1.Group("/catalog")
		{
			catalog.GET("/main-catalog", h.getMainCatalogElements)
			catalog.GET("/child", h.getCatalogChild)

			catalog.DELETE("/", h.deleteCatalogElement)

			catalog.POST("/", h.addCatalogElement)

			catalog.PUT("/", h.changeCatalogElementName)
		}
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
