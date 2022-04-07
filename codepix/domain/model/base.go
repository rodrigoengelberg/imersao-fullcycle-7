package model

import (
	"github.com/asaskevich/govalidator"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdateAt  time.Time `json:"update_at" valid:"-"`
}
