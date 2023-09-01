package models

import (
	"gorm.io/gorm"
)

type DefaultModel struct {
	gorm.Model `gorm:"embedded"`
}
