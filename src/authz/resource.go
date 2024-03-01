package authz

import (
	"fmt"
	"net/http"

	"github.com/authink/ink.go/src/env"
)

type Obj struct {
	Name     string
	NeedRoot bool
	Acts     []string
	AppName  string
}

func (obj *Obj) Resource() string {
	return fmt.Sprintf("%s/%s", obj.AppName, obj.Name)
}

func (obj *Obj) Support(act string) bool {
	for _, v := range obj.Acts {
		if v == act {
			return true
		}
	}
	return false
}

var (
	Apps = Obj{
		Name: "apps",
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
		AppName: env.AppNameAdmin(),
	}
	Staffs = Obj{
		Name: "staffs",
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
		AppName: env.AppNameAdmin(),
	}
	Tokens = Obj{
		Name: "tokens",
		Acts: []string{
			http.MethodGet,
			http.MethodDelete,
		},
		AppName: env.AppNameAdmin(),
	}
	Groups = Obj{
		Name:     "groups",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
		AppName: env.AppNameAdmin(),
	}
	Groupships = Obj{
		Name:     "groupships",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
		AppName: env.AppNameAdmin(),
	}
	Permissons = Obj{
		Name:     "permissions",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
		},
		AppName: env.AppNameAdmin(),
	}
	Policies = Obj{
		Name:     "policies",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
		AppName: env.AppNameAdmin(),
	}
	Departments = Obj{
		Name: "departments",
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
		AppName: env.AppNameAdmin(),
	}
	Logs = Obj{
		Name:     "logs",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
		},
		AppName: env.AppNameAdmin(),
	}
)

func ObjList() []Obj {
	return []Obj{Apps, Staffs, Tokens, Groups, Groupships, Permissons, Policies}
}

func GetObj(appName, name string) *Obj {
	for _, v := range ObjList() {
		if v.AppName == appName && v.Name == name {
			return &v
		}
	}
	return nil
}
