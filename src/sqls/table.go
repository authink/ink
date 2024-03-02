package sqls

type tableName struct {
	App       string
	Staff     string
	AuthToken string
	Group     string
	Dept      string
	DeptLevel string
	DeptStaff string
	Log       string
}

var table = &tableName{
	App:       "s_apps",
	Staff:     "s_staffs",
	AuthToken: "s_auth_tokens",
	Group:     "s_groups",
	Dept:      "s_departments",
	DeptLevel: "s_dept_levels",
	DeptStaff: "s_dept_staffs",
	Log:       "s_logs",
}
