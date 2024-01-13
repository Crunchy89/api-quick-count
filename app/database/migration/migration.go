package migration

import (
	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	data := []interface{}{
		&entities.Role{},
		&entities.User{},
	}

	for _, table := range data {
		if db.Migrator().HasTable(table) {
			db.Migrator().DropTable(table)
			continue
		}
		db.Migrator().CreateTable(table)
	}
}
