package models

import "gorm.io/gorm"

type Catalog0 struct {
	gorm.Model
	Name   string     `json:"name"`
	Child0 []Catalog1 `gorm:"ForeignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Catalog1 struct {
	gorm.Model
	ParentId uint       `json:"parent_id"`
	Name     string     `json:"name"`
	Child1   []Catalog2 `gorm:"ForeignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Catalog2 struct {
	gorm.Model
	ParentId uint       `json:"parent_id"`
	Name     string     `json:"name"`
	Child2   []Catalog3 `gorm:"ForeignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Catalog3 struct {
	gorm.Model
	ParentId uint   `json:"parent_id"`
	Name3    string `json:"name"`
	Count    uint   `json:"count"`
}

type MainCatalog struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

type Catalog struct {
	Name     string `json:"name"`
	Id       uint   `json:"id"`
	ParentId uint   `json:"parent_id"`
}
