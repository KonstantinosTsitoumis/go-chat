package main

import (
	"go_project/app"
	"go_project/dependencies/database/postgres"
)

func main() {
	repo := postgres.NewRepository()
	app.StartChat(app.ChatDependancies{Repo: repo})
}
