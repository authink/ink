package sql

type tableName struct {
	App       string
	Staff     string
	AuthToken string
	Group     string
}

var table = &tableName{
	App:       "s_apps",
	Staff:     "s_staffs",
	AuthToken: "s_auth_tokens",
	Group:     "s_groups",
}
