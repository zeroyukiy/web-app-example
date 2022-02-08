package user

type UserDTO struct {
	// Id          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
