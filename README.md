# gokit-init

A powerful CLI tool for generating boilerplate Go web application projects with various architectures, databases, and Docker support.

```
   ___       _    _ _          ___      _ _   
  / __|___  | |__(_) |_ ___   |_ _|_ _ (_) |_ 
 | (_ / _ \ | / /| |  _|___|   | || ' \| |  _|
  \___\___/ |_\_\|_|\__|     |___|_||_|_|\__|
                                              
  Go Project Generator - Fast & Clean
```

## Features

**Multiple Architecture Patterns**
- Simple structure with handlers and domain models
- Clean Architecture with proper layer separation (handler → service → repository → domain)

**Database Support**
- MySQL with configured connection
- PostgreSQL with configured connection
- SQLite with file-based storage

**Docker Ready**
- Dockerfile for production builds with multi-stage setup
- docker-compose.yml with database services pre-configured
- Environment variables properly configured

**Smart Scaffolding**
- Pre-configured project structure
- Sample CRUD code and boilerplate
- Environment configuration files (.env.example)
- Comprehensive README for generated projects
- go.mod with proper module path

## Installation

### From Source

```bash
git clone https://github.com/fauzymadani/gokit-init.git
cd gokit-init
go build -o gokit-init
sudo mv gokit-init /usr/local/bin/
```

Or install directly:

```bash
go install github.com/fauzymadani/gokit-init@latest
```

## Usage

### Basic Command

```bash
gokit-init new <project-name>
```

### Available Flags

| Flag | Description | Options |
|------|-------------|---------|
| `--db` | Database type | `mysql`, `postgres`, `sqlite` |
| `--module` | Go module path | Custom module path (default: `github.com/user/<project-name>`) |
| `--docker` | Include Docker files | Boolean flag |
| `--clean-arch` | Use Clean Architecture | Boolean flag |

### Examples

#### Simple MVC Project

```bash
gokit-init new myapp
```

This generates:
```
myapp/
├── handler/
├── domain/
├── main.go
├── go.mod
├── .env.example
└── README.md
```

#### Clean Architecture with PostgreSQL

```bash
gokit-init new myapp --clean-arch --db postgres
```

This generates:
```
myapp/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── database.go
│   ├── handler/
│   │   └── user_handler.go
│   ├── service/
│   │   └── user_service.go
│   ├── repository/
│   │   └── user_repository.go
│   └── domain/
│       └── user.go
├── pkg/
├── go.mod
├── .env.example
└── README.md
```

#### Full Stack with Docker

```bash
gokit-init new myapp --clean-arch --db mysql --docker --module github.com/username/myapp
```

This generates everything above plus:
- `Dockerfile` - Multi-stage build for Go application
- `docker-compose.yml` - Docker Compose with app and database services

## Generated Project Structure

### Simple Structure

- **handler/** - HTTP handlers and request processing
- **domain/** - Domain models and business entities
- **main.go** - Application entry point

### Clean Architecture

- **cmd/app/** - Application entry point
- **internal/domain/** - Business entities (domain models)
- **internal/repository/** - Data access layer (database operations)
- **internal/service/** - Business logic layer (use cases)
- **internal/handler/** - Presentation layer (HTTP handlers)
- **internal/config/** - Configuration (database, app settings)
- **pkg/** - Public libraries (reusable packages)

## Commands

### Create New Project

```bash
gokit-init new [project-name] [flags]
```

### Show Version

```bash
gokit-init version
```

### Help

```bash
gokit-init --help
gokit-init new --help
```

## What's Generated?

Every generated project includes:

1. **go.mod** - Go module file with custom module path
2. **main.go** - HTTP server with health check endpoint
3. **.env.example** - Environment variables template
4. **README.md** - Project-specific documentation

### Optional Files (based on flags)

- **Database Config** - Connection setup for MySQL/PostgreSQL/SQLite
- **Dockerfile** - Multi-stage Docker build
- **docker-compose.yml** - Complete stack with database
- **Clean Architecture Files** - Sample entities, repositories, services, controllers

## Quick Start After Generation

After generating a project:

```bash
cd your-project
go mod tidy
cp .env.example .env
# Edit .env with your configuration
go run main.go  # or: go run cmd/app/main.go for clean-arch
```

With Docker:

```bash
cd your-project
cp .env.example .env
docker-compose up --build
```

## Environment Variables

Generated projects use these environment variables:

```env
APP_PORT=8080        # Application port
DB_HOST=localhost    # Database host
DB_PORT=3306         # Database port (3306 for MySQL, 5432 for PostgreSQL)
DB_USER=root         # Database user
DB_PASS=             # Database password
DB_NAME=app          # Database name
```

For SQLite projects, add:
```env
DB_PATH=./app.db     # SQLite database file path
```

## Development

### Project Structure (gokit-init itself)

```
gokit-init/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command and version
│   └── new.go             # New project command
├── internal/
│   ├── banner/            # ASCII logo
│   ├── version/           # Version info
│   ├── config/            # Configuration validation
│   └── generator/         # Code generation logic
│       ├── directories.go # Directory creation
│       ├── files.go       # File utilities
│       ├── mainfile.go    # main.go generation
│       ├── database.go    # Database configs
│       ├── docker.go      # Docker files
│       ├── cleanarch.go   # Clean arch files
│       └── generator.go   # Main orchestrator
├── main.go                # Entry point
└── go.mod
```

### Build

```bash
go build -o gokit-init
```

### Test

```bash
# Test basic generation
./gokit-init new test-app

# Test with all features
./gokit-init new test-app --clean-arch --db postgres --docker --module github.com/test/app
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - feel free to use this tool for any purpose.

## Roadmap

Future enhancements:
- [ ] More database options (MongoDB, Redis)
- [ ] REST API templates with routing
- [ ] GraphQL support
- [ ] Authentication/Authorization templates
- [ ] Testing boilerplate
- [ ] CI/CD configuration files
- [ ] Kubernetes manifests
- [ ] Swagger/OpenAPI documentation

## Author

Built with ❤️ for the Go community

---

**Happy Coding!**

