package main

import (
	"go-chat/app"
	"go-chat/dependencies/database/postgres"
)

func main() {
	repo := postgres.NewRepository()
	app.StartChat(app.ChatDependancies{Repo: repo})
}
