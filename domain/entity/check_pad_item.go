package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/check-pad/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type CheckPadItem struct {
	Base           `json:",inline" valid:"-"`
	Status         CheckPadItemStatus `json:"status" gorm:"column:status;not null" valid:"chackPadItemStatus"`
	CanceledReason *string            `json:"canceled_reason,omitempty" gorm:"column:canceled_reason;type:varchar(255)" valid:"-"`
	Name           *string            `json:"name" gorm:"column:name;not null" valid:"required"`
	Quantity       *int               `json:"quantity" gorm:"column:quantity;not null" valid:"required"`
	UnitPrice      *float64           `json:"unit_price" gorm:"column:unit_price;not null" valid:"required"`
	Discount       *float64           `json:"discount" gorm:"column:discount;not null" valid:"-"`
	FinalPrice     *float64           `json:"final_price" gorm:"column:final_price;not null" valid:"-"`
	Note           *string            `json:"note" gorm:"column:note" valid:"-"`
	Tags           []*string          `json:"tags" gorm:"-" valid:"-"`
	CheckPadID     *string            `json:"check_pad_id" gorm:"column:check_pad_id;type:uuid;not null" valid:"uuid"`
	CheckPad       *CheckPad          `json:"-" valid:"-"`
}

func NewCheckPadItem(name *string, quantity *int, unitPrice *float64, discount *float64, note *string, tags []*string, chekPad *CheckPad) (*CheckPadItem, error) {
	e := CheckPadItem{
		Name:       name,
		Status:     CHECK_PAD_ITEM_PENDING,
		Quantity:   quantity,
		UnitPrice:  unitPrice,
		Discount:   discount,
		Note:       note,
		Tags:       tags,
		CheckPadID: chekPad.ID,
		CheckPad:   chekPad,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())
	e.processPrice()

	err := e.isValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *CheckPadItem) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *CheckPadItem) processPrice() error {
	e.FinalPrice = utils.PFloat64(*e.UnitPrice * float64(*e.Quantity))
	if e.Discount != nil {
		e.FinalPrice = utils.PFloat64(*e.FinalPrice - *e.Discount)
	}
	err := e.isValid()
	return err
}

func (e *CheckPadItem) Cancel(canceledReason *string) error {
	e.Status = CHECK_PAD_ITEM_CANCELED
	e.CanceledReason = canceledReason
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}

func (e *CheckPadItem) Prepare() error {
	e.Status = CHECK_PAD_ITEM_PREPARING
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}

func (e *CheckPadItem) Forward() error {
	e.Status = CHECK_PAD_ITEM_ON_THE_WAY
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}

func (e *CheckPadItem) Deliver() error {
	e.Status = CHECK_PAD_ITEM_DELIVERED
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.isValid()
	return err
}
