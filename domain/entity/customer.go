package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/check-pad/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Customer struct {
	Base `json:",inline" valid:"-"`
}

func NewCustomer(id *string) (*Customer, error) {
	e := Customer{}
	e.ID = id
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Customer) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
