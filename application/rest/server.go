package rest

import (
	"fmt"
	"log"

	"github.com/c-4u/employee-service/application/grpc/pb"
	_ "github.com/c-4u/employee-service/application/rest/docs"
	_service "github.com/c-4u/employee-service/domain/service"
	"github.com/c-4u/employee-service/infrastructure/external"
	"github.com/c-4u/employee-service/infrastructure/repository"
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
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func StartRestServer(port int, keycloak *external.Keycloak, service pb.AuthServiceClient, kafka *external.Kafka) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))
	r.Use(apmgin.Middleware(r))

	authService := _service.NewAuthService(service)
	authMiddlerare := NewAuthMiddleware(authService)
	employeeRepository := repository.NewKeycloakEmployeeRepository(keycloak)
	kafkaRepository := repository.NewKafkaRepository(kafka)
	employeeService := _service.NewEmployeeService(employeeRepository, kafkaRepository)
	employeeRestService := NewEmployeeRestService(employeeService)

	v1 := r.Group("api/v1/employees")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		authorized := v1.Group("", authMiddlerare.Require())
		{
			authorized.POST("", employeeRestService.CreateEmployee)
			authorized.GET("", employeeRestService.SearchEmployees)
			authorized.GET("/:id", employeeRestService.FindEmployee)
			authorized.PUT("/:id", employeeRestService.UpdateEmployee)
			authorized.PUT("/:id/password", employeeRestService.SetPassword)
		}
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Run(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
