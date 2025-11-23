# Chess Opening Analyzer API

A Go-based REST API for analyzing chess openings, built with a modern microservices architecture using Docker containers.

## Architecture

This project consists of three main services:

- **Go API Server**: Main application server built with Go
- **MongoDB Database**: Document database for storing chess data
- **Mongo Express**: Web-based MongoDB admin interface

## Prerequisites

- Docker and Docker Compose
- Git

## Environment Setup

1. Create a `.env` file in the project root with the following variables:

```env
COMPOSE_PROJECT_NAME=chess-opening-analyzer
DB_HOST=database
DB_PORT=27017
DB_USER=root
DB_PASSWORD=rootpass
DB_NAME=goapi
SECRET_KEY=your-secret-key-here
```

## Quick Start

1. Clone the repository:
```bash
git clone git@github.com:E4-Chess-Opening-Analyzer/chess-opening-analyzer-api.git
cd chess-opening-analyzer-api
```

2. Start all services:
```bash
docker-compose up -d
```

3. The API will be available at `http://localhost:8080`

## Services

### API Server (Port 8080)
- Built with Go
- Auto-generates Swagger documentation
- Hot reload with `gow` for development
- API documentation available at `http://localhost:8080/swagger/index.html`

### MongoDB Database (Port 27017)
- MongoDB instance with authentication
- Default credentials: `root/rootpass`
- Database name: `goapi`

### Mongo Express (Port 8081)
- Web-based MongoDB administration tool
- Access at `http://localhost:8081`
- Login credentials: `ugoapi/pgoapi`

## Development

### Project Structure
```
.
├── build/
│   ├── database/
│   │   └── Dockerfile
│   └── go/
│       └── Dockerfile
├── src/
│   ├── controllers/       # HTTP request handlers
│   ├── database/          # Database connection configuration
│   ├── docs/              # Auto-generated Swagger documentation
│   ├── middleware/        # HTTP middleware (auth, cors, etc.)
│   ├── models/            # Data models and database schemas
│   ├── repositories/      # Database access layer
│   ├── routes/            # API route definitions
│   └── main.go            # Application entry point
├── .env                   # Environment variables (not in git)
├── .env.example          # Environment variables template
├── .gitignore            # Git ignore file
├── docker-compose.yaml   # Docker services configuration
└── README.md             # Project documentation
```

### Hot Reload
The Go service uses `gow` for automatic reloading during development. Any changes to the source code will trigger a restart of the application.

### API Documentation
Swagger documentation is automatically generated using `swaggo/swag`. The docs are regenerated on each container start and available at `/swagger/index.html` endpoint.

## Database

The MongoDB instance comes pre-configured with:
- Root user authentication
- Persistent volume storage
- Network isolation within Docker

### Connecting to Database
- **Host**: `database` (within Docker network) or `localhost` (from host)
- **Port**: `27017`
- **Username**: `root`
- **Password**: `rootpass`
- **Database**: `goapi`

## Deployment

### Production Deployment
1. Update environment variables for production
2. Ensure proper security configurations
3. Use production-ready MongoDB setup
4. Configure proper logging and monitoring

### Docker Commands

Start services:
```bash
docker-compose up -d
```

Stop services:
```bash
docker-compose down
```

View logs:
```bash
docker-compose logs -f [service-name]
```

Rebuild services:
```bash
docker-compose up -d --build
```
