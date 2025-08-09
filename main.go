package main

import (
	"facebook-pages-api-go/pkg/facebook"
	"fmt"
	"log"
	"os"
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
		fmt.Println("1. Command line: go run main.go <PAGE_ACCESS_TOKEN> <PAGE_ID> [API_VERSION]")
		fmt.Println("2. Environment variables:")
		fmt.Println("   export PAGE_ACCESS_TOKEN='your_token_here'")
		fmt.Println("   export PAGE_ID='your_page_id_here'")
		fmt.Println("   export API_VERSION='v23.0'  # Optional")
		fmt.Println("   go run main.go")
		return
	}

	if pageID == "" {
		fmt.Println("⚠️  PAGE_ID not provided.")
		fmt.Println("Please provide your Facebook Page ID.")
		fmt.Println("\n📋 Usage: go run main.go <PAGE_ACCESS_TOKEN> <PAGE_ID> [API_VERSION]")
		return
	}

	// Create Facebook client
	client := facebook.NewClient(pageAccessToken)
	
	if apiVersion != "" {
		client.SetAPIVersion(apiVersion)
		fmt.Printf("🔧 Using API version: %s\n", apiVersion)
	} else {
		fmt.Printf("🔧 Using default API version: v23.0\n")
	}

	fmt.Println("\n🚀 Facebook Pages API Go Client Demo")
	fmt.Println("=====================================")
	fmt.Printf("📄 Page ID: %s\n", pageID)
	fmt.Printf("🔑 Token: %s...%s\n", pageAccessToken[:10], pageAccessToken[len(pageAccessToken)-10:])

	// Example 1: Get page information
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
	}

	// Example 2: Get recent posts
	fmt.Println("\n� Getting recent posts...")
	postsResp, err := client.GetPosts(pageID, 3) // Get last 3 posts
	if err != nil {
		log.Printf("Error getting posts: %v", err)
	} else {
		fmt.Printf("Found %d recent posts:\n", len(postsResp.Data))
		for i, p := range postsResp.Data {
			fmt.Printf("%d. Post ID: %s\n", i+1, p.ID)
			if p.Message != "" {
				message := p.Message
				if len(message) > 100 {
					message = message[:100] + "..."
				}
				fmt.Printf("   Message: %s\n", message)
			}
			fmt.Printf("   Created: %s\n", p.CreatedTime.Format("2006-01-02 15:04:05"))
			
			// Get comments for this post
			fmt.Printf("   Getting comments...\n")
			comments, err := client.GetPostComments(p.ID, 3, "chronological")
			if err != nil {
				log.Printf("   Error getting comments: %v", err)
			} else {
				if comments.Summary.TotalCount > 0 {
					fmt.Printf("   Total comments: %d\n", comments.Summary.TotalCount)
					for j, comment := range comments.Data {
						fmt.Printf("   Comment %d: %s - %s\n", j+1, comment.From.Name, 
							func() string {
								if len(comment.Message) > 50 {
									return comment.Message[:50] + "..."
								}
								return comment.Message
							}())
					}
				} else {
					fmt.Printf("   No comments found\n")
				}
			}
			fmt.Println()
		}
	}

	fmt.Println("\n🎉 Demo completed successfully!")
	fmt.Println("\n📚 Available examples:")
	fmt.Println("- go run examples/basic/main.go      # Comprehensive demo")
	fmt.Println("- go run examples/comments/main.go   # Comments demo")
	fmt.Println("- go run examples/insights/main.go   # Analytics demo")
	fmt.Println("- go run examples/photos/main.go     # Photo demo")
}
