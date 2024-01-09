package models

import "github.com/authink/ink.go/src/util"

type Staff struct {
	Model
	Email     string
	Password  string
	Phone     string
	Super     bool
	Active    bool
	Departure bool
}

func NewStaff(email, password, phone string, super bool) (staff *Staff, err error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return
	}

	staff = &Staff{
		Email:    email,
		Password: hashedPassword,
		Phone:    phone,
		Super:    super,
	}
	return
}
