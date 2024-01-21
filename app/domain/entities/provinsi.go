package entities

type Provinsi struct {
	Base
	Provinsi string `gorm:"unique"`
}
