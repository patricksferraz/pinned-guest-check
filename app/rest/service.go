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
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /customers [post]
func (t *RestService) CreateCustomer(c *fiber.Ctx) error {
	customerID, err := t.Service.CreateCustomer(c.Context())
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *customerID})
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
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /places [post]
func (t *RestService) CreatePlace(c *fiber.Ctx) error {
	placeID, err := t.Service.CreatePlace(c.Context())
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *placeID})
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

// ReopenCheckPad godoc
// @Summary reopen a check pad
// @ID reopenCheckPad
// @Tags Check Pad
// @Description Router for reopen a check pad
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/reopen [post]
func (t *RestService) ReopenCheckPad(c *fiber.Ctx) error {
	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	err := t.Service.ReopenCheckPad(c.Context(), &checkPadID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
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

// PayCheckPad godoc
// @Summary pay a check pad
// @ID payCheckPad
// @Tags Check Pad
// @Description Router for pay a check pad
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/pay [post]
func (t *RestService) PayCheckPad(c *fiber.Ctx) error {
	checkPadID := c.Params("check_pad_id")
	if !govalidator.IsUUIDv4(checkPadID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "check_pad_id is not a valid uuid",
		})
	}

	err := t.Service.PayCheckPad(c.Context(), &checkPadID)
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

	checkPadItemID, err := t.Service.AddCheckPadItem(c.Context(), &req.Name, &req.Quantity, &req.UnitPrice, &req.Discount, &req.Note, &req.Tag, &checkPadID)
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
func (t *RestService) CancelCheckPadItem(c *fiber.Ctx) error {
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

// PrepareCheckPadItem godoc
// @Summary prepare a check pad item
// @ID prepareCheckPadItem
// @Tags Check Pad
// @Description Router for prepare a check pad item
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Param check_pad_item_id path string true "Check pad item ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/items/{check_pad_item_id}/prepare [post]
func (t *RestService) PrepareCheckPadItem(c *fiber.Ctx) error {
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

	err := t.Service.PrepareCheckPadItem(c.Context(), &checkPadID, &checkPadItemID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// ForwardCheckPadItem godoc
// @Summary forward a check pad item
// @ID forwardCheckPadItem
// @Tags Check Pad
// @Description Router for forward a check pad item
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Param check_pad_item_id path string true "Check pad item ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/items/{check_pad_item_id}/forward [post]
func (t *RestService) ForwardCheckPadItem(c *fiber.Ctx) error {
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

	err := t.Service.ForwardCheckPadItem(c.Context(), &checkPadID, &checkPadItemID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// DeliverCheckPadItem godoc
// @Summary deliver a check pad item
// @ID deliverCheckPadItem
// @Tags Check Pad
// @Description Router for deliver a check pad item
// @Accept json
// @Produce json
// @Param check_pad_id path string true "Check pad ID"
// @Param check_pad_item_id path string true "Check pad item ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /check-pads/{check_pad_id}/items/{check_pad_item_id}/deliver [post]
func (t *RestService) DeliverCheckPadItem(c *fiber.Ctx) error {
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

	err := t.Service.DeliverCheckPadItem(c.Context(), &checkPadID, &checkPadItemID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// CreateAttendant godoc
// @Summary create a new attendant
// @ID createAttendant
// @Tags Attendant
// @Description Router for create a new attendant
// @Accept json
// @Produce json
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /attendants [post]
func (t *RestService) CreateAttendant(c *fiber.Ctx) error {
	customerID, err := t.Service.CreateAttendant(c.Context())
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *customerID})
}

// FindAttendant godoc
// @Summary find a attendant
// @ID findAttendant
// @Tags Attendant
// @Description Router for find a attendant
// @Accept json
// @Produce json
// @Param attendant_id path string true "Attendant ID"
// @Success 200 {object} Attendant
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /attendants/{attendant_id} [get]
func (t *RestService) FindAttendant(c *fiber.Ctx) error {
	attendantID := c.Params("attendant_id")
	if !govalidator.IsUUIDv4(attendantID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "attendant_id is not a valid uuid",
		})
	}

	attendant, err := t.Service.FindAttendant(c.Context(), &attendantID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(attendant)
}
