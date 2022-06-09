package login

type userLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
