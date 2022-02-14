package entity

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/check-pad/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type WaitPaymentCheckPad struct {
	CheckPadID *string `json:"check_pad_id" valid:"uuid"`
}

func (e *WaitPaymentCheckPad) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

type CancelCheckPad struct {
	CheckPadID *string `json:"check_pad_id" valid:"uuid"`
}

func (e *CancelCheckPad) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

type CheckPad struct {
	Base           `json:",inline" valid:"-"`
	TotalPrice     *float64        `json:"total_price,omitempty" gorm:"column:total_price" valid:"-"`
	TotalDiscount  *float64        `json:"total_discount,omitempty" gorm:"column:total_discount" valid:"-"`
	FinalPrice     *float64        `json:"final_price,omitempty" gorm:"column:final_price" valid:"-"`
	Status         CheckPadStatus  `json:"status" gorm:"column:status;not null" valid:"checkPadStatus"`
	CanceledReason *string         `json:"canceled_reason,omitempty" gorm:"column:canceled_reason;type:varchar(255)" valid:"-"`
	Local          *string         `json:"local" gorm:"column:local;type:varchar(255)" valid:"required"`
	CustomerID     *string         `json:"customer_id" gorm:"column:customer_id;type:uuid;not null" valid:"uuid"`
	Customer       *Customer       `json:"-" valid:"-"`
	PlaceID        *string         `json:"place_id" gorm:"column:place_id;type:uuid;not null" valid:"uuid"`
	Place          *Place          `json:"-" valid:"-"`
	AttendantBy    *string         `json:"attendant_by" gorm:"column:attendant_by;type:uuid" valid:"uuid,optional"`
	Attendant      *Attendant      `json:"-" valid:"-"`
	Items          []*CheckPadItem `json:"-" gorm:"ForeignKey:CheckPadID" valid:"-"`
	items          []*CheckPadItem `json:"-" gorm:"-" valid:"-"`
}

func NewCheckPad(local *string, customer *Customer, place *Place) (*CheckPad, error) {
	e := CheckPad{
		Status:     CHECK_PAD_PENDING,
		Local:      local,
		CustomerID: customer.ID,
		Customer:   customer,
		PlaceID:    place.ID,
		Place:      place,
	}

	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *CheckPad) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *CheckPad) processPrice() error {
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

func (e *CheckPad) WaitPayment() (*WaitPaymentCheckPad, error) {
	if e.Status == CHECK_PAD_AWAITING_PAYMENT {
		return nil, errors.New("the check pad has already been awaiting payment")
	}

	e.Status = CHECK_PAD_AWAITING_PAYMENT
	e.UpdatedAt = utils.PTime(time.Now())

	if err := e.IsValid(); err != nil {
		return nil, err
	}

	return &WaitPaymentCheckPad{CheckPadID: e.ID}, nil
}

func (e *CheckPad) Cancel(canceledReason *string) (*CancelCheckPad, error) {
	if e.Status == CHECK_PAD_CANCELED {
		return nil, errors.New("the check pad has already been canceled")
	}

	if e.Status == CHECK_PAD_PAID {
		return nil, errors.New("the paid check pad cannot be canceled")
	}

	// TODO: adds the best way
	if len(e.items) > 0 || len(e.Items) > 0 {
		return nil, errors.New("the check pad cannot be canceled")
	}

	e.Status = CHECK_PAD_CANCELED
	e.CanceledReason = canceledReason
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return &CancelCheckPad{CheckPadID: e.ID}, err
}

func (e *CheckPad) Pay() error {
	if e.Status == CHECK_PAD_CANCELED {
		return errors.New("the canceled check pad cannot be paid")
	}

	e.Status = CHECK_PAD_PAID
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *CheckPad) AddItem(checkPadItem *CheckPadItem) error {
	if e.Status != CHECK_PAD_OPENED {
		return errors.New("the check pad cannot be changed")
	}

	// NOTE: change if multi check pad items are added
	e.items = append(e.Items, checkPadItem)
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.processPrice()
	return err
}

func (e *CheckPad) SetAttendant(attendant *Attendant) error {
	e.AttendantBy = attendant.ID
	e.Attendant = attendant
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}
