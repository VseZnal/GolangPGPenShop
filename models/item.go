package models

type Item struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	Contact     string `json:"Contact"`
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
