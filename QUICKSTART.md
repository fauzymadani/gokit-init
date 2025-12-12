# Quick Start Guide
## Installation
```bash
cd /home/yzuaf/Documents/dev/go/gokit-init
go build -o gokit-init
```
## Usage Examples
### 1. Simple Web App (No Database)
```bash
./gokit-init new my-simple-app --module github.com/myuser/my-simple-app
cd my-simple-app
go mod tidy
go run main.go
```
### 2. Web App with MySQL
```bash
./gokit-init new my-mysql-app --db mysql --module github.com/myuser/my-mysql-app
cd my-mysql-app
go mod tidy
go run main.go
```
### 3. Clean Architecture with PostgreSQL
```bash
./gokit-init new my-clean-app --clean-arch --db postgres --module github.com/myuser/my-clean-app
cd my-clean-app
go mod tidy
go run cmd/app/main.go
```
### 4. Full Stack with Docker
```bash
./gokit-init new my-full-app --clean-arch --db postgres --docker --module github.com/myuser/my-full-app
cd my-full-app
go mod tidy
# Set up environment
cp .env.example .env
# Run with Docker
docker-compose up --build
```
### 5. Microservice with SQLite
```bash
./gokit-init new my-microservice --clean-arch --db sqlite --module github.com/myuser/my-microservice
cd my-microservice
go mod tidy
go run cmd/app/main.go
```
## Project Structure Examples
### Simple MVC
```
my-simple-app/
├── controller/       # HTTP handlers
├── model/           # Data models
├── view/            # Templates
├── main.go          # Entry point
├── go.mod
├── .env.example
└── README.md
```
### Clean Architecture
```
my-clean-app/
├── cmd/
│   └── app/
│       └── main.go           # Entry point
├── internal/
│   ├── entity/               # Domain models
│   │   └── user.go
│   ├── repository/           # Data access
│   │   └── user_repository.go
│   ├── service/              # Business logic
│   │   └── user_service.go
│   ├── controller/           # HTTP handlers
│   │   └── user_controller.go
│   └── config/               # Configuration
│       └── database.go
├── pkg/                      # Public packages
├── go.mod
├── .env.example
└── README.md
```
### With Docker
```
my-full-app/
├── cmd/app/main.go
├── internal/...
├── Dockerfile               # Multi-stage build
├── docker-compose.yml       # App + Database
├── .env.example
└── ...
```
## Environment Variables
Edit `.env` after copying from `.env.example`:
```bash
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=myapp
```
## Available Commands
```bash
# Show help
./gokit-init --help
# Show new command help
./gokit-init new --help
# Show version
./gokit-init version
```
## Flags Reference
| Flag | Short | Description | Example |
|------|-------|-------------|---------|
| `--db` | | Database type | `--db postgres` |
| `--module` | | Go module path | `--module github.com/user/app` |
| `--docker` | | Include Docker files | `--docker` |
| `--clean-arch` | | Use Clean Architecture | `--clean-arch` |
## Next Steps After Generation
1. **Install dependencies**
   ```bash
   go mod tidy
   ```
2. **Configure environment**
   ```bash
   cp .env.example .env
   # Edit .env with your settings
   ```
3. **Run locally**
   ```bash
   # Simple structure
   go run main.go
   # Clean architecture
   go run cmd/app/main.go
   ```
4. **Run with Docker**
   ```bash
   docker-compose up --build
   ```
5. **Test the server**
   ```bash
   curl http://localhost:8080/health
   # Should return: {"status": "healthy"}
   ```
## Tips
- Always use `--module` flag with your actual GitHub username
- Use `--clean-arch` for larger projects that need proper architecture
- Use `--docker` for easy deployment and development environments
- SQLite is great for development and small projects
- PostgreSQL/MySQL are better for production workloads
## Troubleshooting
**Issue: Cannot connect to database**
- Check your `.env` file has correct credentials
- Ensure database service is running
- For Docker: Run `docker-compose up` to start all services
**Issue: Module import errors**
- Run `go mod tidy` to download dependencies
- Verify your module path in `go.mod` matches your project
**Issue: Port already in use**
- Change `APP_PORT` in `.env` file
- Stop other services using port 8080
