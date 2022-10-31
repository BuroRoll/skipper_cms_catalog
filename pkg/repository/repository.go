package repository

import (
	"Skipper_cms_catalog/pkg/models"
	"gorm.io/gorm"
)

type Catalog interface {
	DeleteCatalogElement(catalogId uint, catalogLvl uint) error

	GetMainCatalog() []models.Catalog0
	GetCatalogFirstLevelElements(parentId uint) []models.Catalog1
	GetCatalogSecondLevelElements(parentId uint) []models.Catalog2
	GetCatalogThirdLevelElements(parentId uint) []models.Catalog3

	AddMainCatalogElement(name string) (*models.Catalog0, error)
	AddCatalogFirstElement(name string, parentId uint) (*models.Catalog1, error)
	AddCatalogSecondElement(name string, parentId uint) (*models.Catalog2, error)
	AddCatalogThirdElement(name string, parentId uint) (*models.Catalog3, error)

	ChangeCatalogZeroLvlElementName(newName string, catalogId uint) (models.Catalog0, error)
	ChangeCatalogFirstLvlElementName(newName string, catalogId uint) (models.Catalog1, error)
	ChangeCatalogSecondLvlElementName(newName string, catalogId uint) (models.Catalog2, error)
	ChangeCatalogThirdLvlElementName(newName string, catalogId uint) (models.Catalog3, error)
}

type Repository struct {
	Catalog
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Catalog: NewCatalogPostgres(db),
	}
}
