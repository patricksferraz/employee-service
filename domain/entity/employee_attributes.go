package entity

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type EmployeeAttr struct {
	Pis string `json:"pis,omitempty" valid:"required"`
}

func NewEmployeeAttr(pis string) (*EmployeeAttr, error) {

	e := &EmployeeAttr{
		Pis: pis,
	}

	err := e.isValid()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *EmployeeAttr) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
