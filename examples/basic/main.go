package main

import (
	"facebook-pages-api-go/pkg/facebook"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var pageAccessToken, pageID, apiVersion string

	// Check for command line arguments first
	if len(os.Args) >= 3 {
		pageAccessToken = os.Args[1]
		pageID = os.Args[2]
		if len(os.Args) >= 4 {
			apiVersion = os.Args[3]
		}
		fmt.Println("📝 Using parameters from command line arguments")
	} else {
		// Fall back to environment variables
		pageAccessToken = os.Getenv("PAGE_ACCESS_TOKEN")
		pageID = os.Getenv("PAGE_ID")
		apiVersion = os.Getenv("API_VERSION")
		fmt.Println("📝 Using parameters from environment variables")
	}

	if pageAccessToken == "" {
		fmt.Println("⚠️  PAGE_ACCESS_TOKEN not provided.")
		fmt.Println("Please provide your Facebook Page Access Token.")
		fmt.Println("\n📋 Usage options:")
		fmt.Println("1. Command line: go run examples/basic/main.go <PAGE_ACCESS_TOKEN> <PAGE_ID> [API_VERSION]")
		fmt.Println("2. Environment variables:")
		fmt.Println("   export PAGE_ACCESS_TOKEN='your_token_here'")
		fmt.Println("   export PAGE_ID='your_page_id_here'")
		fmt.Println("   export API_VERSION='v23.0'  # Optional")
		fmt.Println("   go run examples/basic/main.go")
		return
	}

	if pageID == "" {
		fmt.Println("⚠️  PAGE_ID not provided.")
		fmt.Println("Please provide your Facebook Page ID.")
		fmt.Println("\n📋 Usage: go run examples/basic/main.go <PAGE_ACCESS_TOKEN> <PAGE_ID> [API_VERSION]")
		return
	}

	// Create Facebook client
	client := facebook.NewClient(pageAccessToken)
	
	if apiVersion != "" {
		client.SetAPIVersion(apiVersion)
	}

	fmt.Println("🚀 Facebook Pages API Go Client Demo")
	fmt.Println("=====================================")

	// Example 1: Validate access token
	fmt.Println("\n🔐 Validating access token...")
	err := client.ValidateAccessToken()
	if err != nil {
		log.Printf("❌ Invalid access token: %v", err)
		return
	}
	fmt.Println("✅ Access token is valid!")

	// Example 2: Get user info
	fmt.Println("\n👤 Getting user information...")
	user, err := client.GetUserInfo()
	if err != nil {
		log.Printf("Error getting user info: %v", err)
	} else {
		fmt.Printf("User ID: %s\n", user.ID)
		fmt.Printf("User Name: %s\n", user.Name)
		if user.Email != "" {
			fmt.Printf("Email: %s\n", user.Email)
		}
	}

	// Example 3: Get page information
	fmt.Println("\n📄 Getting page information...")
	page, err := client.GetPage(pageID)
	if err != nil {
		log.Printf("Error getting page: %v", err)
	} else {
		fmt.Printf("Page Name: %s\n", page.Name)
		fmt.Printf("Page ID: %s\n", page.ID)
		fmt.Printf("Category: %s\n", page.Category)
		fmt.Printf("Fan Count: %d\n", page.FanCount)
		fmt.Printf("Is Verified: %t\n", page.IsVerified)
		if page.Website != "" {
			fmt.Printf("Website: %s\n", page.Website)
		}
		if page.About != "" {
			fmt.Printf("About: %s\n", page.About)
		}
	}

	// Example 4: Get managed pages
	fmt.Println("\n📋 Getting managed pages...")
	pages, err := client.GetPages()
	if err != nil {
		log.Printf("Error getting pages: %v", err)
	} else {
		fmt.Printf("Found %d managed pages:\n", len(pages))
		for i, p := range pages {
			fmt.Printf("%d. %s (ID: %s) - Can Post: %t\n", i+1, p.Name, p.ID, p.CanPost)
		}
	}

	// Example 5: Get recent posts
	fmt.Println("\n📰 Getting recent posts...")
	postsResp, err := client.GetPosts(pageID, 5) // Get last 5 posts
	if err != nil {
		log.Printf("Error getting posts: %v", err)
	} else {
		fmt.Printf("Found %d recent posts:\n", len(postsResp.Data))
		for i, p := range postsResp.Data {
			fmt.Printf("%d. Post ID: %s\n", i+1, p.ID)
			if p.Message != "" {
				// Truncate long messages
				message := p.Message
				if len(message) > 100 {
					message = message[:100] + "..."
				}
				fmt.Printf("   Message: %s\n", message)
			}
			fmt.Printf("   Created: %s\n", p.CreatedTime.Format("2006-01-02 15:04:05"))
			fmt.Printf("   Type: %s\n", p.Type)
			if p.Permalink != "" {
				fmt.Printf("   Link: %s\n", p.Permalink)
			}
			fmt.Println()
		}
	}

	// Example 6: Get page insights (last 7 days)
	fmt.Println("\n📊 Getting page insights...")
	since := time.Now().AddDate(0, 0, -7) // 7 days ago
	until := time.Now()
	
	insights, err := client.GetPageInsights(pageID, nil, "day", &since, &until)
	if err != nil {
		log.Printf("Error getting insights: %v", err)
	} else {
		fmt.Printf("Found %d insight metrics:\n", len(insights.Data))
		for i, insight := range insights.Data {
			if i >= 5 { // Show only first 5 metrics
				break
			}
			fmt.Printf("📈 %s (%s)\n", insight.Name, insight.Period)
			if len(insight.Values) > 0 {
				latestValue := insight.Values[len(insight.Values)-1]
				fmt.Printf("   Latest Value: %v (Date: %s)\n", 
					latestValue.Value, 
					latestValue.EndTime.Format("2006-01-02"))
			}
		}
		if len(insights.Data) > 5 {
			fmt.Printf("... and %d more metrics\n", len(insights.Data)-5)
		}
	}

	// Example 7: Upload photo by URL (optional)
	fmt.Println("\n🖼️  Testing photo upload by URL...")
	photoURL := "https://via.placeholder.com/400x300/0066cc/ffffff?text=Facebook+API+Test"
	photoResp, err := client.UploadPhotoByURL(pageID, photoURL, "Test photo uploaded via Facebook Pages API Go Client! 📸\n\nThis demonstrates photo upload functionality.", true)
	if err != nil {
		log.Printf("Error uploading photo: %v", err)
	} else {
		fmt.Printf("✅ Photo uploaded successfully! Photo ID: %s\n", photoResp.ID)
		if photoResp.PostID != "" {
			fmt.Printf("   Post ID: %s\n", photoResp.PostID)
		}
	}

	fmt.Println("\n🎉 Demo completed successfully!")
	fmt.Println("\n📚 Next steps:")
	fmt.Println("1. Check your Facebook page to see the new posts and photo")
	fmt.Println("2. Explore the Facebook Pages API documentation: https://developers.facebook.com/docs/pages-api")
	fmt.Println("3. Customize the client for your specific use case")
	fmt.Println("4. Add error handling and logging for production use")
	fmt.Println("5. Implement pagination for large datasets")
	
	fmt.Println("\n🔧 Available API Methods:")
	fmt.Println("- GetPage() - Get page information")
	fmt.Println("- GetPages() - Get managed pages")
	fmt.Println("- CreatePost() - Create new posts")
	fmt.Println("- GetPosts() - Retrieve posts")
	fmt.Println("- DeletePost() - Delete posts")
	fmt.Println("- UploadPhoto() - Upload photos from file")
	fmt.Println("- UploadPhotoByURL() - Upload photos from URL")
	fmt.Println("- GetPageInsights() - Get page analytics")
	fmt.Println("- GetPostInsights() - Get post analytics")
	fmt.Println("- ValidateAccessToken() - Validate tokens")
}
