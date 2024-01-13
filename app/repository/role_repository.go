package repository

import (
	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"github.com/crunchy89/api-quick-count/app/domain/responses"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleRepository defines the interface for role operations.
type RoleRepository interface {
	Save(role *entities.Role) (*uint, error)
	Create(role *entities.Role) (*uuid.UUID, error)
	GetAll() ([]*responses.RoleResponse, error)
	GetByID(id uint) (*responses.RoleResponse, error)
	GetByUUID(uuid uuid.UUID) (*responses.RoleResponse, error)
	UpdateById(role *entities.Role) error
	UpdateByUuid(role *entities.Role) error
	DeleteById(id uint) error
	DeleteByUuid(uuid uuid.UUID) error
}

// baseRoleRepository is the implementation of RoleRepository.
type baseRoleRepository struct {
	table *gorm.DB
}

// NewRoleRepository creates a new instance of baseRoleRepository.
func NewRoleRepository(db *gorm.DB) RoleRepository {
	table := db.Debug()
	return &baseRoleRepository{table: table}
}

func (r *baseRoleRepository) Save(role *entities.Role) (*uint, error) {
	if err := r.table.Create(role).Error; err != nil {
		return nil, err
	}
	return &role.ID, nil
}

func (r *baseRoleRepository) Create(role *entities.Role) (*uuid.UUID, error) {
	if err := r.table.Create(role).Error; err != nil {
		return nil, err
	}
	return &role.UUID, nil
}

func (r *baseRoleRepository) GetAll() ([]*responses.RoleResponse, error) {
	var roles []*responses.RoleResponse
	if err := r.table.Find(&entities.Role{}).Scan(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *baseRoleRepository) GetByID(id uint) (*responses.RoleResponse, error) {
	role := new(responses.RoleResponse)
	if err := r.table.First(&entities.Role{}, id).Scan(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *baseRoleRepository) GetByUUID(uuid uuid.UUID) (*responses.RoleResponse, error) {
	role := new(responses.RoleResponse)
	if err := r.table.Where("uuid = ?", uuid).First(&entities.Role{}).Scan(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *baseRoleRepository) UpdateById(role *entities.Role) error {
	if err := r.table.First(&role, role.ID).Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *baseRoleRepository) UpdateByUuid(role *entities.Role) error {
	if err := r.table.Where("uuid = ?", role.UUID).Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *baseRoleRepository) DeleteById(id uint) error {
	if err := r.table.Delete(&entities.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *baseRoleRepository) DeleteByUuid(uuid uuid.UUID) error {
	if err := r.table.Where("uuid = ?", uuid).Delete(&entities.Role{}).Error; err != nil {
		return err
	}
	return nil
}
