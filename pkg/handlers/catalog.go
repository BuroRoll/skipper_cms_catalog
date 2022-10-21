package handlers

import (
	"Skipper_cms_catalog/pkg/models/forms/inputForms"
	"Skipper_cms_catalog/pkg/models/forms/outputForms"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BaseUrl /api/v1

// @Description Получение верхнего уровня католога
// @Security BearerAuth
// @Tags Catalog
// @Accept json
// @Produce json
// @Success 	200 		{object} 	[]models.MainCatalog
// @Router /catalog/main-catalog [get]
func (h *Handler) getMainCatalogElements(c *gin.Context) {
	mainCatalog := h.services.GetMainCatalog()
	c.JSON(http.StatusOK, mainCatalog)
}

// @Description Получение детей элемента каталога
// @Security BearerAuth
// @Tags Catalog
// @Accept json
// @Produce json
// @Param 			parent_id 		query 		int 	true "Id родителя элемента"
// @Param 			catalog_level 	query 		int 	true "Уровень элемента католога"
// @Success 	200 		{object} 	[]models.Catalog
// @Router /catalog/child [get]
func (h *Handler) getCatalogChild(c *gin.Context) {
	var params inputForms.CatalogChildInput
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, outputForms.ErrorResponse{
			Error: "Неверная форма запроса",
		})
		return
	}
	catalog := h.services.GetCatalogElements(params.ParentId, params.CatalogLevel)
	c.JSON(http.StatusOK, catalog)
}

// @Description Удаление элемента каталога
// @Security BearerAuth
// @Tags Catalog
// @Accept json
// @Produce json
// @Param 			catalog_id 		query 		int 	true "Id элемента каталога"
// @Param 			catalog_level 	query 		int 	true "Уровень элемента католога"
// @Success 	200 		{object} 	outputForms.DeleteResponse
// @Failure     500         {object}  	outputForms.ErrorResponse
// @Failure     400         {object}  	outputForms.ErrorResponse
// @Router /catalog [delete]
func (h *Handler) deleteCatalogElement(c *gin.Context) {
	var params inputForms.CatalogInput
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, outputForms.ErrorResponse{
			Error: "Неверная форма запроса",
		})
		return
	}
	err := h.services.DeleteCatalogElement(params.CatalogId, params.CatalogLevel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, outputForms.ErrorResponse{
			Error: "Ошибка удаления данных",
		})
		return
	}
	c.JSON(http.StatusOK, outputForms.DeleteResponse{
		Status: "Ok",
	})
}

// @Description Добавление элементов каталога
// @Security BearerAuth
// @Tags Catalog
// @Accept json
// @Produce json
// @Param 		request 	body 		inputForms.AddCatalogInput 	true 	"query params"
// @Success 	200 		{object} 	models.Catalog
// @Router /catalog [post]
func (h *Handler) addCatalogElement(c *gin.Context) {
	var params inputForms.AddCatalogInput
	if err := c.BindJSON(&params); err != nil {
		fmt.Println(err)
		fmt.Println(params)
		c.JSON(http.StatusBadRequest, outputForms.ErrorResponse{
			Error: "Неверная форма запроса",
		})
		return
	}
	fmt.Println("No error")
	fmt.Println(params)
	catalogElement, err := h.services.AddCatalogElement(params.CatalogName, params.CatalogLevel, params.CatalogParentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, outputForms.ErrorResponse{
			Error: "Ошибка добавления данных",
		})
		return
	}
	c.JSON(http.StatusOK, catalogElement)
}
