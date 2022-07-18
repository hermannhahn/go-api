# go-api-gin

Simple API-REST running tasks on the PostgreSQL database.

# Install

1. Rename db_config.example.json to db_config.json
[Update contents to use other database and skip step 2]
2. Run > docker-compose up
3. Run > go run .

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
    
    Delete product
	DELETE -> /products/:id
    
    Update product
	PATCH -> /products/:id
        {
            "name": "PRODUCT NAME",
            "description": "PRODUCT DESCRIPTION",
            "image": "IMAGE URL",
            "price": FLOAT,
            "quantity": INT,
            "active": "true or false"
        }
