# Docker Notes API

A RESTful API for managing notes built with **Go (Gin Framework)** and **PostgreSQL**, containerized with **Docker**.

## 🚀 Features

- Create, Read, Update, Delete (CRUD) notes
- PostgreSQL database with persistent storage
- Docker & Docker Compose for easy deployment
- Clean architecture with separation of concerns (Handler → Service → Repository)
- Environment-based configuration

## 📁 Project Structure

```
docker-notes-api/
├── back-end/
│   ├── main.go              # Application entry point
│   ├── db/
│   │   └── db.go            # Database connection & schema initialization
│   ├── handler/
│   │   ├── note_handler.go  # HTTP request handlers
│   │   └── response.go      # Standardized response helpers
│   ├── model/
│   │   └── note.go          # Note data model
│   ├── repository/
│   │   └── note_repository.go  # Database queries
│   ├── service/
│   │   └── note_service.go  # Business logic
│   ├── Dockerfile           # Go build configuration
│   ├── go.mod               # Go module dependencies
│   └── go.sum               # Dependency checksums
├── docker-compose.yml       # Docker services orchestration
├── .env                     # Environment variables (not in repo)
└── README.md
```

## 🛠️ Tech Stack

- **Language:** Go 1.22
- **Framework:** Gin Web Framework
- **Database:** PostgreSQL 15
- **Containerization:** Docker & Docker Compose

## 📦 API Endpoints

| Method | Endpoint         | Description           |
|--------|------------------|-----------------------|
| GET    | `/`              | Health check          |
| POST   | `/notes`         | Create a new note     |
| GET    | `/notes`         | Get all notes         |
| GET    | `/notes/:id`     | Get note by ID        |
| PUT    | `/notes/:id`     | Update a note         |
| DELETE | `/notes/:id`     | Delete a note         |

### Request/Response Examples

**Create Note**
```bash
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "My Note", "content": "Note content here"}'
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "My Note",
    "content": "Note content here",
    "created_at": "2026-03-25T10:00:00Z",
    "updated_at": "2026-03-25T10:00:00Z"
  },
  "error": null
}
```

**Get All Notes**
```bash
curl http://localhost:8080/notes
```

## ⚙️ Configuration

Create a `.env` file in the project root:

```env
# Database
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=notes_db
DB_HOST=db
DB_PORT=5432

# API
PORT=8080
```

## 🏃 Getting Started

### Prerequisites

- Docker & Docker Compose installed
- Or Go 1.22+ for local development

### Running with Docker (Recommended)

```bash
# Build and start all services
docker-compose up --build

# Run in detached mode
docker-compose up -d --build
```

The API will be available at `http://localhost:8080`

### Running Locally

```bash
# Install dependencies
cd back-end
go mod download

# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=notes_db
export PORT=8080

# Run the server
go run main.go
```

## 🐳 Docker Commands

```bash
# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Stop and remove volumes (⚠️ deletes data)
docker-compose down -v

# Restart services
docker-compose restart

# Rebuild and restart
docker-compose up --build
```

## 📝 Note Model

```go
type Note struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

## 🧪 Testing the API

```bash
# Health check
curl http://localhost:8080/

# Create note
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "Test", "content": "Test content"}'

# Get all notes
curl http://localhost:8080/notes

# Get note by ID
curl http://localhost:8080/notes/1

# Update note
curl -X PUT http://localhost:8080/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated", "content": "New content"}'

# Delete note
curl -X DELETE http://localhost:8080/notes/1
```

## 📄 License

MIT

## 👨‍💻 Author

Developed as a Docker-based notes API demonstration.
