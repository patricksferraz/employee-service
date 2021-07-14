package entity

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Filter struct {
	FirstName string `json:"first_name" valid:"optional"`
	LastName  string `json:"last_name" valid:"optional"`
	PageSize  int    `json:"page_size" valid:"optional"`
	Page      int    `json:"page" valid:"optional"`
}

func (e *Filter) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func NewFilter(firstName, lastName string, pageSize int, page int) (*Filter, error) {

	if pageSize == 0 {
		pageSize = 10
	}

	entity := &Filter{
		FirstName: firstName,
		LastName:  lastName,
		PageSize:  pageSize,
		Page:      page,
	}

	err := entity.isValid()
	if err != nil {
		return nil, err
	}

	return entity, nil
}
