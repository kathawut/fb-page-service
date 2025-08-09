# Facebook Pages API Go - Complete Solution

A comprehensive Go implementation providing both a client library and REST API server for Facebook Pages API v23.0.

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or later
- Facebook Page Access Token
- Facebook Page ID

### 1. Setup
```bash
git clone <your-repo>
cd facebook-pages-api-go
chmod +x setup.sh start_server.sh
```

### 2. Configure
```bash
# Set your Facebook credentials
export PAGE_ACCESS_TOKEN="your_facebook_page_access_token"
export PAGE_ID="your_facebook_page_id"

# Optional: Set custom port
export PORT="8080"
```

### 3. Start API Server
```bash
# Option 1: Simple server (no external dependencies)
./start_server.sh simple

# Option 2: Full-featured server (with Gorilla Mux)
./start_server.sh

# Option 3: Manual start
go run cmd/server/main.go
```

### 4. Test the API
```bash
# Test with built-in client
go run cmd/client/main.go

# Test with curl
curl http://localhost:8080/health
curl "http://localhost:8080/api/pages/$PAGE_ID"
```

## 🌟 What's Included

### 📡 REST API Server
- **Health Check**: `GET /health`
- **Page Info**: `GET /api/pages/{pageId}`
- **Managed Pages**: `GET /api/pages`
- **Page Posts**: `GET /api/pages/{pageId}/posts`
- **Post Comments**: `GET /api/posts/{postId}/comments`
- **Comment Details**: `GET /api/comments/{commentId}`
- **Comment Replies**: `GET /api/comments/{commentId}/replies`

### 📚 Go Client Library
- Direct Go integration
- All Facebook Pages API v23.0 operations
- Custom field selection
- Error handling and validation

### 🧪 Testing Tools
- Built-in HTTP client for testing
- Complete Postman collection
- Environment generators
- Example usage code

### 📖 Implementation Options
- **Standard Library**: Zero external dependencies
- **Gorilla Mux**: Advanced routing with middleware support
- **Docker Ready**: Container deployment support

## 🔧 API Endpoints

| Method | Endpoint | Description | Parameters |
|--------|----------|-------------|------------|
| `GET` | `/health` | Server health check | None |
| `GET` | `/api/pages/{pageId}` | Get page information | `fields` |
| `GET` | `/api/pages` | Get managed pages | None |
| `GET` | `/api/pages/{pageId}/posts` | Get page posts | `limit`, `fields` |
| `GET` | `/api/posts/{postId}/comments` | Get post comments | `limit`, `order`, `fields` |
| `GET` | `/api/comments/{commentId}` | Get comment details | `fields` |
| `GET` | `/api/comments/{commentId}/replies` | Get comment replies | `limit`, `fields` |

## 📖 Usage Examples

### Get Page Information
```bash
curl "http://localhost:8080/api/pages/YOUR_PAGE_ID?fields=id,name,category,fan_count"
```

### Get Recent Posts
```bash
curl "http://localhost:8080/api/pages/YOUR_PAGE_ID/posts?limit=5&fields=id,message,created_time"
```

### Get Post Comments
```bash
curl "http://localhost:8080/api/posts/POST_ID/comments?order=chronological&limit=10"
```

## 🎯 Advanced Features

### Field Selection
Optimize API calls by requesting only needed fields:
```bash
# Get only essential page info
curl "http://localhost:8080/api/pages/123?fields=id,name,fan_count"

# Get detailed post information
curl "http://localhost:8080/api/pages/123/posts?fields=id,message,created_time,type,permalink_url"
```

### Pagination
Handle large datasets with pagination:
```bash
# Limit results
curl "http://localhost:8080/api/pages/123/posts?limit=20"

# Get comments in specific order
curl "http://localhost:8080/api/posts/456/comments?order=reverse_chronological&limit=50"
```

## 🧪 Testing & Development

### Built-in Test Client
```bash
# Test all endpoints
go run cmd/client/main.go

# Test specific server
go run cmd/client/main.go http://localhost:3000
```

### Postman Collection
```bash
# Generate custom environment for your page
./postman/generate_environment.sh

# Import files into Postman:
# - postman/Facebook_Pages_API.postman_collection.json
# - postman/Facebook_Pages_API.postman_environment.json
```

### Direct Go Client Usage
```go
package main

import (
    "facebook-pages-api-go/pkg/facebook"
    "fmt"
    "log"
)

func main() {
    client := facebook.NewClient("your_access_token")
    
    // Get page info
    page, err := client.GetPage("your_page_id")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Page: %s (ID: %s)\n", page.Name, page.ID)
    fmt.Printf("Fans: %d\n", page.FanCount)
}
```

## 🚀 Deployment

### Local Development
```bash
# Start development server
./start_server.sh simple
```

### Docker Deployment
```bash
# Build image
docker build -t facebook-pages-api .

# Run container
docker run -p 8080:8080 \
  -e PAGE_ACCESS_TOKEN="your_token" \
  facebook-pages-api
```

### Production Setup
```bash
# Set production environment
export PORT="80"
export PAGE_ACCESS_TOKEN="production_token"

# Build production binary
go build -o facebook-api cmd/server/main.go

# Run production server
./facebook-api
```

## 📁 Project Structure

```
facebook-pages-api-go/
├── cmd/
│   ├── server/          # Main API server
│   ├── simple-server/   # Standard library server
│   └── client/          # Test client
├── pkg/facebook/        # Core library
│   ├── client.go        # HTTP client
│   ├── pages.go         # Pages operations
│   ├── types.go         # Data structures
│   ├── router.go        # Gorilla Mux router
│   └── simple_router.go # Standard library router
├── examples/            # Usage examples
├── postman/            # Postman collection
├── docs/               # Documentation
└── scripts/            # Setup scripts
```

## 🔧 Configuration

### Environment Variables
```bash
# Required
PAGE_ACCESS_TOKEN="your_facebook_page_access_token"

# Optional
PAGE_ID="your_default_page_id"           # For testing
PORT="8080"                              # Server port
API_VERSION="v23.0"                      # Facebook API version
```

### Access Token Setup
1. Go to [Facebook Developers](https://developers.facebook.com/)
2. Create an app and get Page Access Token
3. Ensure token has necessary permissions:
   - `pages_show_list`
   - `pages_read_engagement`
   - `pages_read_user_content`

## 🔍 Troubleshooting

### Common Issues

**Server won't start:**
```bash
# Check if port is available
lsof -i :8080

# Try different port
PORT=3000 ./start_server.sh
```

**Authentication errors:**
```bash
# Verify token is set
echo $PAGE_ACCESS_TOKEN

# Test token validity
curl "https://graph.facebook.com/me?access_token=$PAGE_ACCESS_TOKEN"
```

**API errors:**
- Check Facebook API version (v23.0)
- Verify page permissions
- Ensure page ID is correct and accessible

### Debug Mode
Enable detailed logging:
```bash
# Add debug flag
go run cmd/server/main.go -debug

# Check server logs
curl http://localhost:8080/health -v
```

## 📚 Documentation

- **[API Documentation](docs/api.md)** - Complete API reference
- **[Client Library](docs/client.md)** - Go client usage guide
- **[Examples](examples/)** - Code examples and tutorials
- **[Postman Guide](postman/README.md)** - Testing with Postman

## 🤝 Contributing

1. Fork the repository
2. Create feature branch
3. Add tests for new functionality
4. Submit pull request

## 📄 License

MIT License - see LICENSE file for details.

## 🆘 Support

- **Issues**: Report bugs and feature requests
- **Documentation**: Check docs/ directory
- **Examples**: See examples/ directory
- **Facebook API**: [Official Facebook Graph API documentation](https://developers.facebook.com/docs/graph-api/)
