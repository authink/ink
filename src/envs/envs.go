package envs

import "github.com/authink/inkstone/env"

func AppNameAdmin() string {
	appNameAdmin := "admin.dev"
	env.GetString("APP_NAME_ADMIN", &appNameAdmin)
	return appNameAdmin
}
