package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	baseURL := "http://localhost:8080"
	
	// Override base URL if provided
	if len(os.Args) > 1 {
		baseURL = os.Args[1]
	}
	
	fmt.Println("🧪 Facebook Pages API Client Test")
	fmt.Printf("🌐 Testing server at: %s\n", baseURL)
	fmt.Println("=" + fmt.Sprintf("%s", "================================="))
	
	client := &http.Client{Timeout: 10 * time.Second}
	
	// Test health check
	fmt.Println("\n🔍 Testing health check...")
	testEndpoint(client, baseURL+"/health")
	
	// Test page info (you'll need to replace with a valid page ID)
	pageID := "471958375999732" // Demo page ID
	fmt.Printf("\n📄 Testing page info for page %s...\n", pageID)
	testEndpoint(client, fmt.Sprintf("%s/api/pages/%s?fields=id,name,category,fan_count", baseURL, pageID))
	
	// Test posts
	fmt.Printf("\n📰 Testing posts for page %s...\n", pageID)
	testEndpoint(client, fmt.Sprintf("%s/api/pages/%s/posts?limit=2&fields=id,message,created_time", baseURL, pageID))
	
	// Test managed pages (will likely fail with page token)
	fmt.Println("\n📋 Testing managed pages...")
	testEndpoint(client, baseURL+"/api/pages")
	
	fmt.Println("\n✅ API testing completed!")
	fmt.Println("\n📝 Note: Some endpoints may fail if:")
	fmt.Println("  - Server is not running")
	fmt.Println("  - Invalid access token")
	fmt.Println("  - Page ID doesn't exist")
	fmt.Println("  - Using page token instead of user token for managed pages")
}

func testEndpoint(client *http.Client, url string) {
	fmt.Printf("  GET %s\n", url)
	
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("  ❌ Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("  ❌ Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("  📊 Status: %d\n", resp.StatusCode)
	
	// Pretty print JSON response
	var jsonData interface{}
	if err := json.Unmarshal(body, &jsonData); err == nil {
		prettyJSON, _ := json.MarshalIndent(jsonData, "    ", "  ")
		fmt.Printf("  📄 Response:\n    %s\n", string(prettyJSON))
	} else {
		fmt.Printf("  📄 Response: %s\n", string(body))
	}
}
