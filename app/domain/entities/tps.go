package entities

import "github.com/google/uuid"

type TPS struct {
	Base
	Nama     string
	Nomor    int `gorm:"unique"`
	DesaUUID uuid.UUID
	Desa     Desa `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
