package entities

import "github.com/google/uuid"

type Paslon struct {
	Base
	Nama       string
	Nomor      int
	Foto       string
	PartaiUUID uuid.UUID
	Partai     Partai `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
