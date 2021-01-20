package model

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}

type CreateUserResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}

type GetUserResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}

type GetLoginResponse struct {
	AccessToken string      `json:"access_token"`
	User        interface{} `json:"user"`
}
