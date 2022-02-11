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

type CheckPad struct {
	Base           `json:",inline" valid:"-"`
	TotalPrice     *float64        `json:"total_price" gorm:"column:total_price;not null" valid:"-"`
	TotalDiscount  *float64        `json:"total_discount" gorm:"column:total_discount;not null" valid:"-"`
	FinalPrice     *float64        `json:"final_price" gorm:"column:final_price;not null" valid:"-"`
	Status         CheckPadStatus  `json:"status" gorm:"column:status;not null" valid:"checkPadStatus"`
	CanceledReason *string         `json:"canceled_reason,omitempty" gorm:"column:canceled_reason;type:varchar(255)" valid:"-"`
	Local          *string         `json:"local" gorm:"column:local;type:varchar(255)" valid:"-"`
	CustomerID     *string         `json:"customer_id" gorm:"column:customer_id;type:uuid;not null" valid:"uuid"`
	Customer       *Customer       `json:"-" valid:"-"`
	PlaceID        *string         `json:"place_id" gorm:"column:place_id;type:uuid;not null" valid:"uuid"`
	Place          *Place          `json:"-" valid:"-"`
	Items          []*CheckPadItem `json:"-" gorm:"ForeignKey:CheckPadID" valid:"-"`
}

func NewCheckPad(local *string, customer *Customer, place *Place) (*CheckPad, error) {
	e := CheckPad{
		Local:      local,
		CustomerID: customer.ID,
		Customer:   customer,
		PlaceID:    place.ID,
		Place:      place,
		Status:     CHECK_PAD_OPENED,
	}

	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.isValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *CheckPad) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *CheckPad) processPrice() error {
	var totalPrice float64
	var totalDiscount float64

	for _, i := range e.Items {
		if i.FinalPrice != nil {
			totalPrice += *i.FinalPrice
		}
		if i.Discount != nil {
			totalDiscount += *i.Discount
		}
	}

	e.FinalPrice = utils.PFloat64(totalPrice - totalDiscount)
	err := e.isValid()
	return err
}

func (e *CheckPad) Reopen() error {
	if e.Status == CHECK_PAD_PAID {
		return errors.New("the paid check pad cannot be reopened")
	}

	e.Status = CHECK_PAD_OPENED
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}

func (e *CheckPad) WaitPayment() error {
	if e.Status == CHECK_PAD_AWAITING_PAYMENT {
		return errors.New("the check pad has already been awaiting payment")
	}

	e.Status = CHECK_PAD_AWAITING_PAYMENT
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}

func (e *CheckPad) Cancel(canceledReason *string) error {
	if e.Status == CHECK_PAD_CANCELED {
		return errors.New("the check pad has already been canceled")
	}

	if e.Status == CHECK_PAD_PAID {
		return errors.New("the paid check pad cannot be canceled")
	}

	e.Status = CHECK_PAD_CANCELED
	e.CanceledReason = canceledReason
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}

func (e *CheckPad) Pay() error {
	e.Status = CHECK_PAD_PAID
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}

func (e *CheckPad) AddItem(checkPadItem *CheckPadItem) error {
	if e.Status == CHECK_PAD_CANCELED || e.Status == CHECK_PAD_PAID || e.Status == CHECK_PAD_AWAITING_PAYMENT {
		return errors.New("the check pad cannot be changed")
	}

	e.Items = append(e.Items, checkPadItem)
	err := e.processPrice()
	return err
}
