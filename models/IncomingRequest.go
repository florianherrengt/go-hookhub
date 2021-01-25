package models

import (
	"os"

	"github.com/google/uuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// IncomingRequest from external service
type IncomingRequest struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Payload string
}

func (i *IncomingRequest) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}

var DB *gorm.DB

func ConnectDataBase() {
	// database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	database, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&IncomingRequest{})

	DB = database
}
