package model

import (
	"github.com/google/uuid"
)

type CreateProjectRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	DepartmentID string `json:"departmentId"`
}

type CreateProjectResponse struct {
	Id           uuid.UUID                `json:"id"`
	Name         string                   `json:"name"`
	Description  string                   `json:"description"`
}

type GetProjectResponse struct {
	Id           uuid.UUID                `json:"id"`
	Name         string                   `json:"name"`
	Description  string                   `json:"description"`
}
