package app

import (
	"fmt"
	"go-chat/app/tools"
	"go-chat/dependencies/database"
)

type ChatDependancies struct {
	Repo database.IRepository
}

func StartChat(deps ChatDependancies) {

	go tools.ReadWorker(deps.Repo)

	username := "Anon"
	for true {
		var text string

		_, _ = fmt.Scanln(&text)

		deps.Repo.SendMessage(username, text)
	}
}
