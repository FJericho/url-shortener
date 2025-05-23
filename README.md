# URL Shortener
![Project Status](https://img.shields.io/badge/status-in%20progress-yellow)

A high-performance URL Shortener API built with **Golang**, designed using **Clean Architecture**, and powered by **Fiber**, **GORM**, **PostgreSQL**, and **Swagger** for API documentation. This service lets users shorten long URLs and redirect them through a unique short code.

---

## Tech Stack

| Layer        | Technology                  |
| ------------ | --------------------------- |
| Language     | Go (Golang)                 |
| Framework    | Fiber                       |
| ORM          | GORM                        |
| DB           | PostgreSQL                  |
| Config       | Viper                       |
| Logger       | Logrus                      |
| Docs         | Swagger (swaggo/swag)       |
| Container    | Docker, Docker Compose      |

---

## Installation

1. Clone the repo
```sh
https://github.com/FJericho/url-shortener.git
```
2. Navigate to the root directory
```sh
cd url-shortener
```
3. Create a .env file in the root directory
```sh
touch .env
# or
echo > .env
```
4. Add environment variable
```sh
APP_PORT=Application port(example: 3000)

DB_HOST= Database host (example: localhost)
DB_PORT= Database port (example: 5432)
DB_USER= Database username (example: postgres)
DB_PASSWORD= Database user password (example: postgres)
DB_NAME=Name of the database to be used (example: shortener_db)
```
5. Add Docker Compose
```sh
services:
  postgres:
    image: postgres:alpine
    restart: always
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
    ports:
      - "${DB_PORT}:5432"
    volumes:
      - db:/var/lib/postgresql/data

volumes:
  db:
    driver: local
```

6. Run Docker Compose
```sh
docker-compose up -d
```

7. Run the App Locally
```sh
go run cmd/main.go
```

## API Documentation

The API documentation can be found 
```sh
http://localhost:3000/swagger/index.html
```

## API Endpoints

### 1. Shorten URL

**Endpoint:**  
`POST /shorten`

**Description:**  
Shortens a given long URL by generating a unique short code.

#### Request Body:
```json
{
  "original": "https://example.com/very/long/url"
}
```

#### Success Response:
**Status:** ``200 OK``

```json
{
  "message": "Short url successfully",
  "data": {
    "original": "https://example.com/very/long/url",
    "short_code": "a1B2c3"
  }
}
```

#### Error Response:
**Status:** `400 Bad Request`

```json
{
  "message": "Invalid request. Please check your input."
}
```

### 2. Redirect to Original URL
**Endpoint:**  
`GET /:short_code`

**Description:**  
Redirects the user to the original URL based on the short code.

#### Example:
```sh
GET localhost:3000/a1B2c3
```

#### Response:
**Status:** ``301 Moved Permanently`` (redirects to the original URL)

#### Error Response:
**Status:** `404 Not Found`

```json
{
  "message": "Short URL not found"
}

```

### 3. Get Original URL from Short Code

**Endpoint:**  
`GET /api/url/:short_code`

**Description:**  
Returns the original URL and its short code without redirection.

#### Example:
```sh
GET localhost:3000/api/url/a1B2c3
```

#### Success Response:
**Status:** ``200 OK``

```json
{
  "message": "Get original url successfully",
  "data": {
    "original": "https://example.com/very/long/url",
    "short_code": "a1B2c3"
  }
}
```

#### Error Response:
**Status:** `400 Bad Request`

```json
{
  "message": "Short URL not found"
}
```

### Notes
>- **Redis is NOT used** in this project. 
>- Short codes are **automatically generated** using a random mix of letters and numbers.
