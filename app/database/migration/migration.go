package migration

import (
	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	data := []interface{}{
		&entities.Role{},
		&entities.User{},
		&entities.Provinsi{},
		&entities.Kabupaten{},
		&entities.Kecamatan{},
		&entities.Desa{},
		&entities.Partai{},
		&entities.Paslon{},
		&entities.TPS{},
		&entities.UserTPS{},
		&entities.SuaraTPS{},
	}

	for _, table := range data {
		if db.Migrator().HasTable(table) {
			db.Migrator().DropTable(table)
		}
	}
	for _, table := range data {
		db.Migrator().CreateTable(table)
	}
}
