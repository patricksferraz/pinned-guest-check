package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-guest-check/utils"
	"github.com/lib/pq"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.SetNilPtrAllowedByRequired(true)
}

type Item struct {
	Base      `json:",inline" valid:"-"`
	Code      *int            `json:"code" gorm:"column:code;not null" valid:"-"`
	Name      *string         `json:"name" gorm:"column:name;not null" valid:"required"`
	Available *bool           `json:"available" gorm:"column:available;not null" valid:"-"`
	Price     *float64        `json:"price" gorm:"column:price;not null" valid:"required"`
	Discount  *float64        `json:"discount,omitempty" gorm:"column:discount" valid:"-"`
	Tags      *pq.StringArray `json:"tags" gorm:"column:tags;type:text[]" valid:"-"`
}

func NewItem(id, name *string, code *int, price, discount *float64, available *bool, tags *[]string) (*Item, error) {
	e := Item{
		Code:      code,
		Name:      name,
		Price:     price,
		Discount:  discount,
		Available: available,
		Tags:      (*pq.StringArray)(tags),
	}
	e.ID = id
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Item) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Item) SetTags(tags *[]string) *Item {
	e.Tags = (*pq.StringArray)(tags)
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetAvailable(available *bool) *Item {
	e.Available = available
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetName(name *string) *Item {
	e.Name = name
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetPrice(price *float64) *Item {
	e.Price = price
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}

func (e *Item) SetDiscount(discount *float64) *Item {
	e.Discount = discount
	e.UpdatedAt = utils.PTime(time.Now())
	return e
}
