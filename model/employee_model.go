package model

import (
	"github.com/google/uuid"
)

type CreateEmployeeRequest struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
	Age      int8   `json:"age"`
	Status   string `json:"status"`
}

type CreateEmployeeResponse struct {
	Id          uuid.UUID                  `json:"id,omitempty"`
	FullName    string                     `json:"fullName"`
	Phone       string                     `json:"phone"`
	Age         int8                       `json:"age"`
	Status      string                     `json:"status"`
	Departments []CreateDepartmentResponse `json:"department,omitempty"`
}

type GetEmployeeResponse struct {
	Id          uuid.UUID                  `json:"id"`
	FullName    string                     `json:"fullName"`
	Phone       string                     `json:"phone"`
	Age         int8                       `json:"age"`
	Status      string                     `json:"status"`
	Departments []CreateDepartmentResponse `json:"department,omitempty"`
}
