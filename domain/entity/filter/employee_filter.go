package filter

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type EmployeeFilter struct {
	FirstName string `json:"first_name" valid:"optional"`
	LastName  string `json:"last_name" valid:"optional"`
	PageSize  int    `json:"page_size" valid:"optional"`
	PageToken string `json:"page_token" valid:"optional"`
}

func (e *EmployeeFilter) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func NewEmployeeFilter(firstName, lastName string, pageSize int, pageToken string) (*EmployeeFilter, error) {

	if pageSize == 0 {
		pageSize = 10
	}

	entity := &EmployeeFilter{
		FirstName: firstName,
		LastName:  lastName,
		PageSize:  pageSize,
		PageToken: pageToken,
	}

	err := entity.isValid()
	if err != nil {
		return nil, err
	}

	return entity, nil
}
