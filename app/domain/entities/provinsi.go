package entities

type Provinsi struct {
	Base
	Provinsi string
	Kode     int `gorm:"unique"`
}
