package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"

	"github.com/hermannhahn/go-api-gin/models"

	"github.com/tkanos/gonfig"
)

// DB is the database connection
var (
	DB  *gorm.DB
	err error
)

// GetConfig returns the configuration for the application
func GetConfig(params ...string) models.Configuration {
	configuration := models.Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./%s_config.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

// Connect returns a connection to the database
func Connect() {
	config := GetConfig("db")
	dsn := "host=" + config.DbHost + " user=" + config.DbUsername + " password=" + config.DbPassword + " dbname=" + config.DbName + " port=" + config.DbPort + " sslmode=disable"
	con, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	con.AutoMigrate(&models.Product{})
	DB = con
}
