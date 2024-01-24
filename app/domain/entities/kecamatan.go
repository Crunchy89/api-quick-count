package entities

type Kecamatan struct {
	Base
	Kecamatan     string
	KabupatenKode int
	Kode          int       `gorm:"unique"`
	Kabupaten     Kabupaten `gorm:"references:Kode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
