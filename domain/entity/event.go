package entity

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Event struct {
	ID       string    `json:"id,omitempty" valid:"uuid"`
	Employee *Employee `json:"employee,omitempty" valid:"-"`
}

func NewEvent(employee *Employee) (*Event, error) {
	e := &Event{
		ID:       uuid.NewV4().String(),
		Employee: employee,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Event) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Event) ToJson() ([]byte, error) {
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
