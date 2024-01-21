package entities

import "github.com/google/uuid"

type Kabupaten struct {
	Base
	Kabupaten    string
	ProvinsiUUID uuid.UUID
	Provinsi     Provinsi `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
