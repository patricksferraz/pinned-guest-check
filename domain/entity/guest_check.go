package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-guest-check/utils"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type GuestCheck struct {
	Base           `json:",inline" valid:"-"`
	TotalPrice     *float64          `json:"total_price,omitempty" gorm:"column:total_price" valid:"-"`
	TotalDiscount  *float64          `json:"total_discount,omitempty" gorm:"column:total_discount" valid:"-"`
	FinalPrice     *float64          `json:"final_price,omitempty" gorm:"column:final_price" valid:"-"`
	Status         GuestCheckStatus  `json:"status" gorm:"column:status;not null" valid:"guestCheckStatus"`
	CanceledReason *string           `json:"canceled_reason,omitempty" gorm:"column:canceled_reason;type:varchar(255)" valid:"-"`
	Local          *string           `json:"local" gorm:"column:local;type:varchar(255)" valid:"required"`
	Token          *string           `json:"-" gorm:"column:token;type:varchar(25);not null" valid:"-"`
	GuestID        *string           `json:"guest_id" gorm:"column:guest_id;type:uuid;not null" valid:"uuid"`
	Guest          *Guest            `json:"-" valid:"-"`
	PlaceID        *string           `json:"place_id" gorm:"column:place_id;type:uuid;not null" valid:"uuid"`
	Place          *Place            `json:"-" valid:"-"`
	AttendedBy     *string           `json:"attended_by,omitempty" gorm:"column:attended_by;type:uuid" valid:"uuid,optional"`
	Attendant      *Employee         `json:"-" gorm:"foreignKey:AttendedBy" valid:"-"`
	Items          []*GuestCheckItem `json:"-" gorm:"ForeignKey:GuestCheckID" valid:"-"`
	items          []*GuestCheckItem `json:"-" gorm:"-" valid:"-"`
}

func NewGuestCheck(local *string, guest *Guest, place *Place) (*GuestCheck, error) {
	token := primitive.NewObjectID().Hex()
	e := GuestCheck{
		Status:  GUEST_CHECK_PENDING,
		Local:   local,
		Token:   &token,
		GuestID: guest.ID,
		Guest:   guest,
		PlaceID: place.ID,
		Place:   place,
	}

	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *GuestCheck) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *GuestCheck) processPrice() error {
	var totalPrice float64
	var totalDiscount float64

	for _, i := range e.items {
		if i.FinalPrice != nil {
			totalPrice += *i.TotalPrice
		}
		if i.Discount != nil {
			totalDiscount += *i.Discount
		}
	}

	e.TotalPrice = &totalPrice
	e.TotalDiscount = &totalDiscount
	e.FinalPrice = utils.PFloat64(*e.TotalPrice - *e.TotalDiscount)
	err := e.IsValid()
	return err
}

func (e *GuestCheck) WaitPayment() error {
	if e.Status == GUEST_CHECK_AWAITING_PAYMENT {
		return errors.New("the guest check has already been awaiting payment")
	}

	if e.Status != GUEST_CHECK_OPENED {
		return errors.New("the guest check cannot wait payment")
	}

	e.Status = GUEST_CHECK_AWAITING_PAYMENT
	e.UpdatedAt = utils.PTime(time.Now())

	if err := e.IsValid(); err != nil {
		return err
	}

	return nil
}

func (e *GuestCheck) Cancel(canceledReason *string) error {
	if e.Status == GUEST_CHECK_CANCELED {
		return errors.New("the guest check has already been canceled")
	}

	if e.Status == GUEST_CHECK_PAID {
		return errors.New("the paid guest check cannot be canceled")
	}

	// TODO: adds the best way
	if len(e.items) > 0 || len(e.Items) > 0 {
		return errors.New("the guest check cannot be canceled")
	}

	e.Status = GUEST_CHECK_CANCELED
	e.CanceledReason = canceledReason
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheck) Pay() error {
	if e.Status == GUEST_CHECK_CANCELED {
		return errors.New("the canceled guest check cannot be paid")
	}

	e.Status = GUEST_CHECK_PAID
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *GuestCheck) AddItem(guestCheckItem *GuestCheckItem) error {
	if e.Status != GUEST_CHECK_OPENED {
		return errors.New("the guest check cannot be changed")
	}

	// NOTE: change if multi guest check items are added
	e.items = append(e.Items, guestCheckItem)
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.processPrice()
	return err
}

func (e *GuestCheck) Open(employee *Employee) error {
	e.Status = GUEST_CHECK_OPENED
	e.AttendedBy = employee.ID
	e.Attendant = employee
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

type SearchGuestChecks struct {
	Pagination `json:",inline" valid:"-"`
}

func NewSearchGuestChecks(pagination *Pagination) (*SearchGuestChecks, error) {
	e := SearchGuestChecks{}
	e.PageToken = pagination.PageToken
	e.PageSize = pagination.PageSize

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}
