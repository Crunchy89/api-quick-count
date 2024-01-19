package repository

import "gorm.io/gorm"

type baseSuaraRepository struct {
	table *gorm.DB
}

type SuaraRepository interface {
}

func NewSuaraRepository(db *gorm.DB) SuaraRepository {
	table := db.Debug()
	return &baseSuaraRepository{table: table}
}
