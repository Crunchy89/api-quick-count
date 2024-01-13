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
	userSuperAdmin := &entities.User{Username: "ferdy", Password: "makannasi", Email: "rocker.hunt@gmail.com", RoleUUID: superAdmin.UUID}
	if err := db.Create(userSuperAdmin).Error; err != nil {
		log.Fatalf("error %s", err.Error())
	}

}
