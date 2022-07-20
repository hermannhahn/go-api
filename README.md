# go-api-gin

Simple Rest API to learn Go Lang.

# Install

1. Rename db_config.example.json to db_config.json
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
    
    Delete product
	DELETE -> /products/:id
    
    Update product
	PATCH -> /products/:id

# RAW

    {
        "name": "PRODUCT NAME",
        "description": "PRODUCT DESCRIPTION",
        "image": "URL",
        "price": FLOAT,
        "quantity": INT,
        "active": "true or false"
    }
