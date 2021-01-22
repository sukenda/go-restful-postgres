package model

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
	Age      int8   `json:"age"`
}

type CreateUserResponse struct {
	Id       uuid.UUID              `json:"id"`
	Username string                 `json:"username"`
	Password string                 `json:"password,omitempty"`
	Email    string                 `json:"email"`
	Employee CreateEmployeeResponse `json:"employee,omitempty"`
}

type GetUserResponse struct {
	Id       uuid.UUID              `json:"id"`
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	Employee CreateEmployeeResponse `json:"employee,omitempty"`
}

type GetLoginResponse struct {
	AccessToken string      `json:"access_token"`
	User        interface{} `json:"user"`
}
