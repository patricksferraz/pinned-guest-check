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

// CreateCheckPad godoc
// @Summary create a new check pad
// @ID createCheckPad
// @Tags Check Pad
// @Description Router for create a new check pad
// @Accept json
// @Produce json
// @Param body body CreateCheckPadRequest true "JSON body for create a new check pad"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads [post]
func (t *RestService) CreateCheckPad(c *fiber.Ctx) error {
	var req CreateCheckPadRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	checkPadID, err := t.Service.CreateCheckPad(c.Context(), &req.Local, &req.CustomerID, &req.PlaceID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *checkPadID})
}

// FindCheckPad godoc
// @Summary find a check pad
// @ID findCheckPad
// @Tags Check Pad
// @Description Router for find a check pad
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Success 200 {object} CheckPad
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id} [get]
func (t *RestService) FindCheckPad(c *fiber.Ctx) error {
	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	checkPad, err := t.Service.FindCheckPad(c.Context(), &checkPadID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(checkPad)
}

// WaitPaymentCheckPad godoc
// @Summary wait payment a check pad
// @ID waitPaymentCheckPad
// @Tags Check Pad
// @Description Router for wait payment a check pad
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/wait-payment [post]
func (t *RestService) WaitPaymentCheckPad(c *fiber.Ctx) error {
	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	err := t.Service.WaitPaymentCheckPad(c.Context(), &checkPadID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// CancelCheckPad godoc
// @Summary cancel a check pad
// @ID cancelCheckPad
// @Tags Check Pad
// @Description Router for cancel a check pad
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Param body body CancelCheckPadRequest true "JSON body for cancel a check pad"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/cancel [post]
func (t *RestService) CancelCheckPad(c *fiber.Ctx) error {
	var req CancelCheckPadRequest

	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	err := t.Service.CancelCheckPad(c.Context(), &checkPadID, &req.CanceledReason)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// AddCheckPadItem godoc
// @Summary add a check pad item
// @ID addCheckPadItem
// @Tags Check Pad
// @Description Router for add a check pad item
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Param body body AddCheckPadItemRequest true "JSON body for add a new check pad item"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/items [post]
func (t *RestService) AddCheckPadItem(c *fiber.Ctx) error {
	var req AddCheckPadItemRequest

	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	checkPadItemID, err := t.Service.AddCheckPadItem(c.Context(), &req.Name, &req.Code, &req.Quantity, &req.UnitPrice, &req.Discount, &req.Note, &req.Tag, &checkPadID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *checkPadItemID})
}

// FindCheckPadItem godoc
// @Summary find a check pad item
// @ID findCheckPadItem
// @Tags Check Pad
// @Description Router for find a check pad item
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Param check_pad_item_id path string true "Check pad item ID"
// @Success 200 {object} CheckPadItem
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/items/{check_pad_item_id} [get]
func (t *RestService) FindCheckPadItem(c *fiber.Ctx) error {
	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	checkPadItemID := c.Params("check_pad_item_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_item_id is not a valid uuid",
		})
	}

	checkPadItem, err := t.Service.FindCheckPadItem(c.Context(), &checkPadID, &checkPadItemID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(checkPadItem)
}

// CancelCheckPadItem godoc
// @Summary cancel a check pad item
// @ID cancelCheckPadItem
// @Tags Check Pad
// @Description Router for cancel a check pad item
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Param check_pad_item_id path string true "Check pad item ID"
// @Param body body CancelCheckPadItemRequest true "JSON body for cancel a check pad item"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/items/{check_pad_item_id}/cancel [post]
func (t *RestService) CancelCheckPadItem(c *fiber.Ctx) error { // TODO: add in rest->kafka<-kafka resources
	var req CancelCheckPadItemRequest

	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	checkPadItemID := c.Params("check_pad_item_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_item_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	err := t.Service.CancelCheckPadItem(c.Context(), &checkPadID, &checkPadItemID, &req.CanceledReason)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}
