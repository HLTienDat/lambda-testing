package model

import "time"

type Employee struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Dob       time.Time `json:"dob"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CitizenId string    `json:"citizenId"`
	Address   string    `json:"address"`
}
