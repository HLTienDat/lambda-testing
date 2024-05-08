package entity

import (
	"time"
)

type EmployeeData struct {
	Name      string    `json:"name"`
	Dob       time.Time `json:"dob"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CitizenID string    `json:"citizenId"`
	Address   string    `json:"address"`
}

type CreateEmployeeRequest struct {
	RequestID   string       `json:"requestId" validate:"required"`
	RequestTime time.Time    `json:"requestTime"`
	Data        EmployeeData `json:"data" validate:"required"`
}

type UpdateEmployeeRequest struct {
	RequestID   string       `json:"requestId" validate:"required"`
	RequestTime time.Time    `json:"requestTime"`
	EmployeeID  string       `json:"employeeId"`
	Data        EmployeeData `json:"data" validate:"required"`
}

type PostEmployeeResponse struct {
	RequestID   string    `json:"requestId" `
	CreatedTime time.Time `json:"createdTime"`
	Status      string    `json:"status"`
}

// type UpdateEmployeeResponse struct {
// 	RequestID   string    `json:"requestId" `
// 	CreatedTime time.Time `json:"createdTime"`
// 	Status      string    `json:"status"`
// }
