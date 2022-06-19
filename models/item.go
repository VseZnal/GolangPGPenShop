package models

type Item struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	Contact     string `json:"Contact"`
	Phone       string `json:"Phone"`
}
