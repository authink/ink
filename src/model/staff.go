package model

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

func NewStaff(email, password, phone string, super bool) *Staff {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		panic(err)
	}

	return &Staff{
		Email:    email,
		Password: hashedPassword,
		Phone:    phone,
		Super:    super,
	}
}
