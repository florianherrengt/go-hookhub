package datasource

import (
	"github.com/florianherrengt/hubhook/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB connection
var DB *gorm.DB

// ConnectDataBase adds the connection to DB
func ConnectDataBase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// database, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.HookEvent{})

	DB = database

}
