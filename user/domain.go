package user

type UpdateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (ur UpdateRequest) Validate() (err error) {
	if ur.Password == "" {
		return errEmptyPassword
	}
	if ur.Name == "" {
		return errEmptyName
	}
	return
}

type CreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleType string `json:"role_type"`
}
