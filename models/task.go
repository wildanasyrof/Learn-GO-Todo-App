package models

type Task struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status string `gorm:"default:'pending'" json:"status"`
}
