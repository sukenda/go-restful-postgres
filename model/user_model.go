package model

type CreateUserRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Profile  string `json:"profile"`
}

type CreateUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Profile  string `json:"profile"`
}

type GetUserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Profile  string `json:"profile"`
}

type GetLoginResponse struct {
	AccessToken string      `json:"access_token"`
	User        interface{} `json:"user"`
}
