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

> Note: Ignored from git via .gitignore. This file is required to run the project. Create it manually.

## API Endpoints

1. Create Student:

   - Method: POST
   - URL: `/api/students/create`
   - Body:

   ```json
    {
     "name": "test",
     "email": "test@example.com",
     "age": 21
    }
    ```

2. Get Student by ID:

   - Method: GET
   - URL: `/api/students/{id}`

3. Get All Students:

    - Method: GET
    - URL: `/api/students`

4. Update Student by ID:

    - Method: PUT
    - URL: `/api/students/update/{id}`
    - Body:
    ```json
    {
     "name": "test Updated",
     "email": "test.updated@example.com",
     "age": 22
    }
    ```

5. Delete Student by ID:
    - Method: DELETE
    - URL: `/api/students/deleteById/{id}`

## Database

Create a database file at `storage/storage.db` and run the project to auto-create the required tables.

> Note: Ignored from git via .gitignore

## Testing with Postman

A Postman Collection can be created to test the above API endpoints.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
