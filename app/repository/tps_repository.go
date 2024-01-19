package repository

import "gorm.io/gorm"

type baseTPSRepository struct {
	table *gorm.DB
}

type TPSRepository interface {
}

func NewTPSRepository(db *gorm.DB) TPSRepository {
	table := db.Debug()
	return &baseTPSRepository{table: table}
}
