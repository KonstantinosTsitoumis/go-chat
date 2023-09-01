package postgres

import (
	"fmt"
	"go_project/config"
	"go_project/dependencies/database/postgres/models"
	"go_project/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) SendMessage(username string, text string) types.Message {
	message := models.Message{
		Text:      text,
		UserName:  username,
		CreatedAt: time.Now(),
	}

	r.db.Create(&message)

	return types.Message{
		Id:        message.Id,
		UserName:  message.UserName,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}
}

func (r Repository) GetMessages(fromDate time.Time) []types.Message {
	var messages []models.Message

	r.db.Where("created_at > ?", fromDate).Find(&messages)

	var result []types.Message

	// TODO: to tool function
	for _, m := range messages {

		modelToType := types.Message{
			Id:        m.Id,
			Text:      m.Text,
			UserName:  m.UserName,
			CreatedAt: m.CreatedAt,
		}

		result = append(result, modelToType)
	}

	return result
}

func NewRepository() *Repository {
	db, connectionErr := gorm.Open(postgres.Open(config.DB_URL), &gorm.Config{})

	if connectionErr != nil {
		panic("Could not connect to database")
	}

	migrationErr := db.AutoMigrate(&models.Message{})

	if migrationErr != nil {
		panic("Migration failed")
	}

	fmt.Println("Connected and migrated")

	return &Repository{db: db}
}
