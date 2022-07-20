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
func getConfig(params ...string) models.Configuration {
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
	user := getConfig("db").DbUsername
	pass := getConfig("db").DbPassword
	host := getConfig("db").DbHost
	name := getConfig("db").DbName
	port := getConfig("db").DbPort
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + name + " port=" + port + " sslmode=disable"
	con, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	con.AutoMigrate(&models.Product{})
	DB = con
}
