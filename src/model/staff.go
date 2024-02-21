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
	hashedPassword, err := inkstone.HashPassword(password)
	if err != nil {
		panic(err)
	}

	return &Staff{
		Email:    email,
		Password: hashedPassword,
		Phone:    phone,
		Super:    super,
		Active:   true,
	}
}
