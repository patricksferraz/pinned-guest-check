package rest

import (
	"fmt"
	"log"

	_ "github.com/c-4u/check-pad/app/rest/docs"
	"github.com/c-4u/check-pad/domain/service"
	"github.com/c-4u/check-pad/infra/db"
	"github.com/c-4u/check-pad/infra/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Check Pad Swagger API
// @version 1.0
// @description Swagger API for Check Pad Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email contato@coding4u.com.br

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func StartRestServer(pg *db.PostgreSQL, port int) {
	r := fiber.New()
	r.Use(cors.New())

	repository := repo.NewRepository(pg)
	service := service.NewService(repository)
	restService := NewRestService(service)

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/swagger/*", fiberSwagger.WrapHandler)
	{
		v1.Post("/customers", restService.CreateCustomer)
		v1.Get("/customers/:customer_id", restService.FindCustomer)

		v1.Post("/places", restService.CreatePlace)
		v1.Get("/places/:place_id", restService.FindPlace)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
