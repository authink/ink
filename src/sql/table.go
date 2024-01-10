package sql

type table struct {
	App       string
	Staff     string
	AuthToken string
}

var Table = &table{
	App:       "s_apps",
	Staff:     "s_staff",
	AuthToken: "s_auth_tokens",
}
