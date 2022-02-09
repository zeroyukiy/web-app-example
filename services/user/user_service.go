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

func (s *UserService) GetAllUsers() ([]UserDTO, error) {
	rawUsers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	var users []UserDTO
	for i := 0; i < len(rawUsers); i++ {
		users = append(users, UserDTO{
			FirstName:   rawUsers[i].FirstName,
			LastName:    rawUsers[i].LastName,
			Description: rawUsers[i].Description,
			Address:     rawUsers[i].Address,
		})
	}

	return users, nil
}
