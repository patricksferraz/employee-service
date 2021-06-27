package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base          `json:",inline" valid:"required"`
	Username      string        `json:"username,omitempty" valid:"required"`
	FirstName     string        `json:"first_name,omitempty" valid:"alpha,required"`
	LastName      string        `json:"last_name,omitempty" valid:"alpha,required"`
	Email         string        `json:"email,omitempty" valid:"email"`
	Enabled       bool          `json:"enabled,omitempty" valid:"-"`
	EmailVerified bool          `json:"email_verified,omitempty" valid:"-"`
	Attributes    *EmployeeAttr `json:"attributes,omitempty" valid:"-"`
}

func NewEmployee(id, username, firstName, lastName, email, pis string, enabled, emailVerified bool, createdAt time.Time) (*Employee, error) {

	employee := &Employee{
		Username:      username,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		Enabled:       enabled,
		EmailVerified: emailVerified,
	}

	attr, err := NewEmployeeAttr(pis)
	if err != nil {
		return nil, err
	}

	employee.Attributes = attr

	if id == "" {
		employee.ID = uuid.NewV4().String()
	} else {
		employee.ID = id
	}

	if createdAt.IsZero() {
		employee.CreatedAt = time.Now()
	} else {
		employee.CreatedAt = createdAt
	}

	err = employee.isValid()
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *Employee) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Employee) Enable() error {
	e.Enabled = true
	err := e.isValid()
	return err
}

func (e *Employee) Disable() error {
	e.Enabled = false
	err := e.isValid()
	return err
}
