package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/pinned-guest-check/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Pagination struct {
	PageToken *string `json:"page_token" valid:"-"`
	PageSize  *int    `json:"page_size" valid:"-"`
}

func NewPagination(pageToken *string, pageSize *int) (*Pagination, error) {
	if pageSize == nil {
		pageSize = utils.PInt(10)
	}

	e := Pagination{
		PageToken: pageToken,
		PageSize:  pageSize,
	}

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Pagination) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
