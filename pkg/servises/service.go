package service

import (
	"Skipper_cms_catalog/pkg/models"
	"Skipper_cms_catalog/pkg/repository"
)

type Catalog interface {
	DeleteCatalogElement(catalogId uint, catalogLvl uint) error
	GetMainCatalog() []models.MainCatalog
	GetCatalogElements(parentId uint, catalogLevel uint) []models.Catalog
	AddCatalogElement(name string, level uint, parentId uint) (models.Catalog, error)
}

type Service struct {
	Catalog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Catalog: NewCatalogService(repos.Catalog),
	}
}
