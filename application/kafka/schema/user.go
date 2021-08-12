package schema

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	Base       `json:",inline" valid:"required"`
	Username   string `json:"username,omitempty" valid:"required"`
	EmployeeID string `json:"employee_id,omitempty" valid:"required"`
}

func NewUser() *User {
	return &User{}
}
