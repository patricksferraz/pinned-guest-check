package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/c-4u/check-pad/domain/service"
	"github.com/gofiber/fiber/v2"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateCustomer godoc
// @Summary create a new customer
// @ID createCustomer
// @Tags Customer
// @Description Router for create a new customer
// @Accept json
// @Produce json
// @Success 200 {object} PostCustomerResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /customers [post]
func (t *RestService) CreateCustomer(c *fiber.Ctx) error {
	customerID, err := t.Service.CreateCustomer(c.Context())
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(PostCustomerResponse{ID: *customerID})
}

// FindCustomer godoc
// @Summary find a customer
// @ID findCustomer
// @Tags Customer
// @Description Router for find a customer
// @Accept json
// @Produce json
// @Param customer_id path string true "Customer ID"
// @Success 200 {object} Customer
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /customers/{customer_id} [get]
func (t *RestService) FindCustomer(c *fiber.Ctx) error {
	customerID := c.Params("customer_id")
	if !govalidator.IsUUIDv4(customerID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "customer_id is not a valid uuid",
		})
	}

	customer, err := t.Service.FindCustomer(c.Context(), &customerID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

// CreatePlace godoc
// @Summary create a new place
// @ID createPlace
// @Tags Place
// @Description Router for create a new place
// @Accept json
// @Produce json
// @Success 200 {object} PostPlaceResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /places [post]
func (t *RestService) CreatePlace(c *fiber.Ctx) error {
	placeID, err := t.Service.CreatePlace(c.Context())
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(PostPlaceResponse{ID: *placeID})
}

// FindPlace godoc
// @Summary find a place
// @ID findPlace
// @Tags Place
// @Description Router for find a place
// @Accept json
// @Produce json
// @Param place_id path string true "Place ID"
// @Success 200 {object} Place
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /places/{place_id} [get]
func (t *RestService) FindPlace(c *fiber.Ctx) error {
	placeID := c.Params("place_id")
	if !govalidator.IsUUIDv4(placeID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "place_id is not a valid uuid",
		})
	}

	place, err := t.Service.FindPlace(c.Context(), &placeID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(place)
}
