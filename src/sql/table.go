package sql

type table struct {
	App string
}

var Table = &table{
	App: "s_apps",
}