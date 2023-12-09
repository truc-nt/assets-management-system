package services

import (
	"server/internal/db"
	"server/internal/models"
)

type IUserService interface {
	GetUsers(param *models.GetUsersParam) ([]*models.User, error)
	GetUserById(id uint32) (*models.User, error)
	FindUserByUsername(username string) error
	UpdateUser(id uint32, user *models.User) error
	CreateUser(UserRegisterParam models.UserRegisterParam) (uint32, error)
	AuthenticateUser(userAuthParam models.UserAuthParam) (models.User, error)
	SetLoginUser(id uint32) error
	SetLogoutUser(id uint32) error
}

type UserService struct {
	Repository models.IUserRepository
}

func NewUserService() IUserService {
	return &UserService{
		Repository: models.NewUserRepository(db.DB),
	}
}

func (s *UserService) FindUserByUsername(username string) error {
	err := s.Repository.FindUserByUsername(username)
	return err
}

func (s *UserService) CreateUser(UserRegisterParam models.UserRegisterParam) (uint32, error) {
	id, err := s.Repository.CreateUser(UserRegisterParam)
	return id, err
}

func (s *UserService) AuthenticateUser(userAuthParam models.UserAuthParam) (models.User, error) {
	user, err := s.Repository.AuthenticateUser(db.DB, userAuthParam)
	return user, err
}

func (s *UserService) SetLoginUser(id uint32) error {
	err := s.Repository.SetLoginUser(id)
	return err
}

func (s *UserService) SetLogoutUser(id uint32) error {
	err := s.Repository.SetLogoutUser(id)
	return err
}

func (s *UserService) GetUsers(param *models.GetUsersParam) ([]*models.User, error) {
	return s.Repository.GetUsers(param)
}

func (s *UserService) GetUserById(id uint32) (*models.User, error) {
	return s.Repository.GetUserById(id)
}

func (s *UserService) UpdateUser(id uint32, user *models.User) error {
	_, err := s.Repository.GetUserById(id)
	if err != nil {
		return err
	}
	user.ID = id
	return s.Repository.UpdateUser(user)
}
