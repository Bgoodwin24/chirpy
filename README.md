# **Chirpy API**

Chirpy is a social network similar to Twitter(X). Chirpy contains the back-end logic to receive requests and create responses to them, any desired front-end logic will need to be added to it (ie. JS, CSS, etc). 

[![Go Version](https://img.shields.io/github/go-mod/go-version/Bgoodwin24/chirpy)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

# **Highlights**
* Built with clean Go architecture
* RESTful API design
* Real-world authentication implementation
* Scalable database structure
* Production-ready security practices

# Table of Contents
- [Highlights](#highlights)
- [Quick Start](#quick-start)
- [Supported Features](#supported-features)
- [Setup](#setup)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Database Setup](#database-setup)
  - [Running the API](#running-the-api)
- [Required Dependencies](#required-dependencies)
- [Authentication Headers Required](#authentication-headers-required)
- [Resources](#resources)
- [API Endpoints](#api-endpoints)
- [Example Responses](#example-responses)
- [License](#license)

# :rocket: Quick Start
1. Install PostgreSQL and create a database
2. Clone repository (`git clone github.com/Bgoodwin24/chirpy`) and install Go dependencies (see Required Dependencies)
3. Set up environment variables (see Configuration section)
4. Run database migrations (see Database Setup section)
5. Start the server `go run .`

## **Prerequisites**
* Install Go 1.23.5 or later
* PostgreSQL installed and running
* Goose database migration tool `go install github.com/pressly/goose/v3/cmd/goose@latest`

# **Supported Features**
* User authentication/authorization
* CRUD operations for chirps (posts)
* Token-based security
* Password hashing
* Premium user features (Chirpy Red)
* Database persistence

# :gear: **Setup**

## **Configuration**
* Create a `.env` file in the root directory with:
    * `PORT` - Server Port (default:8080)
    * `JWT_SECRET` - Secret key for JWT tokens
    * `DB_URL` - PostgreSQL connection string
    * `POLKA_KEY` - API Key to third party payment processor
    ## Example .env file
        PORT=8080
        JWT_SECRET=your-secret-here
        DB_URL=postgres://user:password@localhost:5432/dbname
        POLKA_KEY=your-polka-key-here

## **Database Setup**
1. Create a PostgreSQL database
2. Run the migrations. Syntax:
```bash
goose -dir sql/schema postgres "your-connection-string-here" up
```
3. Verify database connection

## **Running the API**
1. Install dependencies: `go mod download`
2. Start server: `go run .`
3. API will be available at `http://localhost:8080`

# **Required Dependencies**
* github.com/golang-jwt/jwt/v5 v5.2.1
* github.com/google/uuid v1.6.0
* github.com/joho/godotenv v1.5.1
* github.com/lib/pq v1.10.9
* golang.org/x/crypto v0.32.0

# Authentication Headers Required
* JWT Authorization: `Authorization: Bearer \<token\>`
* Admin API Key: `ApiKey: \<api_key\>`

# Resources

## **User**

```json
{
    "ID":          "uuid",
	"CreatedAt":   "timestamp",
	"UpdatedAt":   "timestamp",
	"Email":       "user@example.com",
	"IsChirpyRed": "bool",
}
```

## **Chirp**

```json
{
	"ID":        "uuid",
	"CreatedAt": "timestamp",
	"UpdatedAt": "timestamp",
	"UserID":    "uuid",
	"Body":      "body",
}
```

# API Endpoints

## Authentication
* `POST /api/users`
    * Creates a new user account
    * Request: `{"email": "user@example.com", "password": "userspassword"}`
    * Returns: User object

* `POST /api/login`
    * Authenticates a user
    * Request: `{"email": "user@example.com", "password": "userspassword"}`
    * Returns: JWT access token and refresh token

* `POST /api/refresh`
    * Refreshes an access token using a refresh token
    * Requires: Valid refresh token in Authorization header
    * Returns: New access token

* `POST /api/revoke`
    * Revokes a refresh token
    * Requires: Valid refresh token in authorization header

## Chirps
* `POST /api/chirps`
    * Creates a new chirp
    * Requires: Valid access token
    * Request: `{"body": "Hello World!"}`
    * Returns: Chirp object

* `GET /api/chirps`
    * Retrieves all chirps
    * Optional author_id parameter

* `GET /api/chirps/{chirpID}`
    * Retrieves a specific chirp

* `DELETE /api/chirps/{chirpID}`
    * Deletes a chirp
    * Requires: Valid access token from chirp author

## Users
* `PUT /api/users`
    * Updates user information
    * Requires: Valid access token

## Webhooks
* `POST /api/polka/webhooks`
    * Handles webhook from Polka for Chirpy Red upgrades
    * Requires: Valid Polka API Key

## Authorization/Admin
* `Get /app/`
    * Serves the frontend static files

* `GET /api/healthz`
    * Health check endpoint

* `POST /admin/reset`
    * Resets the database
    * Requires: Valid API Key

* `GET /admin/metrics`
    * Returns server metrics
    * Requires: Valid API Key

# Example Responses

## Successful login:
```json
{
    "token": "jwt-token-here",
    "refresh_token": "refresh-token-here"
}
```

## Create User
```json
{
    "id": "uuid-here",
    "email": "user@example.com",
    "created_at": "2024-01-01T00:00:00Z",
    "is_chirpy_red": false
}
```

## Create Chirp
```json
{
    "id": "uuid-here",
    "body": "Hello World!",
    "author_id": "user-uuid-here",
    "created_at": "2024-01-01T00:00:00Z"
}
```

## Status Codes
* 200 - OK: Request successful
* 201 - Created: Resource successfully created
* 400 - Bad Request: Invalid input/parameters
* 401 - Unauthorized: Missing or invalid authentication
* 403 - Forbidden: Valid auth but insufficient permissions
* 404 - Not Found: Resource doesn't exist
* 500 - Internal Server Error: Server-side error

## Error Response Format
```json
{
    "error": "Error message description"
}
```

## Request/Response Limits
* Chirp body: Maximum 140 characters
* Profanity filter replaces the following words with ****:
    * "kerfuffle"
    * "sharbert"
    * "fornax"
    :warning:These are just placeholders as an example:warning:

## Testing
1. Run test suite with: `go test ./...`

# License
* This project is licensed under the [MIT-License](LICENSE).
