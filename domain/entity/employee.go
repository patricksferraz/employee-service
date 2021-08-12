package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/employee-service/utils"
	"github.com/paemuri/brdoc"
	pisvalidator "github.com/patricksferraz/pisvalidator"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.TagMap["pis"] = govalidator.Validator(func(str string) bool {
		return pisvalidator.ValidatePis(str)
	})

	govalidator.TagMap["cpf"] = govalidator.Validator(func(str string) bool {
		return brdoc.IsCPF(str)
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base      `json:",inline" valid:"required"`
	FirstName string  `json:"first_name,omitempty" gorm:"column:first_name;type:varchar(50);not null" valid:"required"`
	LastName  string  `json:"last_name,omitempty" gorm:"column:last_name;type:varchar(255);not null" valid:"required"`
	Email     string  `json:"email,omitempty" gorm:"column:email;type:varchar(255);not null;unique" valid:"email"`
	Pis       string  `json:"pis,omitempty" gorm:"column:pis;type:varchar(25);not null;unique" valid:"pis"`
	Cpf       string  `json:"cpf,omitempty" gorm:"column:cpf;type:varchar(25);not null;unique" valid:"cpf"`
	Enabled   bool    `json:"enabled,omitempty" valid:"-"`
	Token     *string `json:"-" gorm:"column:token;type:varchar(25);not null" bson:"token" valid:"-"`
}

func NewEmployee(firstName, lastName, email, pis, cpf string) (*Employee, error) {

	utils.CleanNonDigits(&pis)
	utils.CleanNonDigits(&cpf)
	token := primitive.NewObjectID().Hex()
	employee := &Employee{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Pis:       pis,
		Cpf:       cpf,
		Enabled:   true,
		Token:     &token,
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
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) Disable() error {
	e.Enabled = false
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) SetFirstName(firstName string) error {
	e.FirstName = firstName
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) SetLastName(lastName string) error {
	e.LastName = lastName
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) SetEmail(email string) error {
	e.Email = email
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}
