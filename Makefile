.PHONY: start start_prod test build clean

# Set the environment variable and run the application in local development mode
start:
	ENV=local go run main.go

# Run the application in production mode
start_prod:
	go run main.go

# Run tests
test:
	go test ./...

# Build the application
build:
	go build -o bin/web-service-gin main.go

# Clean the build artifacts
clean:
	rm -rf bin