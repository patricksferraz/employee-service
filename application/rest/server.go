package rest

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/patricksferraz/employee-service/application/rest/docs"
	_service "github.com/patricksferraz/employee-service/domain/service"
	"github.com/patricksferraz/employee-service/infrastructure/db"
	"github.com/patricksferraz/employee-service/infrastructure/external"
	"github.com/patricksferraz/employee-service/infrastructure/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
	"google.golang.org/grpc"
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
func StartRestServer(database *db.Postgres, authConn *grpc.ClientConn, kafkaProducer *external.KafkaProducer, port int) {
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

	authService := _service.NewAuthService(authConn)
	authMiddlerare := NewAuthMiddleware(authService)
	repository := repository.NewRepository(database, kafkaProducer)
	service := _service.NewService(repository)
	restService := NewRestService(service)

	v1 := r.Group("api/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		authorized := v1.Group("/employees", authMiddlerare.Require())
		{
			authorized.POST("", restService.CreateEmployee)
			authorized.GET("", restService.SearchEmployees)
			authorized.GET("/:id", restService.FindEmployee)
			authorized.PUT("/:id", restService.UpdateEmployee)
		}
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Run(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
