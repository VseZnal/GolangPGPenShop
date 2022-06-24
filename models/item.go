package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"Description"`
	Contact     string `json:"Contact"`
	//	Creator     User   `json:"foreignKey:Id"`
}

type CreateItem struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Contact     string `json:"contact" binding:"required"`
}

type UpdateItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
}
