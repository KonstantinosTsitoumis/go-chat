package database

import "go_project/types"

type IRepository interface {
	GetMessages() []types.Message
	SendMessage(username string, text string) types.Message
}
