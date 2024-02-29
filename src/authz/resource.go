package authz

import (
	"net/http"
)

type Obj struct {
	Name     string
	NeedRoot bool
	Acts     []string
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
	}
	Staffs = Obj{
		Name: "staffs",
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
	}
	Tokens = Obj{
		Name: "tokens",
		Acts: []string{
			http.MethodGet,
			http.MethodDelete,
		},
	}
	Groups = Obj{
		Name:     "groups",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
		},
	}
	Groupships = Obj{
		Name:     "groupships",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
	}
	Policies = Obj{
		Name:     "policies",
		NeedRoot: true,
		Acts: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
		},
	}
)

func List() []Obj {
	return []Obj{Apps, Staffs, Tokens, Groups}
}

func Get(name string) *Obj {
	for _, v := range List() {
		if v.Name == name {
			return &v
		}
	}
	return nil
}
