package rest

import (
	"net/http"

	"github.com/c-4u/employee-service/domain/service"
	"github.com/gin-gonic/gin"
)

type EmployeeRestService struct {
	EmployeeService *service.EmployeeService
}

func NewEmployeeRestService(service *service.EmployeeService) *EmployeeRestService {
	return &EmployeeRestService{
		EmployeeService: service,
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
// @Param body body EmployeeRequest true "JSON body to create a new employee"
// @Success 200 {object} ID
// @Failure 401 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees [post]
func (s *EmployeeRestService) CreateEmployee(ctx *gin.Context) {
	var json EmployeeRequest

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

	id, err := s.EmployeeService.CreateEmployee(
		ctx,
		json.Username,
		json.FirstName,
		json.LastName,
		json.Email,
		json.Pis,
		json.Enabled,
		json.EmailVerified,
	)
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

	ctx.JSON(http.StatusOK, ID{ID: *id})
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
// @Success 200 {object} EmployeeResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees/{id} [get]
func (s *EmployeeRestService) FindEmployee(ctx *gin.Context) {
	var req ID

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

	employee, err := s.EmployeeService.FindEmployee(ctx, req.ID)
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

// SetPassword godoc
// @Security ApiKeyAuth
// @Summary set a employee password
// @Description Router for set a employee password
// @ID setPassword
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param body body PasswordInfo true "JSON body to set a employee password"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees/{id}/password [put]
func (s *EmployeeRestService) SetPassword(ctx *gin.Context) {
	var req ID
	var json PasswordInfo

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

	err := s.EmployeeService.SetPassword(ctx, req.ID, json.Password, json.Temporary)
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

	ctx.JSON(
		http.StatusOK,
		HTTPResponse{
			Code:    http.StatusOK,
			Message: "password updated successfully"},
	)
}

// SearchEmployees godoc
// @Security ApiKeyAuth
// @Summary search employees by filter
// @ID searchEmployees
// @Tags Employee
// @Description Search for employee employees by `filter`. if the page and page size are empty, 0 and 10 will be considered respectively.
// @Accept json
// @Produce json
// @Param body query SearchEmployeesRequest true "JSON body for search employees"
// @Success 200 {array} EmployeeResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees [get]
func (s *EmployeeRestService) SearchEmployees(ctx *gin.Context) {
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

	employees, err := s.EmployeeService.SearchEmployees(ctx, body.FirstName, body.LastName, body.PageSize, body.Page)
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

	ctx.JSON(http.StatusOK, employees)
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
func (s *EmployeeRestService) UpdateEmployee(ctx *gin.Context) {
	var req ID
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

	err := s.EmployeeService.UpdateEmployee(ctx, req.ID, json.FirstName, json.LastName, json.Email)
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
