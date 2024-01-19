package entities

import "github.com/google/uuid"

type Paslon struct {
	Base
	Nama       string `gorm:"unique"`
	Nomor      int    `gorm:"unique"`
	Foto       string
	PartaiUUID uuid.UUID
	Partai     Partai `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
