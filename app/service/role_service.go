package service

import (
	"github.com/centrifugal/centrifuge"
	"github.com/crunchy89/api-quick-count/app/domain/responses"
	"github.com/crunchy89/api-quick-count/app/repository"
)

type RoleService interface {
	GetAllRole() ([]*responses.RoleResponse, error)
}

type baseRoleService struct {
	roleRepo repository.RoleRepository
	socket   *centrifuge.Node
}

func NewRoleService(
	roleRepo repository.RoleRepository,
	socket *centrifuge.Node,
) RoleService {
	return &baseRoleService{
		roleRepo: roleRepo,
		socket:   socket,
	}
}

func (r *baseRoleService) GetAllRole() ([]*responses.RoleResponse, error) {
	return r.roleRepo.GetAll()
}
