package rest

import (
	"fmt"
	"log"

	_ "dev.azure.com/c4ut/TimeClock/_git/employee-service/application/rest/docs"
	_service "dev.azure.com/c4ut/TimeClock/_git/employee-service/domain/service"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/infrastructure/external"
	"dev.azure.com/c4ut/TimeClock/_git/employee-service/infrastructure/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
)

// @title Employee Swagger API
// @version 1.0
// @description Swagger API for Golang Project Employee.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email contato@coding4u.com.br

// @BasePath /api/v1
func StartRestServer(keycloak *external.Keycloak, port int) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(apmgin.Middleware(r))

	employeeRepository := repository.NewKeycloakEmployeeRepository(keycloak)
	employeeService := _service.NewEmployeeService(employeeRepository)
	employeeRestService := NewEmployeeRestService(employeeService)

	v1 := r.Group("api/v1/employees")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.POST("/", employeeRestService.CreateEmployee)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Run(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
