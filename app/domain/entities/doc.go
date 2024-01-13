package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `gorm:"primaryKey"`
	UUID      uuid.UUID      `gorm:"type:char(36);unique"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *Base) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.New()
	return nil
}
