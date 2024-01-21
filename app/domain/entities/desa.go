package entities

import "github.com/google/uuid"

type Desa struct {
	Base
	Desa          string
	KecamatanUUID uuid.UUID
	Kecamatan     Kecamatan `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
