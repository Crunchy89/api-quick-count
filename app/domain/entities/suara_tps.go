package entities

import "github.com/google/uuid"

type SuaraTPS struct {
	Base
	UserTPSUUID uuid.UUID
	PaslonUUID  uuid.UUID
	UserTPS     UserTPS `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Paslon      Paslon  `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Suara       int
}
