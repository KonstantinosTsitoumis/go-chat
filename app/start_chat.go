package app

import (
	"fmt"
	"go_project/dependencies/database"
)

type ChatDependancies struct {
	Repo database.IRepository
}

func StartChat(deps ChatDependancies) {

	//deps.Repo.SendMessage("Test", "hello")
	messages := deps.Repo.GetMessages()

	//fmt.Printf("test")

	for _, m := range messages {
		fmt.Println(m.CreatedAt, "  ", m.Id, "  ", m.UserName+": "+m.Text)
	}
}
