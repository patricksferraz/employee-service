package schema

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type UserEvent struct {
	ID   string `json:"id,omitempty" valid:"uuid"`
	User *User  `json:"user,omitempty" valid:"-"`
}

func NewUserEvent() *UserEvent {
	return &UserEvent{}
}

func (e *UserEvent) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *UserEvent) ParseJson(data []byte) error {
	err := json.Unmarshal(data, e)
	if err != nil {
		return err
	}

	err = e.isValid()
	if err != nil {
		return err
	}

	return nil
}
