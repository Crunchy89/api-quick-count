package entities

import (
	"github.com/crunchy89/api-quick-count/utils/password"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username string
	Password string
	Active   bool
	IsOnline bool
	RoleUUID uuid.UUID
	Role     Role `gorm:"references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.New()
	result, err := password.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = result
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	result, err := password.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = result
	return nil
}
