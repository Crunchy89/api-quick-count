package entities

type Kabupaten struct {
	Base
	Kabupaten    string
	Kode         int `gorm:"unique"`
	ProvinsiKode int
	Provinsi     Provinsi `gorm:"references:Kode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
