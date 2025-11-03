# Markdown Note API

A simple REST API for managing markdown notes with grammar fixing and HTML parsing features.

## Features

- ✍️ CRUD operations for markdown notes
- 🔧 Automatic grammar fixing (multi-language support)
- 📄 Parse markdown to HTML
- 📁 Direct .md file upload
- 💾 SQLite database (gorm)
- 🐳 Docker support

## Tech Stack

- **Language**: Go
- **Database**: SQLite (gorm)
- **Router**: net/http (standard library)
- **Port**: 3000

## Project Structure

```
markdown-note/
├── app/
│   └── main.go        # Application entry point
├── internal/
│   ├── handlers/      # HTTP request handlers
│   ├── routes/        # Route registrations
│   ├── services/      # Business logic
│   ├── repo/          # Database repositories
│   ├── models/        # Data models
│   └── lib/           # Utilities
├── data/
│   └── data.db        # SQLite database
├── docker-compose.yml
└── Dockerfile
```

## Quick Start

### Install

```bash
# Clone repository
git clone https://github.com/chesta132/markdown-note-go.git
cd markdown-note-go
```

### Using Docker (Recommended)

```bash
# Run with docker compose
docker compose up
```

### Manual Setup

```bash
# Install dependencies
go mod download

# Run server
go run app/main.go
```

Server will run on `http://localhost:3000` as default

## API Endpoints

### Notes

#### Get All Notes

```http
GET /notes
```

**Response:**

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": [
    {
      "id": "1",
      "title": "My Note",
      "markdown": "# Hello World",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

#### Get Single Note

```http
GET /notes/{id}
```

**Response:**

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": {
    "id": "1",
    "title": "My Note",
    "markdown": "# Hello World",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Create Note

```http
POST /notes
Content-Type: multipart/form-data
```

**Form Data:**

- `title` (optional): Note title (defaults to filename if not provided)
- `file` (required): Markdown file with .md extension (max 5MB)

**Response:** `201 Created`

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": {
    "id": "1",
    "title": "My Note",
    "markdown": "# Hello World",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Update Note

```http
PUT /notes/{id}
Content-Type: application/json
```

**Body:**

```json
{
  "title": "Updated Title",
  "markdown": "# Updated content"
}
```

**Response:**

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": {
    "id": "1",
    "title": "Updated Title",
    "markdown": "# Updated content",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

#### Delete Note

```http
DELETE /notes/{id}
```

**Response:**

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": {
    "id": "1",
    "title": "My Note",
    "markdown": "# Hello World",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Fix Grammar

```http
PATCH /notes/{id}/fix-grammar?lang=en-US
```

**Query Params:**

- `lang` (required): Language code (see supported languages below)

**Response:**

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": {
    "id": "1",
    "title": "My Note",
    "markdown": "# Corrected markdown content",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

Note: This endpoint automatically updates the note with corrected grammar.

#### Get Parsed HTML

```http
GET /notes/{id}/html
```

**Response:** HTML content (not JSON)

```html
<h1>Hello World</h1>
<p>This is parsed markdown content.</p>
```

## Response Format

All endpoints (except `/notes/{id}/html`) return JSON with this format:

**Success Response:**

```json
{
  "meta": {
    "status": "SUCCESS"
  },
  "data": { ... }
}
```

**Error Response:**

```json
{
  "meta": {
    "status": "ERROR"
  },
  "data": {
    "code": "ERROR_CODE",
    "message": "Human readable error message",
    "details": "Additional error details (optional)"
  }
}
```

### Error Codes

| Code           | Description                                |
| -------------- | ------------------------------------------ |
| `NOT_FOUND`    | Resource not found                         |
| `SERVER_ERROR` | Internal server error                      |
| `CLIENT_ERROR` | Invalid request from client                |
| `BAD_GATEWAY`  | External service error (e.g., grammar API) |

### HTTP Status Codes

- `200 OK` - Successful GET, PUT, PATCH, DELETE
- `201 Created` - Successful POST
- `204 No Content` - Successful operation with no content
- `400 Bad Request` - Invalid client request
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error
- `502 Bad Gateway` - External service error

## Configuration

Edit in `app/main.go`:

```go
const (
    PORT    = "3000"         // Server port
    DB_PATH = "data/data.db" // Database path
)
```

## Validation Rules

- File upload maximum size: **5MB**
- Only accepts files with `.md` extension
- Title field is optional (defaults to filename)
- Grammar fixing requires a valid language code

## Notes

- Database file is automatically created if it doesn't exist
- Grammar fixing requires connection to external service
- All timestamps are in ISO 8601 format
- File content is stored as text in the database

---

**Challenge from:** [roadmap.sh/projects/url-shortening-service](https://roadmap.sh/projects/url-shortening-service)

