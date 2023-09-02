package app

import (
	"bufio"
	"fmt"
	"go-chat/dependencies/database"
	"os"
)

func StartChatSender(Repo database.IRepository, username string) {
	for true {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		text := scanner.Text()

		if text == "post-trade is the real deal" {
			fmt.Println("Quiting chat...")
			break
		}

		Repo.SendMessage(username, text)
	}
}
