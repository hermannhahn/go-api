package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hermannhahn/go-api-gin/models"

	"github.com/tkanos/gonfig"
)

// DB is the database connection
var (
	DB  *gorm.DB
	err error
)

// DBConfig returns the configuration for the database
func DBConfig() models.Configuration {
	configuration := models.Configuration{}
	fileName := "db_config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

// APIConfig returns the configuration for the API
func APIConfig() models.APIConfiguration {
	configuration := models.APIConfiguration{}
	fileName := "api_config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

// Connect returns a connection to the database
func Connect() {
	config := DBConfig()
	dsn := "host=" + config.DbHost + " user=" + config.DbUsername + " password=" + config.DbPassword + " dbname=" + config.DbName + " port=" + config.DbPort + " sslmode=disable"
	con, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	con.AutoMigrate(&models.Product{})
	DB = con
}
