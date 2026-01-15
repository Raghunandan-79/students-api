# Students API

A simple CRUD API built in Go that manages student records using SQLite as a database

## Features

- Create a student
- Get a student by ID
- Get all students
- Update a student by ID
- Delete a student by ID

## Tech Stack

- **Language:** Go 1.25+
- **Database:** SQLite
- **Routing:** `net/http`
- **Validation:** using `go-playground/validator` 
- **Config:** YAML (`local.yaml`)
- **Logging:** slog

> Note: Only basic struct validation is used (no advanced business validation yet).

## Project Structure

```bash
students-api/
├── cmd/
│   └── students-api/
│       └── main.go
├── config/
│   └── local.yaml
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── http/
│   │   └── handlers/
│   │       └── student.go
│   ├── storage/
│   │   ├── sqlite/
│   │   │   └── sqlite.go
│   │   └── storage.go
│   ├── types/
│   │   └── types.go
│   └── utils/
│       └── response/
│           └── response.go
├── storage/
│   └── storage.db
├── .gitignore
├── go.mod
└── go.sum
```

## Running the Project

```bash
go run ./cmd/students-api/main.go -config ./config/local.yaml
```

## Example local.yaml:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

## API Endpoints

1. Create Student
POST /api/atudents

- Body:

```json
{
  "name": "Raghu",
  "email": "raghu@example.com",
  "age": 21
}
```

