package env

import "github.com/authink/inkstone"

func AppNameAdmin() string {
	appNameAdmin := "admin.dev"
	inkstone.GetEnvString("APP_NAME_ADMIN", &appNameAdmin)
	return appNameAdmin
}
