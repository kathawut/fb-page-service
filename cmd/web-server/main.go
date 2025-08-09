package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	// Set up file server for static files
	webDir := filepath.Join(currentDir, "web")
	fs := http.FileServer(http.Dir(webDir))
	
	// Handle all requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If requesting root, serve privacy policy
		if r.URL.Path == "/" {
			http.ServeFile(w, r, filepath.Join(webDir, "privacy-policy.html"))
			return
		}
		
		// For other paths, serve static files
		fs.ServeHTTP(w, r)
	})

	// Add specific route for privacy policy
	http.HandleFunc("/privacy-policy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(webDir, "privacy-policy.html"))
	})

	// Add specific route for terms of service
	http.HandleFunc("/terms-of-service", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(webDir, "terms-of-service.html"))
	})

	// Add a simple health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status":"ok","service":"privacy-policy-server"}`)
	})

	fmt.Printf("🌐 Privacy Policy Server starting on port %s\n", port)
	fmt.Printf("📄 Privacy Policy: http://localhost:%s/privacy-policy\n", port)
	fmt.Printf("📋 Terms of Service: http://localhost:%s/terms-of-service\n", port)
	fmt.Printf("🏠 Root URL: http://localhost:%s/\n", port)
	fmt.Println("✨ Static file server for web directory")

	// Start server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
