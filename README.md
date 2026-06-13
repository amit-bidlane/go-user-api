# Go User API

A RESTful API built with Go, GoFiber, PostgreSQL, and SQLC to manage users with dynamically calculated age.

## Tech Stack

- **GoFiber** — HTTP framework
- **PostgreSQL** — Database
- **SQLC** — Type-safe SQL query generation
- **Uber Zap** — Structured logging
- **go-playground/validator** — Input validation
- **Docker** — Containerization

## Project Structure

```text
/cmd/server/main.go
/config/
/db/migrations/
/db/sqlc/<generated>
/internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/
```

## Setup and Run

### Using Docker (Recommended)

1. Make sure Docker is installed and running
2. Clone the repository
3. Run:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:3000`

### Running Locally

1. Make sure Go and PostgreSQL are installed
2. Create a database called `userdb`
3. Run the migration:

```bash
psql -U postgres -d userdb -f db/migrations/schema.sql
```

4. Set environment variables:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=userdb
export SERVER_PORT=3000
```

5. Run the app:

```bash
go run cmd/server/main.go
```

## API Endpoints

> **Note:** All `dob` fields must be in `YYYY-MM-DD` format. Example: `"1990-05-10"`

### Create User

```http
POST /users
Content-Type: application/json

{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Get User by ID

```http
GET /users/:id
```

### Update User

```http
PUT /users/:id
Content-Type: application/json

{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

### Delete User

```http
DELETE /users/:id
```

### List All Users (with pagination)

```http
GET /users?page=1&limit=10
```

## Running Tests

```bash
go test ./internal/service/...
```
