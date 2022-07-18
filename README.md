# go-api-gin

Simple API-REST in Go getting data from PostgreSQL.

# Install

1. Rename db_config.example.json to db_config.json
2. docker-compose up
3. go run .

# Routes

    Show all products.
    GET -> /products
    
    Search product by name, description or price.
	GET -> /products/s/:query
    
    Show product by ID.
	GET -> /products/:id
    
    Add new product
	POST -> /products
        {
            "name": "PRODUCT NAME",
            "description": "PRODUCT DESCRIPTION",
            "image": "IMAGE URL",
            "price": FLOAT,
            "quantity": INT,
            "active": "true or false"
        }
    
    Show all products
	DELETE -> /products/:id
    
    Show all products
	PATCH -> /products/:id
