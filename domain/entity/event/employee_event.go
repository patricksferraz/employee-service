package event

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/employee-service/domain/entity"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type EmployeeEvent struct {
	ID       string           `json:"id,omitempty" valid:"uuid"`
	Employee *entity.Employee `json:"employee,omitempty" valid:"-"`
}

func NewEmployeeEvent(employee *entity.Employee) (*EmployeeEvent, error) {
	e := &EmployeeEvent{
		ID:       uuid.NewV4().String(),
		Employee: employee,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *EmployeeEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *EmployeeEvent) ToJson() ([]byte, error) {
	err := e.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(e)
	if err != nil {
		return nil, nil
	}

	return result, nil
}
