package model

import (
	"github.com/authink/inkstone"
)

type Staff struct {
	inkstone.Model
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
		Password: inkstone.HashPassword(password),
		Phone:    phone,
		Super:    super,
		Active:   true,
	}
}

func (s *Staff) Reset(password string) {
	s.Password = inkstone.HashPassword(password)
}
