package entities

import "github.com/google/uuid"

type UserTPS struct {
	Base
	TotalSuara int
	Sah        int
	Batal      int
	C1         string
	TPSUUID    uuid.UUID
	TPS        TPS `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserUUID   uuid.UUID
	User       User `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
