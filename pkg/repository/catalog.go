package repository

import (
	"Skipper_cms_catalog/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CatalogPostgres struct {
	db *gorm.DB
}

func NewCatalogPostgres(db *gorm.DB) *CatalogPostgres {
	return &CatalogPostgres{db: db}
}

func (c CatalogPostgres) DeleteCatalogElement(catalogId uint, catalogLvl uint) error {
	switch catalogLvl {
	case 0:
		return c.db.Unscoped().Where("id = ?", catalogId).Select(clause.Associations).Delete(&models.Catalog0{}).Error
	case 1:
		return c.db.Unscoped().Where("id = ?", catalogId).Select(clause.Associations).Delete(&models.Catalog1{}).Error
	case 2:
		return c.db.Unscoped().Where("id = ?", catalogId).Select(clause.Associations).Delete(&models.Catalog2{}).Error
	case 3:
		return c.db.Unscoped().Where("id = ?", catalogId).Select(clause.Associations).Delete(&models.Catalog3{}).Error
	}
	return nil
}

func (c CatalogPostgres) GetMainCatalog() []models.Catalog0 {
	var catalog []models.Catalog0
	c.db.Select("name", "id").Find(&catalog)
	return catalog
}

func (c CatalogPostgres) GetCatalogFirstLevelElements(parentId uint) []models.Catalog1 {
	var catalog []models.Catalog1
	c.db.Where("parent_id = ?", parentId).Find(&catalog)
	return catalog
}

func (c CatalogPostgres) GetCatalogSecondLevelElements(parentId uint) []models.Catalog2 {
	var catalog []models.Catalog2
	c.db.Where("parent_id = ?", parentId).Find(&catalog)
	return catalog
}

func (c CatalogPostgres) GetCatalogThirdLevelElements(parentId uint) []models.Catalog3 {
	var catalog []models.Catalog3
	c.db.Where("parent_id = ?", parentId).Find(&catalog)
	return catalog
}

func (c CatalogPostgres) AddMainCatalogElement(name string) (*models.Catalog0, error) {
	catalog := models.Catalog0{
		Name: name,
	}
	err := c.db.Save(&catalog)
	return &catalog, err.Error
}

func (c CatalogPostgres) AddCatalogFirstElement(name string, parentId uint) (*models.Catalog1, error) {
	catalog := models.Catalog1{
		Name:     name,
		ParentId: parentId,
	}
	err := c.db.Save(&catalog)
	return &catalog, err.Error
}

func (c CatalogPostgres) AddCatalogSecondElement(name string, parentId uint) (*models.Catalog2, error) {
	catalog := models.Catalog2{
		Name:     name,
		ParentId: parentId,
	}
	err := c.db.Save(&catalog)
	return &catalog, err.Error
}

func (c CatalogPostgres) AddCatalogThirdElement(name string, parentId uint) (*models.Catalog3, error) {
	catalog := models.Catalog3{
		Name3:    name,
		ParentId: parentId,
	}
	err := c.db.Save(&catalog)
	return &catalog, err.Error
}

func (c CatalogPostgres) ChangeCatalogZeroLvlElementName(newName string, catalogId uint) (models.Catalog0, error) {
	var catalog models.Catalog0
	err := c.db.Where("id = ?", catalogId).Find(&catalog)
	err = c.db.Model(&catalog).Update("name", newName)
	if err.Error != nil {
		return models.Catalog0{}, err.Error
	}
	return catalog, nil
}

func (c CatalogPostgres) ChangeCatalogFirstLvlElementName(newName string, catalogId uint) (models.Catalog1, error) {
	var catalog models.Catalog1
	err := c.db.Where("id = ?", catalogId).Find(&catalog)
	err = c.db.Model(&catalog).Update("name", newName)
	if err.Error != nil {
		return models.Catalog1{}, err.Error
	}
	return catalog, nil
}

func (c CatalogPostgres) ChangeCatalogSecondLvlElementName(newName string, catalogId uint) (models.Catalog2, error) {
	var catalog models.Catalog2
	err := c.db.Where("id = ?", catalogId).Find(&catalog)
	err = c.db.Model(&catalog).Update("name", newName)
	if err.Error != nil {
		return models.Catalog2{}, err.Error
	}
	return catalog, nil
}

func (c CatalogPostgres) ChangeCatalogThirdLvlElementName(newName string, catalogId uint) (models.Catalog3, error) {
	var catalog models.Catalog3
	err := c.db.Where("id = ?", catalogId).Find(&catalog)
	err = c.db.Model(&catalog).Update("name3", newName)
	if err.Error != nil {
		return models.Catalog3{}, err.Error
	}
	return catalog, nil
}
