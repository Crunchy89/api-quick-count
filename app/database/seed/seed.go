package seed

import (
	"log"

	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	superAdmin := &entities.Role{Name: "SuperAdmin"}
	if err := db.Create(superAdmin).Error; err != nil {
		log.Fatalf("error %s", err.Error())
	}
	userSuperAdmin := &entities.User{Username: "superadmin", Password: "silahkanakses", RoleUUID: superAdmin.UUID}
	if err := db.Create(userSuperAdmin).Error; err != nil {
		log.Fatalf("error %s", err.Error())
	}

}
