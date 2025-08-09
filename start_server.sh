#!/bin/bash

# Facebook Pages API Server Startup Script

echo "🚀 Facebook Pages API Server"
echo "============================"

# Help function
show_help() {
    echo ""
    echo "Usage: $0 [simple|help] [options]"
    echo ""
    echo "Options:"
    echo "  simple    Use standard library router (no dependencies)"
    echo "  help      Show this help message"
    echo ""
    echo "Environment Variables:"
    echo "  PAGE_ACCESS_TOKEN  Facebook Page Access Token (required)"
    echo "  PAGE_ID           Facebook Page ID (optional, for testing)"
    echo "  PORT              Server port (default: 8080)"
    echo ""
    echo "Examples:"
    echo "  $0                Start server with Gorilla Mux"
    echo "  $0 simple         Start simple server (standard library)"
    echo "  PORT=3000 $0      Start server on port 3000"
    echo ""
    echo "Testing:"
    echo "  curl http://localhost:8080/health"
    echo "  go run cmd/client/main.go"
    echo ""
    exit 0
}

# Check for help flag
if [ "$1" = "help" ] || [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    show_help
fi

# Check if access token is provided
if [ -z "$PAGE_ACCESS_TOKEN" ]; then
    echo "❌ Error: PAGE_ACCESS_TOKEN environment variable is required"
    echo ""
    echo "💡 Get your token at: https://developers.facebook.com/"
    echo ""
    echo "Usage:"
    echo "  export PAGE_ACCESS_TOKEN='your_page_access_token'"
    echo "  $0"
    echo ""
    echo "Or:"
    echo "  PAGE_ACCESS_TOKEN='your_token' $0"
    echo ""
    echo "For help: $0 --help"
    exit 1
fi

# Set default port if not provided
PORT=${PORT:-8080}

echo "✅ Configuration:"
echo "   🔑 Access token: ${PAGE_ACCESS_TOKEN:0:20}..."
echo "   🌐 Port: $PORT"
if [ -n "$PAGE_ID" ]; then
    echo "   📄 Page ID: $PAGE_ID"
fi
echo ""

# Check if we should use simple server or regular server
if [ "$1" = "simple" ]; then
    echo "🎯 Starting Simple Server (standard library only)..."
    echo "📚 No external dependencies required"
    echo ""
    echo "🌐 Server will be available at: http://localhost:$PORT"
    echo "🔍 Health check: http://localhost:$PORT/health"
    echo ""
    echo "📋 Available endpoints:"
    echo "  GET /health                           - Health check"
    echo "  GET /api/pages/{pageId}               - Get page info"
    echo "  GET /api/pages                        - Get managed pages"
    echo "  GET /api/pages/{pageId}/posts         - Get page posts"
    echo "  GET /api/posts/{postId}/comments      - Get post comments"
    echo "  GET /api/comments/{commentId}         - Get specific comment"
    echo "  GET /api/comments/{commentId}/replies - Get comment replies"
    echo ""
    echo "🚀 Starting server..."
    go run cmd/simple-server/main.go
else
    echo "🎯 Starting Full Server (with Gorilla Mux)..."
    echo "📚 Using Gorilla Mux for advanced routing"
    echo ""
    echo "🌐 Server will be available at: http://localhost:$PORT"
    echo "🔍 Health check: http://localhost:$PORT/health"
    echo ""
    echo "📋 Available endpoints:"
    echo "  GET /health                           - Health check"
    echo "  GET /api/pages/{pageId}               - Get page info"
    echo "  GET /api/pages                        - Get managed pages"
    echo "  GET /api/pages/{pageId}/posts         - Get page posts"
    echo "  GET /api/posts/{postId}/comments      - Get post comments"
    echo "  GET /api/comments/{commentId}         - Get specific comment"
    echo "  GET /api/comments/{commentId}/replies - Get comment replies"
    echo ""
    echo "🚀 Starting server..."
    go run cmd/server/main.go
fi
