package repository_test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectingToDatabaseTest() (*gorm.DB, error) {
	dsn := "root:root@TechCodeForUs123@tcp(techcode.cloud:1234)/test?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
