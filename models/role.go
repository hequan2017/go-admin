package models

type Role struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Menu []Menu `json:"menu" gorm:"many2many:role_menu;"`
}
