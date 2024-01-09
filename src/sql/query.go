package sql

import "fmt"

type query struct {
	InsertApp string
}

var Query = &query{
	InsertApp: fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret)", Table.App),
}
