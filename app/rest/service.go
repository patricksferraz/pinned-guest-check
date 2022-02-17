package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-guest-check/domain/service"
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

// CreateGuestCheck godoc
// @Summary create a new guest check
// @ID createGuestCheck
// @Tags Guest Check
// @Description Router for create a new guest check
// @Accept json
// @Produce json
// @Param body body CreateGuestCheckRequest true "JSON body for create a new guest check"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guest-checks [post]
func (t *RestService) CreateGuestCheck(c *fiber.Ctx) error {
	var req CreateGuestCheckRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	guestCheckID, err := t.Service.CreateGuestCheck(c.Context(), &req.Local, &req.GuestID, &req.PlaceID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *guestCheckID})
}

// FindGuestCheck godoc
// @Summary find a gust check
// @ID findGuestCheck
// @Tags Guest Check
// @Description Router for find a gust check
// @Accept json
// @Produce json
// @Param guest_check_id path string true "Guest pad ID"
// @Success 200 {object} GuestCheck
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guest-checks/{guest_check_id} [get]
func (t *RestService) FindGuestCheck(c *fiber.Ctx) error {
	guestCheckID := c.Params("guest_check_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_id is not a valid uuid",
		})
	}

	guestCheck, err := t.Service.FindGuestCheck(c.Context(), &guestCheckID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(guestCheck)
}

// WaitPaymentGuestCheck godoc
// @Summary wait payment a guest check
// @ID waitPaymentGuestCheck
// @Tags Guest Check
// @Description Router for wait payment a guest check
// @Accept json
// @Produce json
// @Param guest_check_id path string true "Guest check ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guest-checks/{guest_check_id}/wait-payment [post]
func (t *RestService) WaitPaymentGuestCheck(c *fiber.Ctx) error {
	guestCheckID := c.Params("guest_check_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_id is not a valid uuid",
		})
	}

	err := t.Service.WaitPaymentGuestCheck(c.Context(), &guestCheckID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// CancelGuestCheck godoc
// @Summary cancel a guest check
// @ID cancelGuestCheck
// @Tags Guest Check
// @Description Router for cancel a guest check
// @Accept json
// @Produce json
// @Param guest_check_id path string true "Guest check ID"
// @Param body body CancelGuestCheckRequest true "JSON body for cancel a guest check"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guest-checks/{guest_check_id}/cancel [post]
func (t *RestService) CancelGuestCheck(c *fiber.Ctx) error {
	var req CancelGuestCheckRequest

	guestCheckID := c.Params("guest_check_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	err := t.Service.CancelGuestCheck(c.Context(), &guestCheckID, &req.CanceledReason)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// AddGuestCheckItem godoc
// @Summary add a guest check item
// @ID addGuestCheckItem
// @Tags Guest Check
// @Description Router for add a guest check item
// @Accept json
// @Produce json
// @Param guest_check_id path string true "Guest check ID"
// @Param body body AddGuestCheckItemRequest true "JSON body for add a new guest check item"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guest-checks/{guest_check_id}/items [post]
func (t *RestService) AddGuestCheckItem(c *fiber.Ctx) error {
	var req AddGuestCheckItemRequest

	guestCheckID := c.Params("guest_check_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	guestCheckItemID, err := t.Service.AddGuestCheckItem(c.Context(), &req.Name, &req.Code, &req.Quantity, &req.UnitPrice, &req.Discount, &req.Note, &req.Tag, &guestCheckID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *guestCheckItemID})
}

// FindGuestCheckItem godoc
// @Summary find a guest check item
// @ID findGuestCheckItem
// @Tags Guest Check
// @Description Router for find a guest check item
// @Accept json
// @Produce json
// @Param guest_check_id path string true "Guest check ID"
// @Param guest_check_item_id path string true "Guest check item ID"
// @Success 200 {object} GuestCheckItem
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guest-checks/{guest_check_id}/items/{guest_check_item_id} [get]
func (t *RestService) FindGuestCheckItem(c *fiber.Ctx) error {
	guestCheckID := c.Params("guest_check_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_id is not a valid uuid",
		})
	}

	guestCheckItemID := c.Params("guest_check_item_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_item_id is not a valid uuid",
		})
	}

	guestCheckItem, err := t.Service.FindGuestCheckItem(c.Context(), &guestCheckID, &guestCheckItemID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(guestCheckItem)
}

// CancelGuestCheckItem godoc
// @Summary cancel a guest check item
// @ID cancelGuestCheckItem
// @Tags Guest Check
// @Description Router for cancel a guest check item
// @Accept json
// @Produce json
// @Param guest_check_id path string true "Guest check ID"
// @Param guest_check_item_id path string true "Guest check item ID"
// @Param body body CancelGuestCheckItemRequest true "JSON body for cancel a guest check item"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guest-checks/{guest_check_id}/items/{guest_check_item_id}/cancel [post]
func (t *RestService) CancelGuestCheckItem(c *fiber.Ctx) error { // TODO: add in rest->kafka<-kafka resources
	var req CancelGuestCheckItemRequest

	guestCheckID := c.Params("guest_check_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_id is not a valid uuid",
		})
	}

	guestCheckItemID := c.Params("guest_check_item_id")
	if !govalidator.IsUUIDv4(guestCheckID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_check_item_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	err := t.Service.CancelGuestCheckItem(c.Context(), &guestCheckID, &guestCheckItemID, &req.CanceledReason)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}
