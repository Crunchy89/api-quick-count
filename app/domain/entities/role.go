package entities

type Role struct {
	Base
	Edit   bool `gorm:"default:false"`
	Create bool `gorm:"default:false"`
	Delete bool `gorm:"default:false"`
	Name   string
}
