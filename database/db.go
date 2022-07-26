package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hermannhahn/go-api/models"

	"github.com/tkanos/gonfig"
)

// DB is the database connection
var (
	DB  *gorm.DB
	err error
)

// APIConfig returns the configuration for the API
func APIConfig() models.APIConfiguration {
	configuration := models.APIConfiguration{}
	fileName := "api_config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

// Connect returns a connection to the database
func Connect() {
	user := string(os.Getenv("POSTGRES_USER"))
	password := string(os.Getenv("POSTGRES_PASSWORD"))
	host := string(os.Getenv("POSTGRES_HOST"))
	port := string(os.Getenv("POSTGRES_PORT"))
	dbname := string(os.Getenv("POSTGRES_DB"))

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{})
	DB = db
}
