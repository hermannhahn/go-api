package data

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"go-api-gin/models"

	"github.com/tkanos/gonfig"
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
func Connect() *sql.DB {
	user := getConfig("db").DbUsername
	pass := getConfig("db").DbPassword
	host := getConfig("db").DbHost
	name := getConfig("db").DbName
	connStr := "user=" + user + " dbname=" + name + " password=" + pass + " host=" + host
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

// ShowProducts returns a list of products
func ShowProducts(db *sql.DB) []models.Product {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.Active)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

// ShowProduct returns a product
func ShowProduct(db *sql.DB, id string) models.Product {
	row := db.QueryRow("SELECT * FROM products WHERE \"ID\" = $1", id)
	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.Active)
	if err != nil {
		panic(err)
	}
	return product
}

// SearchProducts returns a list of products with a search term
func SearchProducts(db *sql.DB, search string) []models.Product {
	rows, err := db.Query("SELECT * FROM products WHERE \"Name\" LIKE $1", "%"+search+"%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.Active)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}
