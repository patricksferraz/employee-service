package rest

import "time"

type Base struct {
	ID        string    `json:"id,omitempty" binding:"uuid"`
	CreatedAt time.Time `json:"created_at,omitempty" time_format:"RFC3339"`
	UpdatedAt time.Time `json:"updated_at,omitempty" time_format:"RFC3339"`
}

type Employee struct {
	Base      `json:",inline"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Pis       string `json:"pis,omitempty"`
	Cpf       string `json:"cpf,omitempty"`
	Enabled   bool   `json:"enabled"`
}

type CreateEmployeeRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Pis       string `json:"pis" binding:"required"`
	Cpf       string `json:"cpf" binding:"required"`
}

type CreateEmployeeResponse struct {
	ID string `json:"id"`
}

type IDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type HTTPResponse struct {
	Code    int    `json:"code,omitempty" example:"200"`
	Message string `json:"message,omitempty" example:"a message"`
}

type HTTPError struct {
	Code  int    `json:"code,omitempty" example:"400"`
	Error string `json:"error,omitempty" example:"status bad request"`
}

type SearchEmployeesRequest struct {
	Filter `json:",inline"`
}

type Filter struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	PageSize  int    `json:"page_size" form:"page_size" default:"10"`
	PageToken string `json:"page_token" form:"page_token"`
}

type SearchEmployeesResponse struct {
	NextPageToken string     `json:"next_page_token"`
	Employees     []Employee `json:"employees"`
}

type UpdateEmployeeRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
