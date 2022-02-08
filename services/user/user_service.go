package user

import (
	"web-app-example/interfaces"
	"web-app-example/models/user"
)

type UserService struct {
	repository interfaces.IUserRepository
}

func (s *UserService) CreateUser(data UserDTO) (user.User, error) {
	u := user.NewUser(data.FirstName, data.LastName, data.Address, data.Description)

	return s.repository.Store(*u)
}

func NewUserService(r interfaces.IUserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}
