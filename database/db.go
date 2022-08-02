package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hermannhahn/go-api/models"
)

// DB is the database connection
var (
	DB  *gorm.DB
	err error
)

// Connect returns a connection to the database
func Connect() {

	// Get database connection details from environment variables
	user := string(os.Getenv("DATABASE_USER"))
	password := string(os.Getenv("DATABASE_PASSWORD"))
	host := string(os.Getenv("DATABASE_HOST"))
	port := string(os.Getenv("DATABASE_PORT"))
	dbname := string(os.Getenv("DATABASE_DB"))

	// Create connection string
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"

	// Create a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create the database tables if they don't exist
	db.AutoMigrate(&models.Product{})
	DB = db
}
