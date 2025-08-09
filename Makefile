# Facebook Pages API Go - Makefile

.PHONY: help setup build run run-simple test clean install dev docker

# Default target
help:
	@echo "🚀 Facebook Pages API Go - Available Commands"
	@echo "=============================================="
	@echo ""
	@echo "Setup & Installation:"
	@echo "  make setup        - Complete project setup"
	@echo "  make install      - Install dependencies"
	@echo "  make build        - Build all binaries"
	@echo ""
	@echo "Development:"
	@echo "  make dev          - Start development server (auto-reload)"
	@echo "  make run          - Start API server (Gorilla Mux)"
	@echo "  make run-simple   - Start simple server (standard library)"
	@echo "  make test         - Run tests"
	@echo "  make test-api     - Test API endpoints"
	@echo ""
	@echo "Deployment:"
	@echo "  make docker       - Build Docker image"
	@echo "  make clean        - Clean build artifacts"
	@echo ""
	@echo "Utilities:"
	@echo "  make postman      - Generate Postman environment"
	@echo "  make check-token  - Validate Facebook token"
	@echo ""
	@echo "Environment Variables:"
	@echo "  PAGE_ACCESS_TOKEN - Facebook Page Access Token (required)"
	@echo "  PAGE_ID          - Facebook Page ID (optional)"
	@echo "  PORT             - Server port (default: 8080)"

# Setup everything
setup:
	@echo "🔧 Running complete setup..."
	./setup.sh

# Install dependencies
install:
	@echo "📦 Installing dependencies..."
	go mod tidy
	go get github.com/gorilla/mux

# Build all components
build:
	@echo "🔨 Building all components..."
	go build -o bin/facebook-pages-api-go .
	go build -o bin/server cmd/server/main.go
	go build -o bin/simple-server cmd/simple-server/main.go
	go build -o bin/client cmd/client/main.go
	@echo "✅ Build complete! Binaries in bin/"

# Run full server (Gorilla Mux)
run:
	@echo "🚀 Starting API server (Gorilla Mux)..."
	@./start_server.sh

# Run simple server (standard library)
run-simple:
	@echo "🚀 Starting simple API server..."
	@./start_server.sh simple

# Development mode with auto-reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	@if command -v air > /dev/null; then \
		echo "🔄 Starting development server with auto-reload..."; \
		air -c .air.toml; \
	else \
		echo "📦 Installing air for auto-reload..."; \
		go install github.com/cosmtrek/air@latest; \
		echo "🔄 Starting development server with auto-reload..."; \
		air -c .air.toml; \
	fi

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./... -v

# Test API endpoints
test-api:
	@echo "🧪 Testing API endpoints..."
	@if [ -z "$$PAGE_ACCESS_TOKEN" ]; then \
		echo "❌ PAGE_ACCESS_TOKEN is required for API testing"; \
		exit 1; \
	fi
	go run cmd/client/main.go

# Generate Postman environment
postman:
	@echo "📊 Generating Postman environment..."
	./postman/generate_environment.sh

# Check Facebook token
check-token:
	@echo "🔍 Checking Facebook token..."
	@if [ -z "$$PAGE_ACCESS_TOKEN" ]; then \
		echo "❌ PAGE_ACCESS_TOKEN environment variable is required"; \
		exit 1; \
	fi
	@curl -s "https://graph.facebook.com/me?access_token=$$PAGE_ACCESS_TOKEN" | \
		jq -r 'if .error then "❌ Token invalid: " + .error.message else "✅ Token valid for: " + .name end'

# Docker build
docker:
	@echo "🐳 Building Docker image..."
	docker build -t facebook-pages-api:latest .
	@echo "✅ Docker image built: facebook-pages-api:latest"
	@echo "🚀 To run: docker run -p 8080:8080 -e PAGE_ACCESS_TOKEN='your_token' facebook-pages-api:latest"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf bin/
	rm -f facebook-pages-api-go
	rm -f cmd/server/server
	rm -f cmd/simple-server/simple-server
	rm -f cmd/client/client
	go clean
	@echo "✅ Clean complete"

# Development dependencies
dev-deps:
	@echo "📦 Installing development dependencies..."
	go install github.com/cosmtrek/air@latest
	@echo "✅ Development dependencies installed"

# Quick start (setup + run)
quick-start: setup run-simple

# Production build
prod-build:
	@echo "🏭 Building for production..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/facebook-pages-api-prod cmd/server/main.go
	@echo "✅ Production build complete: bin/facebook-pages-api-prod"
