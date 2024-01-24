package service

import (
	"errors"

	"github.com/centrifugal/centrifuge"
	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"github.com/crunchy89/api-quick-count/app/domain/payloads"
	"github.com/crunchy89/api-quick-count/app/domain/responses"
	"github.com/crunchy89/api-quick-count/app/repository"
	"github.com/crunchy89/api-quick-count/utils/jwt"
	"github.com/crunchy89/api-quick-count/utils/password"
	"github.com/google/uuid"
)

type UserService interface {
	AuthService(data payloads.Auth) (*responses.AuthResponse, error)
	GetJWTResponse(uuid uuid.UUID) (*responses.UserResponseJWT, error)
	GetUserByUsername(username string) (*entities.User, error)
}

type baseUserService struct {
	roleRepo repository.RoleRepository
	userRepo repository.UserRepository
	socket   *centrifuge.Node
	jwt      jwt.JWTHelper
}

func NewUserService(
	roleRepo repository.RoleRepository,
	userRepo repository.UserRepository,
	socket *centrifuge.Node,
	jwt jwt.JWTHelper,
) UserService {
	return &baseUserService{
		roleRepo: roleRepo,
		userRepo: userRepo,
		socket:   socket,
		jwt:      jwt,
	}
}

func (r *baseUserService) AuthService(data payloads.Auth) (*responses.AuthResponse, error) {
	user, err := r.userRepo.GetUserByUsername(data.Username)
	if err != nil {
		return nil, err
	}
	check := password.Verify(data.Password, user.Password)
	if !check {
		return nil, errors.New("password salah")
	}
	role, err := r.roleRepo.GetByUUID(user.RoleUUID)
	if err != nil {
		return nil, err
	}
	token := r.jwt.GenerateTokenPublic(user.UUID)
	userRes := new(responses.UserResponseJWT)
	userRes.Username = user.Username
	userRes.Active = user.Active
	userRes.IsOnline = user.IsOnline
	userRes.Role = "petugas"
	if role.ID == 1 {
		userRes.Role = "admin"
	}

	res := &responses.AuthResponse{
		Token: token,
		User:  userRes,
	}
	return res, nil
}
func (r *baseUserService) GetUserByUsername(username string) (*entities.User, error) {
	return r.userRepo.GetUserByUsername(username)
}

func (r *baseUserService) GetJWTResponse(uuid uuid.UUID) (*responses.UserResponseJWT, error) {
	return r.userRepo.GetUserByUUIDJWT(uuid)
}
