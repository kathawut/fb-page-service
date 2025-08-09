# Facebook Pages API - Quick Reference

## 🚀 Ways to Run the API

### 1. Quick Start (Recommended)
```bash
# Setup everything
./setup.sh

# Set credentials
export PAGE_ACCESS_TOKEN="your_token"
export PAGE_ID="your_page_id"

# Start simple server
./start_server.sh simple
```

### 2. Using Make Commands
```bash
# Complete setup
make setup

# Start development server
make run-simple

# Run tests
make test-api
```

### 3. Manual Commands
```bash
# Install dependencies
go mod tidy
go get github.com/gorilla/mux

# Build servers
go build -o bin/server cmd/server/main.go
go build -o bin/simple-server cmd/simple-server/main.go

# Run server
PAGE_ACCESS_TOKEN="your_token" go run cmd/server/main.go
```

### 4. Docker Deployment
```bash
# Build image
docker build -t facebook-pages-api .

# Run container
docker run -p 8080:8080 \
  -e PAGE_ACCESS_TOKEN="your_token" \
  facebook-pages-api

# Or use docker-compose
docker-compose up
```

## 🧪 Testing the API

### Health Check
```bash
curl http://localhost:8080/health
```

### Test All Endpoints
```bash
# Built-in test client
go run cmd/client/main.go

# Comprehensive test script
./test_api.sh

# Make command
make test-api
```

### Manual Testing
```bash
# Get page info
curl "http://localhost:8080/api/pages/YOUR_PAGE_ID"

# Get posts
curl "http://localhost:8080/api/pages/YOUR_PAGE_ID/posts?limit=5"

# Get comments
curl "http://localhost:8080/api/posts/POST_ID/comments"
```

## 📊 Available Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Server health check |
| GET | `/api/pages/{pageId}` | Get page information |
| GET | `/api/pages` | Get managed pages |
| GET | `/api/pages/{pageId}/posts` | Get page posts |
| GET | `/api/posts/{postId}/comments` | Get post comments |
| GET | `/api/comments/{commentId}` | Get comment details |
| GET | `/api/comments/{commentId}/replies` | Get comment replies |

## 🔧 Configuration

### Environment Variables
```bash
# Required
export PAGE_ACCESS_TOKEN="your_facebook_page_access_token"

# Optional
export PAGE_ID="your_page_id"        # For testing
export PORT="8080"                   # Server port
export API_VERSION="v23.0"          # Facebook API version
```

### Query Parameters
```bash
# Field selection
?fields=id,name,category,fan_count

# Limit results
?limit=10

# Comment ordering
?order=chronological  # or reverse_chronological
```

## 🎯 Server Options

### Simple Server (Recommended)
- Uses Go standard library only
- No external dependencies
- Lightweight and fast
- Perfect for development

```bash
./start_server.sh simple
```

### Full Server
- Uses Gorilla Mux router
- Advanced routing features
- Middleware support
- Better for production

```bash
./start_server.sh
```

## 📚 Tools & Utilities

### Postman Collection
```bash
# Generate custom environment
./postman/generate_environment.sh

# Import these files into Postman:
# - postman/Facebook_Pages_API.postman_collection.json
# - postman/Facebook_Pages_API.postman_environment.json
```

### Development Tools
```bash
# Auto-reload server (install air first)
make dev

# Check Facebook token
make check-token

# Generate Postman environment
make postman
```

## 🐛 Troubleshooting

### Server Won't Start
```bash
# Check if port is in use
lsof -i :8080

# Try different port
PORT=3000 ./start_server.sh
```

### Authentication Errors
```bash
# Check token is set
echo $PAGE_ACCESS_TOKEN

# Validate token
curl "https://graph.facebook.com/me?access_token=$PAGE_ACCESS_TOKEN"
```

### API Errors
- Ensure using Facebook API v23.0
- Check page permissions
- Verify page ID is correct
- Some endpoints need user tokens, not page tokens

## 📖 Getting Help

```bash
# Show server help
./start_server.sh --help

# Show make commands
make help

# Check documentation
ls docs/
```

## 🚀 Production Deployment

### Binary Deployment
```bash
# Build production binary
make prod-build

# Deploy
./bin/facebook-pages-api-prod
```

### Docker Deployment
```bash
# Build and run
docker build -t facebook-pages-api .
docker run -p 80:8080 -e PAGE_ACCESS_TOKEN="$TOKEN" facebook-pages-api
```

### Docker Compose
```bash
# Set environment variables in .env file
echo "PAGE_ACCESS_TOKEN=your_token" > .env

# Deploy
docker-compose up -d
```
