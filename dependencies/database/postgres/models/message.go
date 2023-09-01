package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Text      string
	UserName  string
	CreatedAt time.Time
}
