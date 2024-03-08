package models

import (
	"github.com/authink/inkstone/util"
	"github.com/authink/orm/model"
)

// @model
// @db s_staffs
type Staff struct {
	model.Base
	Email     string
	Password  string
	Phone     string
	Super     bool
	Active    bool
	Departure bool
}

func NewStaff(email, password, phone string, super bool) *Staff {
	return &Staff{
		Email:    email,
		Password: util.HashPassword(password),
		Phone:    phone,
		Super:    super,
		Active:   true,
	}
}

func (s *Staff) Reset(password string) {
	s.Password = util.HashPassword(password)
}
