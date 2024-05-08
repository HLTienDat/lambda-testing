package entity

import (
	model "lambda-test/internal/repository/model"
	"time"
)

type GetEmployeesRequest struct {
	RequestID   string    `json:"requestId" validate:"required"`
	RequestTime time.Time `json:"requestTime"`
}

type GetEmployeeRequest struct {
	RequestID   string    `json:"requestId" validate:"required"`
	RequestTime time.Time `json:"requestTime"`
	EmployeeID  string    `json:"employeeId" validate:"required"`
}

type GetEmployeesResponse struct {
	RequestID   string           `json:"requestId" `
	CreatedTime time.Time        `json:"createdTime"`
	Employees   []model.Employee `json:"employees"`
}

type GetEmployeeResponse struct {
	RequestID   string         `json:"requestId" `
	CreatedTime time.Time      `json:"createdTime"`
	Employee    model.Employee `json:"employeeId" `
}
