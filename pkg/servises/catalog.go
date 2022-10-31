package service

import (
	"Skipper_cms_catalog/pkg/models"
	"Skipper_cms_catalog/pkg/repository"
	"encoding/json"
	"fmt"
)

type CatalogService struct {
	repo repository.Catalog
}

func NewCatalogService(repo repository.Catalog) *CatalogService {
	return &CatalogService{repo: repo}
}

func (c CatalogService) DeleteCatalogElement(catalogId uint, catalogLvl uint) error {
	return c.repo.DeleteCatalogElement(catalogId, catalogLvl)
}

func (c CatalogService) GetMainCatalog() []models.MainCatalog {
	mainCatalog := c.repo.GetMainCatalog()
	j, _ := json.Marshal(mainCatalog)
	var mainCatalogElements []models.MainCatalog
	json.Unmarshal(j, &mainCatalogElements)
	return mainCatalogElements
}

func (c CatalogService) GetCatalogElements(parentId uint, catalogLevel uint) []models.Catalog {
	switch catalogLevel {
	case 1:
		catalogElements := c.repo.GetCatalogFirstLevelElements(parentId)
		return TransformCatalogElements(catalogElements)
	case 2:
		catalogElements := c.repo.GetCatalogSecondLevelElements(parentId)
		return TransformCatalogElements(catalogElements)
	case 3:
		catalogElements := c.repo.GetCatalogThirdLevelElements(parentId)
		fmt.Println(catalogElements)
		return TransformCatalogElements(catalogElements)
	}
	return nil
}

func (c CatalogService) AddCatalogElement(name string, level uint, parentId uint) (models.Catalog, error) {
	var element any
	var err error
	switch level {
	case 0:
		element, err = c.repo.AddMainCatalogElement(name)
	case 1:
		element, err = c.repo.AddCatalogFirstElement(name, parentId)
	case 2:
		element, err = c.repo.AddCatalogSecondElement(name, parentId)
	case 3:
		element, err = c.repo.AddCatalogThirdElement(name, parentId)
	}
	catalogElement := TransformCatalogElement(element)
	if err != nil {
		return catalogElement, err
	}
	return catalogElement, nil
}

func (c CatalogService) ChangeCatalogElementName(newName string, level uint, catalogId uint) (models.Catalog, error) {
	var element any
	var err error
	switch level {
	case 0:
		element, err = c.repo.ChangeCatalogZeroLvlElementName(newName, catalogId)
		fmt.Println(element)
	case 1:
		element, err = c.repo.ChangeCatalogFirstLvlElementName(newName, catalogId)
	case 2:
		element, err = c.repo.ChangeCatalogSecondLvlElementName(newName, catalogId)
	case 3:
		element, err = c.repo.ChangeCatalogThirdLvlElementName(newName, catalogId)
	}

	catalogElement := TransformCatalogElement(element)
	fmt.Println(element)
	if err != nil {
		return catalogElement, err
	}
	return catalogElement, nil
}

func TransformCatalogElements(catalog any) []models.Catalog {
	j, _ := json.Marshal(catalog)
	var catalogElements []models.Catalog
	json.Unmarshal(j, &catalogElements)
	return catalogElements
}

func TransformCatalogElement(catalog any) models.Catalog {
	j, _ := json.Marshal(catalog)
	var catalogElement models.Catalog
	json.Unmarshal(j, &catalogElement)
	return catalogElement
}
