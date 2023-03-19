package user

type RegisterUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

func FormatRegisterUser(user User) RegisterUserResponse {
	format := RegisterUserResponse{}

	format.FirstName = user.FirstName
	format.LastName = user.LastName
	format.Email = user.Email
	format.Role = user.Role
	format.ID = user.ID
	format.Token = "123"

	return format
}
