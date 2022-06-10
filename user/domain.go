package user

type createRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleType string `json:"roleType"`
}
