package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/pinned-guest-check/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base        `json:",inline" valid:"-"`
	GuestChecks []*GuestCheck `json:"guest_checks" gorm:"ForeignKey:AttendedBy" valid:"-"`
}

func NewEmployee(id *string) (*Employee, error) {
	e := Employee{}
	e.ID = id
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Employee) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
