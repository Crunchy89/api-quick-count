package entities

import "github.com/google/uuid"

type TPS struct {
	Base
	Nama     int `gorm:"unique"`
	Sah      int
	Batal    int
	UserUUID uuid.UUID
	User     User `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
