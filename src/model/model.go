package model

import "time"

type Model struct {
	Id        uint32
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
