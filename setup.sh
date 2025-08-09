#!/bin/bash

# Facebook Pages API Go - Complete Setup Script

echo "🚀 Facebook Pages API Go - Complete Setup"
echo "=========================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first: https://golang.org/dl/"
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Install dependencies
echo "📦 Installing dependencies..."
go mod tidy
go get github.com/gorilla/mux

# Build all components
echo "🔨 Building all components..."

# Build main client
echo "  � Building main client..."
go build -o facebook-pages-api-go .

# Build API servers
echo "  🌐 Building API servers..."
go build -o cmd/server/server cmd/server/main.go
go build -o cmd/simple-server/simple-server cmd/simple-server/main.go

# Build test client
echo "  🧪 Building test client..."
go build -o cmd/client/client cmd/client/main.go

if [ $? -eq 0 ]; then
    echo "✅ All builds successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

# Make scripts executable
echo "🔧 Setting up scripts..."
chmod +x start_server.sh
chmod +x postman/generate_environment.sh

echo "📋 Running tests..."
go test ./...

if [ $? -eq 0 ]; then
    echo "✅ All tests passed!"
else
    echo "⚠️  Some tests failed, but continuing setup..."
fi

# Create .env.example if it doesn't exist
if [ ! -f .env.example ]; then
    echo "📝 Creating .env.example..."
    cat > .env.example << 'EOF'
# Facebook Pages API Configuration
PAGE_ACCESS_TOKEN=your_facebook_page_access_token_here
PAGE_ID=your_facebook_page_id_here
PORT=8080
API_VERSION=v23.0
EOF
fi

echo ""
echo "🎉 Setup completed successfully!"
echo ""
echo "📋 Quick Start Guide:"
echo "===================="
echo ""
echo "1. 🔑 Get Facebook Credentials:"
echo "   • Go to https://developers.facebook.com/"
echo "   • Create a Facebook App"
echo "   • Get Page Access Token with required permissions"
echo ""
echo "2. ⚙️  Configure Environment:"
echo "   export PAGE_ACCESS_TOKEN='your_token_here'"
echo "   export PAGE_ID='your_page_id_here'"
echo ""
echo "3. 🚀 Start API Server:"
echo "   ./start_server.sh simple    # Standard library (recommended)"
echo "   ./start_server.sh           # With Gorilla Mux"
echo ""
echo "4. 🧪 Test the API:"
echo "   go run cmd/client/main.go   # Built-in test client"
echo "   curl http://localhost:8080/health     # Health check"
echo ""
echo "5. 📊 Use Postman Collection:"
echo "   ./postman/generate_environment.sh     # Generate custom environment"
echo "   # Then import postman/*.json files into Postman"
echo ""
echo "🔧 Available Commands:"
echo "  ./start_server.sh [simple]  - Start API server"
echo "  go run cmd/client/main.go   - Test client"
echo "  go run main.go              - Direct Facebook client"
echo ""
echo "📚 Documentation:"
echo "  README.md                   - Main documentation"
echo "  cmd/server/README.md        - API server guide"
echo "  postman/README.md           - Postman testing guide"
echo ""
echo "🆘 Need help? Check the documentation or run:"
echo "   ./start_server.sh --help"
echo "4. Or try specific examples:"
echo "   go run examples/basic/main.go"
echo "   go run examples/photos/main.go" 
echo "   go run examples/insights/main.go"
echo ""
echo "📚 Documentation:"
echo "   - README.md - General overview"
echo "   - docs/API.md - Complete API documentation"
echo "   - Facebook API Docs: https://developers.facebook.com/docs/pages-api"
echo ""
echo "🔧 Available features:"
echo "   ✅ Page information retrieval"
echo "   ✅ Post creation and management"
echo "   ✅ Photo uploads (file & URL)"
echo "   ✅ Page insights and analytics"
echo "   ✅ Access token validation"
echo "   ✅ Error handling with detailed messages"
echo "   ✅ Configurable API versions"
echo ""
