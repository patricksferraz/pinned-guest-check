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

type Place struct {
	Base `json:",inline" valid:"-"`
}

func NewPlace() (*Place, error) {
	e := Place{}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.isValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Place) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
