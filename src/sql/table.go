package sql

type table struct {
	App   string
	Staff string
}

var Table = &table{
	App:   "s_apps",
	Staff: "s_staff",
}
