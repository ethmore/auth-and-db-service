Authentication and Database Service
===

This GO project serves as a microservice for [eCommerce](https://github.com/users/ethmore/projects/4) project.


## Service tasks:

- Authenticate and authorize users/sellers
- CRUD operations for 
    
    - Users
    - Sellers
    - Products 
    - User addresses 
    - Cart Info
    - Order Info 
    - Payment Info



# Installation

Ensure GO is installed on your system
```
go mod download
````

```
go run .
```

## Test
```
curl http://localhost:3002/test
```
### It should return:
```
StatusCode        : 200
StatusDescription : OK
Content           : {"message":"OK"}
```

## Example .env file
This file should be placed inside `dotEnv` folder
```
# Password salt for user registration
SALT = exampleSALT

# PostgreSQL Credentials
HOST = localhost
PORT = 5432
USER = postgres
PASSWORD = examplePassword
DB_NAME = eCommerce

# MongoDB URI
MONGODB_URI = mongodb://localhost:27017

# Secret Token for JWT
TOKEN = exampleToken

# Cors URLs
BFF_URL = http://localhost:3001
ADD_PRODUCT_SEARCH_SERVICE = http://localhost:3006/addProduct
```