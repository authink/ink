package sql

type query interface {
	Insert() string
	Get()    string
}
