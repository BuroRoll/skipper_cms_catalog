package inputForms

type CatalogInput struct {
	CatalogId    uint `form:"catalog_id" binding:"required"`
	CatalogLevel uint `form:"catalog_level" binding:"required"`
}

type CatalogChildInput struct {
	CatalogLevel uint `form:"catalog_level" binding:"required"`
	ParentId     uint `form:"parent_id" binding:"required"`
}

type AddCatalogInput struct {
	CatalogName     string `json:"catalog_name" binding:"required"`
	CatalogLevel    uint   `json:"catalog_level"`
	CatalogParentId uint   `json:"catalog_parent_id"`
}

type ChangeCatalogNameInput struct {
	NewCatalogName string `json:"newCatalogName"`
	CatalogLevel   uint   `json:"catalog_level"`
	CatalogId      uint   `json:"catalog_id"`
}
