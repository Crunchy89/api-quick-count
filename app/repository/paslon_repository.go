package repository

import (
	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type basePaslonRepository struct {
	table *gorm.DB
}

type PaslonRepository interface {
}

func NewPaslonRepository(db *gorm.DB) PaslonRepository {
	table := db.Debug()
	return &basePaslonRepository{table: table}
}

func (r *basePaslonRepository) Save(paslon *entities.Paslon) (*uint, error) {
	if err := r.table.Create(paslon).Error; err != nil {
		return nil, err
	}
	return &paslon.ID, nil
}

func (r *basePaslonRepository) Create(paslon *entities.Paslon) (*uuid.UUID, error) {
	if err := r.table.Create(paslon).Error; err != nil {
		return nil, err
	}
	return &paslon.UUID, nil
}
