package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"id,omitempty" valid:"uuid"`
	CreatedAt time.Time `json:"created_at,omitempty" valid:"required"`
}
