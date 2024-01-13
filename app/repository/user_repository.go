package repository

import (
	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
}

type baseUserRepository struct {
	table *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) UserRepository {
	table := db.Debug()
	return &baseUserRepository{
		table: table,
	}
}

func (u baseUserRepository) SaveUser(user *entities.User) (*uint, error) {
	if err := u.table.Create(user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}

func (u baseUserRepository) CreateUser(user *entities.User) (*uuid.UUID, error) {
	if err := u.table.Create(user).Error; err != nil {
		return nil, err
	}
	return &user.UUID, nil
}

func (u baseUserRepository) UpdateUserByID(user *entities.User) error {
	if err := u.table.First(&user, user.ID).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u baseUserRepository) UpdateUserByUUID(user *entities.User) error {
	if err := u.table.Where("uuid = ?", user.UUID).Save(user).Error; err != nil {
		return err
	}
	return nil
}
