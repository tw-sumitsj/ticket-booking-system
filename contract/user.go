package contract

type UserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Authenticated bool `json:"authenticated"`
}
