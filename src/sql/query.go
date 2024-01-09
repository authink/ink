package sql

import "fmt"

type query struct {
	InsertApp   string
	InsertStaff string
}

var Query = &query{
	InsertApp: fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret)", Table.App),

	InsertStaff: fmt.Sprintf("INSERT INTO %s (email, password, phone, super) VALUES (:email, :password, :phone, :super)", Table.Staff),
}
