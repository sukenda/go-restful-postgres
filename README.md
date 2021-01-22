# Go Restful Api

Go restful api build for learning go using echo and postgresql and gorm

# Relationships
* User -> Employee using One-to-One
* Employee -> Department using Many To Many
* Department -> Project using One-to-Many

# API Spec

## Register account

Request :

- Method : POST
- Endpoint : `/api/register`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
  "username": "string",
  "password": "string",
  "email": "string",
  "phone": "string"
}
```

Response :

```json 
{
  "code": "int",
  "status": "string",
  "data": {
    "id": "string",
    "username": "string",
    "password": "string",
    "email": "string",
    "phone": "string"
  }
}
```

## Login

Request :

- Method : POST
- Endpoint : `/api/login`
- Header :
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json 
{
  "username": "string",
  "password": "string"
}
```

Response :

```json 
{
  "code": "int",
  "status": "string",
  "data": {
    "access_token": "string",
    "user": {
      "id": "string",
      "username": "string",
      "email": "string",
      "phone": "string"
    }
  }
}
```

## Create Product

Request :

- Method : POST
- Endpoint : `/api/products`
- Header :
    - Authorization: Bearer xxxxx
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
  "name": "string",
  "price": "int64",
  "quantity": "int32"
}
```

Response :

```json
{
  "code": "int",
  "status": "string",
  "data": {
    "id": "string",
    "name": "string",
    "price": "int64",
    "quantity": "int32"
  }
}
```

## Update Product

Request :

- Method : PUT
- Endpoint : `/api/products`
- Header :
    - Authorization: Bearer xxxxx
    - Content-Type: application/json
    - Accept: application/json
- Body :

```json
{
  "id": "string",
  "name": "string",
  "price": "int64",
  "quantity": "int32"
}
```

Response :

```json
{
  "code": "int",
  "status": "string",
  "data": {
    "id": "string",
    "name": "string",
    "price": "int64",
    "quantity": "int32"
  }
}
```

## Get Product By ID

Request :

- Method : GET
- Endpoint : `/api/products/:id`
- Header :
    - Authorization: Bearer xxxxx
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
  "code": "int",
  "status": "string",
  "data": {
    "id": "string",
    "name": "string",
    "price": "int64",
    "quantity": "int32"
  }
}
```

## Get Product All

Request :

- Method : GET
- Endpoint : `/api/products`
- Header :
    - Authorization: Bearer xxxxx
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
  "code": "int",
  "status": "string",
  "data": [
    {
      "id": "string",
      "name": "string",
      "price": "int64",
      "quantity": "int32"
    },
    {
      "id": "string",
      "name": "string",
      "price": "int64",
      "quantity": "int32"
    }
  ]
}
```

## Delete Product By ID

Request :

- Method : DELETE
- Endpoint : `/api/products/:id`
- Header :
    - Authorization: Bearer xxxxx
    - Content-Type: application/json
    - Accept: application/json

Response :

```json
{
  "code": "int",
  "status": "string",
  "data": "string"
}
```

# Reference

https://github.com/khannedy/golang-clean-architecture
