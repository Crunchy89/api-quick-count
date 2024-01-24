package migration

import (
	"log"

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
			if err := db.Migrator().DropTable(table); err != nil {
				log.Fatalf("error drop table %s", err)
				return
			}
		}
	}
	for _, table := range data {
		if err := db.Migrator().CreateTable(table); err != nil {
			log.Fatalf("error create table %s", err)
			return
		}
	}
}
