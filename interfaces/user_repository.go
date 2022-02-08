package interfaces

import (
	model "web-app-example/models/user"
)

type IUserRepository interface {
	Store(u model.User) (model.User, error)
}
