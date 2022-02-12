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
		customer := v1.Group("/customers")
		customer.Post("", restService.CreateCustomer)
		customer.Get("/:customer_id", restService.FindCustomer)

		place := v1.Group("/places")
		place.Post("", restService.CreatePlace)
		place.Get("/:place_id", restService.FindPlace)

		checkPad := v1.Group("/check-pads")
		checkPad.Post("", restService.CreateCheckPad)
		checkPad.Get("/:check_pad_id", restService.FindCheckPad)
		checkPad.Post("/:check_pad_id/reopen", restService.ReopenCheckPad)
		checkPad.Post("/:check_pad_id/wait-payment", restService.WaitPaymentCheckPad)
		checkPad.Post("/:check_pad_id/cancel", restService.CancelCheckPad)
		checkPad.Post("/:check_pad_id/pay", restService.PayCheckPad)

		checkPadItem := checkPad.Group("/:check_pad_id/items")
		checkPadItem.Post("", restService.AddCheckPadItem)
		checkPadItem.Get("/:check_pad_item_id", restService.FindCheckPadItem)
		checkPadItem.Post("/:check_pad_item_id/cancel", restService.CancelCheckPadItem)
		checkPadItem.Post("/:check_pad_item_id/prepare", restService.PrepareCheckPadItem)
		checkPadItem.Post("/:check_pad_item_id/forward", restService.ForwardCheckPadItem)
		checkPadItem.Post("/:check_pad_item_id/deliver", restService.DeliverCheckPadItem)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
