alias r:=run

# List all available tasks by default
@default:
    just --list

# 🔧 Prepare environment with Postgres
@prepare:
    docker run --name postgres16 -p 5434:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres:16

# 🛠 Create Database inside running container
@createdb:
	docker exec -it postgres16 createdb --username=postgres --owner=postgres school_ride_backend

# ▶️ Start PostgreSQL container
@startdb:
    @echo "Starting the Backend Database Server..."
    docker start postgres16

# ⏹ Stop PostgreSQL container
@stopdb:
    @echo "Stopping the Backend Database Server..."
    docker stop postgres16

# 🧹 Remove the container
@rmdb:
    docker rm -f postgres16

# 🔍 View logs
@dblogs:
    docker logs -f postgres16

# 🛠 Restart the database
@restartdb:
    just stopdb
    just startdb

# 📦 Install dependencies
@install:
    @echo "Installing dependencies..."
    go mod tidy
    go mod download
    go install github.com/air-verse/air@latest
    go install github.com/swaggo/swag/cmd/swag@latest

# 📖 Generate Swagger documentation
@swag:
    @echo "Generating Swagger documentation..."
    swag init -g ./cmd/server/main.go -o cmd/docs
    swag fmt
    @echo "Swagger documentation generated successfully."

# 🏃‍♂️ Run the backend server
@run:
    just startdb
    just swag
    @echo "Running the Backend Server..."
    air