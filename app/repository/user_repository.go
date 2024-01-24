package repository

import (
	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"github.com/crunchy89/api-quick-count/app/domain/responses"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveUser(user *entities.User) (*uint, error)
	CreateUser(user *entities.User) (*uuid.UUID, error)
	UpdateUserByID(user *entities.User) error
	UpdateUserByUUID(user *entities.User) error
	GetUserByUsername(username string) (*entities.User, error)
	GetUserByUUIDJWT(uuid uuid.UUID) (*responses.UserResponseJWT, error)
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

func (u baseUserRepository) GetUserByUsername(username string) (*entities.User, error) {
	user := new(entities.User)
	if err := u.table.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u baseUserRepository) GetUserByUUIDJWT(uuid uuid.UUID) (*responses.UserResponseJWT, error) {
	user := new(entities.User)
	userResponse := new(responses.UserResponseJWT)
	if err := u.table.Where("uuid = ?", uuid).First(&entities.User{}).Scan(user).Error; err != nil {
		return nil, err
	}
	userResponse.Active = user.Active
	userResponse.IsOnline = user.IsOnline
	userResponse.Username = user.Username
	userResponse.Role = "Petugas"
	if user.ID == 1 {
		userResponse.Role = "Admin"
	}
	return userResponse, nil
}
