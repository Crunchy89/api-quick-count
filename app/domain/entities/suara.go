package entities

import "github.com/google/uuid"

type Suara struct {
	Base
	Suara      int
	TPSUUID    uuid.UUID
	TPS        TPS `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PaslonUUID uuid.UUID
	Paslon     Paslon `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
