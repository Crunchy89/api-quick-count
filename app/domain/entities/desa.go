package entities

type Desa struct {
	Base
	Desa          string
	KecamatanKode int
	Kecamatan     Kecamatan `gorm:"references:Kode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
