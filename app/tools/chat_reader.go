package tools

import (
	"fmt"
	"go-chat/dependencies/database"
	"time"
)

func ReadWorker(Repo database.IRepository) {
	fmt.Println("Started reading chat")

	readFromDate := time.Now()

	for true {
		messages := Repo.GetMessages(readFromDate)

		for _, m := range messages {
			fmt.Println(m.CreatedAt, m.UserName+": "+m.Text)
		}

		if len(messages) != 0 {
			readFromDate = messages[len(messages)-1].CreatedAt
		}
	}
}
