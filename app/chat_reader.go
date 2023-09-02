package app

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
			readFromDate = GetLastSliceElement(messages).CreatedAt
		}

		time.Sleep(2 * time.Millisecond)
	}
}
