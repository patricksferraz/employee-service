package rest

import (
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
// @Summary Create a employee
// @Description create employee
// @ID createEmployee
// @Tags Employee
// @Accept json
// @Produce json
// @Param body body Employee true "JSON body to create a new employee"
// @Success 200 {object} Employee
// @Failure 401 {object} HTTPError
// @Router /employee [post]
func (s *EmployeeRestService) CreateEmployee(ctx *gin.Context) {
	var json Employee

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := s.EmployeeService.CreateEmployee(
		ctx,
		json.Username,
		json.FirstName,
		json.LastName,
		json.Email,
		json.Attributes.Pis,
		json.Enabled,
		json.EmailVerified,
	)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}
