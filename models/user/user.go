package user

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int
	Email       string
	Description string
	Address     string
	Image       string
}

func NewUser(name string, lastName string, address string, description string) *User {
	return &User{
		FirstName:   name,
		LastName:    lastName,
		Address:     address,
		Description: description,
	}
}
