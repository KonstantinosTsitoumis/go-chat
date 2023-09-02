package database

import (
	"go-chat/types"
	"time"
)

type IRepository interface {
	GetMessages(fromDate time.Time) *[]types.Message
	SendMessage(username string, text string)
}
