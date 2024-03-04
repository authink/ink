package tables

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/sql"
)

const (
	TB_APPS        = "s_apps"
	TB_STAFFS      = "s_staffs"
	TB_TOKENS      = "s_auth_tokens"
	TB_GOUPS       = "s_groups"
	TB_DEPTS       = "s_departments"
	TB_DEPT_LEVELS = "s_dept_levels"
	TB_DEPT_STAFFS = "s_dept_staffs"
	TB_LOGS        = "s_logs"
)

var (
	App       app
	Staff     staff
	AuthToken authToken
	Group     group
	Dept      dept
	DeptLevel deptLevel
	DeptStaff deptStaff
	Log       log
)

func init() {
	sql.Bind(&App, &models.App{})
	sql.Bind(&Staff, &models.Staff{})
	sql.Bind(&AuthToken, &models.AuthToken{})
	sql.Bind(&Group, &models.Group{})
	sql.Bind(&Dept, &models.Department{})
	sql.Bind(&DeptLevel, &models.DeptLevel{})
	sql.Bind(&DeptStaff, &models.DeptStaff{})
	sql.Bind(&Log, &models.Log{})
}
