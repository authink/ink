package models

type App struct {
	Model
	Name   string
	Secret string
}

func NewApp(name, secret string) *App {
	return &App{
		Name:   name,
		Secret: secret,
	}
}
