package repository

import "gorm.io/gorm"

type baseTotalSuaraRepository struct {
	table *gorm.DB
}

type TotalSuaraRepository interface {
}

func NewTotalSuaraRepository(db *gorm.DB) TotalSuaraRepository {
	table := db.Debug()
	return &baseTotalSuaraRepository{table: table}
}
