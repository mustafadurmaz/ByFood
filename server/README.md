📦 Backend Setup

# Navigate to the backend directory
cd backend

# Start PostgreSQL with Docker
docker-compose up -d

# Install dependencies
go mod tidy

# Run the server
go run main.go
