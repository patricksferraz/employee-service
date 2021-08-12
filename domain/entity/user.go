package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	Base       `json:",inline" valid:"required"`
	Username   string    `json:"username" gorm:"column:username;type:varchar(50);not null" valid:"required"`
	EmployeeID string    `json:"employee_id" gorm:"column:employee_id;type:uuid;not null" valid:"uuid"`
	Employee   *Employee `json:"-" valid:"-"`
}

func NewUser(id, username string, employee *Employee) (*User, error) {
	e := &User{
		Username:   username,
		EmployeeID: employee.ID,
		Employee:   employee,
	}
	e.CreatedAt = time.Now()

	if id == "" {
		e.ID = uuid.NewV4().String()
	} else {
		e.ID = id
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *User) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
