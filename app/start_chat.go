package app

import (
	"fmt"
	"go_project/dependencies/database"
	"time"
)

type ChatDependancies struct {
	Repo database.IRepository
}

func StartChat(deps ChatDependancies) {

	//deps.Repo.SendMessage("Test", "hello")
	getMTime := time.Date(2023, time.September, 01, 7, 57, 0, 0, time.UTC)

	fmt.Println(getMTime)

	messages := deps.Repo.GetMessages(getMTime)

	//fmt.Printf("test")

	for _, m := range messages {
		fmt.Println(m.CreatedAt, "  ", m.Id, "  ", m.UserName+": "+m.Text)
	}
}
