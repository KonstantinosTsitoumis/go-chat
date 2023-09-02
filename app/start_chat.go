package app

import (
	"go-chat/dependencies/database"
)

type ChatDependencies struct {
	Repo database.IRepository
}

func StartChat(deps ChatDependencies) {
	go ReadWorker(deps.Repo)

	username := CreateAnonymousUsername()

	StartChatSender(deps.Repo, username)
}
