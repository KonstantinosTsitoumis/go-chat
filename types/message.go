package types

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id        uuid.UUID
	Text      string
	UserName  string
	CreatedAt time.Time
}
