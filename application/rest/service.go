package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patricksferraz/employee-service/domain/service"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateEmployee godoc
// @Security ApiKeyAuth
// @Summary create a employee
// @Description create employee
// @ID createEmployee
// @Tags Employee
// @Accept json
// @Produce json
// @Param body body CreateEmployeeRequest true "JSON body to create a new employee"
// @Success 200 {object} CreateEmployeeResponse
// @Failure 401 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees [post]
func (s *RestService) CreateEmployee(ctx *gin.Context) {
	var json CreateEmployeeRequest

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	id, err := s.Service.CreateEmployee(ctx, json.FirstName, json.LastName, json.Email, json.Pis, json.Cpf)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, CreateEmployeeResponse{ID: *id})
}

// FindEmployee godoc
// @Security ApiKeyAuth
// @Summary find a employee
// @Description Router for find a employee
// @ID findEmployee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} Employee
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees/{id} [get]
func (s *RestService) FindEmployee(ctx *gin.Context) {
	var req IDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	employee, err := s.Service.FindEmployee(ctx, req.ID)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// SearchEmployees godoc
// @Security ApiKeyAuth
// @Summary search employees by filter
// @ID searchEmployees
// @Tags Employee
// @Description Search for employee employees by `filter`. if the page size is empty, 10 will be considered.
// @Accept json
// @Produce json
// @Param body query SearchEmployeesRequest true "JSON body for search employees"
// @Success 200 {array} SearchEmployeesResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees [get]
func (s *RestService) SearchEmployees(ctx *gin.Context) {
	var body SearchEmployeesRequest

	if err := ctx.ShouldBindQuery(&body); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	nextPageToken, employees, err := s.Service.SearchEmployees(ctx, body.FirstName, body.LastName, body.PageSize, body.PageToken)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"next_page_token": *nextPageToken,
			"employees":       employees,
		},
	)
}

// UpdateEmployee godoc
// @Security ApiKeyAuth
// @Summary update a employee
// @Description Router for update a employee
// @ID updateEmployee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param body body UpdateEmployeeRequest true "JSON body to update a new employee"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees/{id} [put]
func (s *RestService) UpdateEmployee(ctx *gin.Context) {
	var req IDRequest
	var json UpdateEmployeeRequest

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	err := s.Service.UpdateEmployee(ctx, req.ID, json.FirstName, json.LastName, json.Email)
	if err != nil {
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, HTTPResponse{Code: http.StatusOK, Message: "updated successfully"})
}
