package user

import (
	"web-app-example/interfaces"
	model "web-app-example/models/user"
)

type UserService struct {
	repository interfaces.IUserRepository
}

func NewUserService(r interfaces.IUserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) CreateUser(data UserDTO) (model.User, error) {
	u := model.NewUser(data.FirstName, data.LastName, data.Address, data.Description)

	return s.repository.Store(*u)
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repository.GetAll()
}
