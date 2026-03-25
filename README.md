# Docker Notes API

A RESTful API for managing notes built with **Go (Gin Framework)** and **PostgreSQL**, containerized with **Docker**. Features pagination and search functionality.

## 🚀 Features

- Create, Read, Update, Delete (CRUD) notes
- **Pagination support** for efficient data retrieval
- **Search functionality** by title or content
- PostgreSQL database with persistent storage (Docker volumes)
- Docker & Docker Compose for easy deployment
- Clean architecture with separation of concerns (Handler → Service → Repository)
- Environment-based configuration
- Automatic schema initialization on startup
- Standardized JSON response format

## 📁 Project Structure

```
docker-notes-api/
├── back-end/
│   ├── main.go              # Application entry point & route definitions
│   ├── db/
│   │   └── db.go            # Database connection pool & schema initialization
│   ├── handler/
│   │   ├── note_handler.go  # HTTP request handlers (controllers)
│   │   └── response.go      # Standardized response helpers (Success/Error)
│   ├── model/
│   │   └── note.go          # Note data model with JSON bindings
│   ├── repository/
│   │   └── note_repository.go  # SQL queries & database operations
│   ├── service/
│   │   └── note_service.go  # Business logic & validation
│   ├── Dockerfile           # Multi-stage Go build configuration
│   ├── go.mod               # Go module dependencies
│   └── go.sum               # Dependency checksums
├── docker-compose.yml       # Docker services orchestration (PostgreSQL + API)
├── .env                     # Environment variables (not in repo)
└── README.md
```

## 🛠️ Tech Stack

| Component      | Technology          |
|----------------|---------------------|
| **Language**   | Go 1.22             |
| **Framework**  | Gin Web Framework   |
| **Database**   | PostgreSQL 15       |
| **Driver**     | lib/pq              |
| **Container**  | Docker & Compose    |

## 🏗️ Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Client    │────▶│   Handler   │────▶│   Service   │────▶│ Repository  │
│  (HTTP Req) │     │  (Routing)  │     │  (Logic)    │     │   (SQL)     │
└─────────────┘     └─────────────┘     └─────────────┘     └─────────────┘
                                                                   │
                                                                   ▼
                                                           ┌─────────────┐
                                                           │ PostgreSQL  │
                                                           │   Database  │
                                                           └─────────────┘
```

### Layer Responsibilities

| Layer        | File                      | Responsibility                           |
|--------------|---------------------------|------------------------------------------|
| **Handler**  | `note_handler.go`         | Parse HTTP requests, return responses    |
| **Service**  | `note_service.go`         | Business logic, input validation         |
| **Repository**| `note_repository.go`     | SQL queries, database connections        |
| **Model**    | `note.go`                 | Data structures, JSON schemas            |
| **DB**       | `db.go`                   | Connection pool, schema migrations       |

## 📦 API Endpoints

| Method | Endpoint         | Description           | Request Body                    |
|--------|------------------|-----------------------|---------------------------------|
| GET    | `/`              | Health check          | -                               |
| POST   | `/notes`         | Create a new note     | `{title, content}`              |
| GET    | `/notes`         | Get all notes (paginated) | -                            |
| GET    | `/notes/:id`     | Get note by ID        | -                               |
| PUT    | `/notes/:id`     | Update a note         | `{title, content}`              |
| DELETE | `/notes/:id`     | Delete a note         | -                               |

### Response Format

All responses follow a standardized format:

**Success Response:**
```json
{
  "success": true,
  "data": { ... },
  "error": null
}
```

**Error Response:**
```json
{
  "success": false,
  "data": null,
  "error": "Error message here"
}
```

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

**Get All Notes (with Pagination & Search):**
```bash
# Default: page=1, limit=10
curl "http://localhost:8080/notes"

# Custom pagination
curl "http://localhost:8080/notes?page=2&limit=5"

# Search by title or content
curl "http://localhost:8080/notes?search=meeting"

# Combine pagination and search
curl "http://localhost:8080/notes?page=1&limit=5&search=important"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "data": [
      {
        "id": 5,
        "title": "Meeting Notes",
        "content": "Important discussion points",
        "created_at": "2026-03-25T10:00:00Z",
        "updated_at": "2026-03-25T10:00:00Z"
      }
    ],
    "total": 25,
    "page": 1,
    "limit": 5
  },
  "error": null
}
```

**Get Note by ID:**
```bash
curl http://localhost:8080/notes/1
```

**Update Note:**
```bash
curl -X PUT http://localhost:8080/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated Title", "content": "Updated content"}'
```

**Delete Note:**
```bash
curl -X DELETE http://localhost:8080/notes/1
```

## 📄 Pagination & Query Parameters

### GET /notes

| Parameter | Type   | Default | Description                          |
|-----------|--------|---------|--------------------------------------|
| `page`    | int    | 1       | Page number (1-indexed)              |
| `limit`   | int    | 10      | Number of items per page             |
| `search`  | string | ""      | Search term for title or content     |

### Query Examples

```bash
# First page with 10 items (default)
curl "http://localhost:8080/notes"

# Second page with 5 items
curl "http://localhost:8080/notes?page=2&limit=5"

# Search notes containing "project"
curl "http://localhost:8080/notes?search=project"

# Search with custom pagination
curl "http://localhost:8080/notes?page=1&limit=20&search=todo"
```

### Response Fields

| Field   | Type | Description                      |
|---------|------|----------------------------------|
| `data`  | []Note | Array of notes for current page |
| `total` | int  | Total number of notes matching search |
| `page`  | int  | Current page number              |
| `limit` | int  | Items per page                   |

## 🗄️ Database Schema

```sql
CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

| Column     | Type        | Constraints              |
|------------|-------------|--------------------------|
| id         | SERIAL      | PRIMARY KEY, AUTO INCREMENT |
| title      | TEXT        | NOT NULL                 |
| content    | TEXT        | NOT NULL                 |
| created_at | TIMESTAMPTZ | DEFAULT NOW()            |
| updated_at | TIMESTAMPTZ | DEFAULT NOW()            |

## ⚙️ Configuration

Create a `.env` file in the project root:

```env
# Database Configuration
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=notes_db
DB_HOST=db
DB_PORT=5432

# API Configuration
PORT=8080
```

### Environment Variables

| Variable     | Description              | Default | Required |
|--------------|--------------------------|---------|----------|
| `DB_USER`    | PostgreSQL username      | -       | ✅       |
| `DB_PASSWORD`| PostgreSQL password      | -       | ✅       |
| `DB_NAME`    | Database name            | -       | ✅       |
| `DB_HOST`    | Database host            | `db`    | ✅       |
| `DB_PORT`    | Database port            | `5432`  | ✅       |
| `PORT`       | API server port          | `8080`  | ❌       |

## 🏃 Getting Started

### Prerequisites

- Docker & Docker Compose installed **OR**
- Go 1.22+ for local development

### Running with Docker (Recommended)

```bash
# Build and start all services
docker-compose up --build

# Run in detached mode (background)
docker-compose up -d --build

# View logs
docker-compose logs -f
```

The API will be available at `http://localhost:8080`

### Running Locally

```bash
# Navigate to backend directory
cd back-end

# Install dependencies
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

> **Note:** For local development, ensure PostgreSQL is running on your machine with the configured credentials.

## 🐳 Docker Commands

```bash
# Build and start services
docker-compose up --build

# Start in detached mode
docker-compose up -d

# View logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f backend
docker-compose logs -f db

# Stop services
docker-compose down

# Stop and remove volumes (⚠️ deletes all data)
docker-compose down -v

# Restart services
docker-compose restart

# Rebuild and restart
docker-compose up --build --force-recreate

# Check running containers
docker-compose ps
```

## 📝 Note Model

```go
type Note struct {
    ID        int    `json:"id"`
    Title     string `json:"title" binding:"required"`
    Content   string `json:"content" binding:"required"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

### Validation Rules

| Field     | Requirement | Description          |
|-----------|-------------|----------------------|
| `title`   | Required    | Cannot be empty      |
| `content` | Required    | Cannot be empty      |

## 🧪 Testing the API

```bash
# Health check
curl http://localhost:8080/

# Create note
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "Test", "content": "Test content"}'

# Get all notes (default pagination)
curl http://localhost:8080/notes

# Get notes with pagination
curl "http://localhost:8080/notes?page=2&limit=5"

# Search notes
curl "http://localhost:8080/notes?search=meeting"

# Get note by ID
curl http://localhost:8080/notes/1

# Update note
curl -X PUT http://localhost:8080/notes/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated", "content": "New content"}'

# Delete note
curl -X DELETE http://localhost:8080/notes/1
```

## 🔍 Key Implementation Details

### Repository Layer (`note_repository.go`)

- Uses `database/sql` with `lib/pq` driver
- Prepared statements for SQL queries
- **Pagination with LIMIT and OFFSET**
- **Search with ILIKE (case-insensitive)**
- Total count query for pagination metadata
- Proper resource cleanup with `defer rows.Close()`
- Returns rows affected for delete operations

### Service Layer (`note_service.go`)

- Input validation (title & content required)
- Business logic separation from handlers
- Error message formatting
- **Pagination parameters forwarding**

### Handler Layer (`note_handler.go`)

- Gin context for request/response handling
- JSON binding with validation
- **Query parameter parsing (page, limit, search)**
- Consistent error handling
- Standardized response format with pagination metadata

### Database Layer (`db.go`)

- Connection pooling via `sql.DB`
- Connection health check with `DB.Ping()`
- Automatic table creation on startup
- Uses `TIMESTAMPTZ` for timezone-aware timestamps

## 🐛 Troubleshooting

| Issue                          | Solution                                    |
|--------------------------------|---------------------------------------------|
| Connection refused to database | Ensure DB container is running: `docker-compose ps` |
| Port 8080 already in use       | Change `PORT` in `.env` or stop conflicting service |
| Migration fails                | Check database credentials in `.env`        |
| Container exits immediately    | View logs: `docker-compose logs backend`    |
| No notes returned              | Check search term or try without search filter |

## 📄 License

MIT

## 👨‍💻 Author

Developed as a Docker-based notes API demonstration showcasing clean architecture with Go, PostgreSQL, and Docker.
