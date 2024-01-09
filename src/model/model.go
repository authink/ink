package model

import "time"

type Model struct {
	Id        uint32
	CreatedAt time.Time
	UpdateAt  time.Time
}
