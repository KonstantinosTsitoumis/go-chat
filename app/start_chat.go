package app

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"go-chat/app/tools"
	"go-chat/dependencies/database"
	"os"
)

type ChatDependancies struct {
	Repo database.IRepository
}

func StartChat(deps ChatDependancies) {

	go tools.ReadWorker(deps.Repo)

	username := "Anon" + uuid.New().String()[0:5]

	for true {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		text := scanner.Text()

		if text == "post-trade is the real deal" {
			fmt.Println("Quiting chat...")
			break
		}

		deps.Repo.SendMessage(username, text)
	}
}
