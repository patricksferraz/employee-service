package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	pisvalidatior "github.com/patricksferraz/pisvalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.TagMap["pis"] = govalidator.Validator(func(str string) bool {
		return pisvalidatior.ValidatePis(str)
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base          `json:",inline" valid:"required"`
	Username      string `json:"username,omitempty" valid:"required"`
	FirstName     string `json:"first_name,omitempty" valid:"alpha,required"`
	LastName      string `json:"last_name,omitempty" valid:"alpha,required"`
	Email         string `json:"email,omitempty" valid:"email"`
	Pis           string `json:"pis,omitempty" attr:"pis" valid:"pis"`
	Enabled       bool   `json:"enabled,omitempty" valid:"-"`
	EmailVerified bool   `json:"email_verified,omitempty" valid:"-"`
}

func NewEmployee(username, firstName, lastName, email, pis string, enabled, emailVerified bool) (*Employee, error) {

	employee := &Employee{
		Username:      username,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		Pis:           pis,
		Enabled:       enabled,
		EmailVerified: emailVerified,
	}
	employee.ID = uuid.NewV4().String()
	employee.CreatedAt = time.Now()

	if err := employee.isValid(); err != nil {
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
