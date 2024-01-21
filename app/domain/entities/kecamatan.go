package entities

import "github.com/google/uuid"

type Kecamatan struct {
	Base
	Kecamatan     string
	KabupatenUUID uuid.UUID
	Kabupaten     Kabupaten `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
