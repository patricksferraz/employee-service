package rest

import (
	"fmt"
	"net/http"

	"dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/service"
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
// @Summary create a employee
// @Description create employee
// @ID createEmployee
// @Tags Employee
// @Accept json
// @Produce json
// @Param body body Employee true "JSON body to create a new employee"
// @Success 200 {object} Employee
// @Failure 401 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /employees [post]
func (s *EmployeeRestService) CreateEmployee(ctx *gin.Context) {
	var json Employee

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

	employee, err := s.EmployeeService.CreateEmployee(
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

	ctx.JSON(http.StatusOK, employee)
}

// FindEmployee godoc
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
func (s *EmployeeRestService) FindEmployee(ctx *gin.Context) {
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
	fmt.Println(employee)

	ctx.JSON(http.StatusOK, employee)
}
