# Link Shortener

A simple and fast URL shortener service built with Go and Redis.

## Features

- Shorten long URLs into compact short links
- Redirect short links to original URLs
- 30-day expiration for stored links

## Getting Started

### Prerequisites

- Docker and Docker Compose

### Installation & Running

1. Clone the repository:
```bash
git clone https://github.com/dusiburg/link-shortener
cd link-shortener
```

2. Start the application with docker compose:
```bash
docker compose up
```

The service will be available at `http://localhost:8080`

## API Endpoints

### Create Short Link

**Request:**
```
GET /l/{original-url}
```

**Example:**
```bash
curl "http://localhost:8080/l/https://example.com/very/long/url"
```

**Response:**
```json
{
  "status": "success",
  "short_link": "http://localhost:8080/r/abc123xyz",
  "original_link": "https://example.com/very/long/url",
  "expires_in": "30d"
}
```

### Redirect to Original Link

**Request:**
```
GET /r/{short-id}
```

**Example:**
```bash
curl "http://localhost:8080/r/abc123xyz"
```

This will redirect to the original URL.

### Health Check

**Request:**
```
GET /ping
```

**Response:**
```json
{
  "status": "success"
}
```

## Configuration

Environment variables can be set via `.env` file or docker compose:

- `SERVER_HOST` - Server host (default: `localhost`)
- `SERVER_PORT` - Server port (default: `8080`)
- `REDIS_URL` - Redis connection URL (required)

## Project Structure

```
.
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── db/                  # Database connections
│   ├── handlers/            # HTTP request handlers
│   ├── http/                # HTTP server setup
│   ├── repositories/        # Data access layer
│   └── services/            # Business logic
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile               # Multi-stage Docker build
└── README                   # This file
```

## Development

To build locally without Docker:

```bash
go build -o bin/application ./cmd
./bin/application
```

Make sure Redis is running before starting the application.