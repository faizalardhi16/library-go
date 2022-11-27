package users

type RegisterUserFormatter struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

func FormatRegisterUser(user User, token string) RegisterUserFormatter {
	format := RegisterUserFormatter{}

	format.ID = user.ID
	format.FirstName = user.FirstName
	format.LastName = user.LastName
	format.Email = user.Email
	format.Role = user.Role
	format.Token = token

	return format
}
