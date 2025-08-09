package main

import (
	"facebook-pages-api-go/pkg/facebook"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get access token from environment variable (optional - can also be provided via request)
	accessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	
	// Create simple router (supports request-level authentication)
	router := facebook.NewSimpleRouter(accessToken)
	
	// Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Facebook Pages API Server (Simple Router) starting on port %s\n", port)
	fmt.Println("� Authentication options:")
	fmt.Println("  1. Query parameter: ?access_token=YOUR_TOKEN")
	fmt.Println("  2. Authorization header: Bearer YOUR_TOKEN")
	fmt.Println("  3. Environment variable: PAGE_ACCESS_TOKEN")
	fmt.Println()
	fmt.Println("�📚 Available endpoints:")
	fmt.Println("  GET /health                           - Health check")
	fmt.Println("  GET /api/pages/{pageId}               - Get page info")
	fmt.Println("  GET /api/pages                        - Get managed pages")
	fmt.Println("  GET /api/pages/{pageId}/posts         - Get page posts")
	fmt.Println("  GET /api/posts/{postId}/comments      - Get post comments")
	fmt.Println("  GET /api/comments/{commentId}         - Get specific comment")
	fmt.Println("  GET /api/comments/{commentId}/replies - Get comment replies")
	fmt.Println()
	fmt.Println("📖 Query parameters:")
	fmt.Println("  ?fields=field1,field2  - Select specific fields")
	fmt.Println("  ?limit=10             - Limit number of results")
	fmt.Println("  ?order=chronological  - Order comments (chronological|reverse_chronological)")
	fmt.Println()
	fmt.Printf("🌐 Server running at http://localhost:%s\n", port)
	fmt.Println("✨ Using standard library only (no external dependencies)")

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, router))
}
