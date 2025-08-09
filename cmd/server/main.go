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
	
	// Create router (supports request-level authentication)
	router := facebook.NewRouter(accessToken)
	
	// Setup routes
	r := router.SetupRoutes()
	
	// Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Facebook Pages API Server starting on port %s\n", port)
	fmt.Println("📚 Available endpoints:")
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

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, r))
}
