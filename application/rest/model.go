package rest

import "time"

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type Employee struct {
	Base          `json:",inline"`
	Username      string `json:"username" binding:"required"`
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Pis           string `json:"pis" binding:"required"`
	Enabled       bool   `json:"enabled" binding:"required"`
	EmailVerified bool   `json:"email_verified" binding:"required"`
}

type HTTPError struct {
	Code  int    `json:"code,omitempty" example:"400"`
	Error string `json:"error,omitempty" example:"status bad request"`
}

type ID struct {
	ID string `json:"id" uri:"id" binding:"required,uuid"`
}

type HTTPResponse struct {
	Code    int    `json:"code,omitempty" example:"200"`
	Message string `json:"message,omitempty" example:"a message"`
}

type PasswordInfo struct {
	Password  string `json:"password" example:"mypassword"`
	Temporary bool   `json:"temporary" example:"false"`
}

type SearchEmployeesRequest struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	PageSize  int    `json:"page_size" form:"page_size" default:"10"`
	Page      int    `json:"page" form:"page" default:"0"`
}
