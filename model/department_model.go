package model

import "github.com/google/uuid"

type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateDepartmentResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Projects    []CreateProjectResponse
}

type GetDepartmentResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Projects    []CreateProjectResponse
}
