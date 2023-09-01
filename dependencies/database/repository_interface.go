package database

import (
	"go_project/types"
	"time"
)

type IRepository interface {
	GetMessages(fromDate time.Time) []types.Message
	SendMessage(username string, text string) types.Message
}
