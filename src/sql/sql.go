package sql

type tableName struct {
	App       string
	Staff     string
	AuthToken string
}

var table = &tableName{
	App:       "s_apps",
	Staff:     "s_staff",
	AuthToken: "s_auth_tokens",
}
