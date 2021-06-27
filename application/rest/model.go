package rest

import "time"

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type EmployeeAttr struct {
	Pis string `json:"pis" binding:"required"`
}

type Employee struct {
	Base          `json:",inline"`
	Username      string        `json:"username" binding:"required"`
	FirstName     string        `json:"first_name" binding:"required"`
	LastName      string        `json:"last_name" binding:"required"`
	Email         string        `json:"email" binding:"required"`
	Enabled       bool          `json:"enabled" binding:"required"`
	EmailVerified bool          `json:"email_verified" binding:"required"`
	Attributes    *EmployeeAttr `json:"attributes"`
}

type HTTPError struct {
	Error string `json:"error" example:"status bad request"`
}
