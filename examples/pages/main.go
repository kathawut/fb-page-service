package main

import (
	"facebook-pages-api-go/pkg/facebook"
	"fmt"
	"log"
	"os"
)

func main() {
	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	pageID := os.Getenv("PAGE_ID")

	// Accept command line arguments as fallback
	if len(os.Args) >= 3 {
		pageID = os.Args[1]
		pageAccessToken = os.Args[2]
	}

	if pageAccessToken == "" || pageID == "" {
		fmt.Println("Usage:")
		fmt.Println("  Set environment variables: PAGE_ID and PAGE_ACCESS_TOKEN")
		fmt.Println("  Or pass as arguments: go run main.go <PAGE_ID> <PAGE_ACCESS_TOKEN>")
		os.Exit(1)
	}

	client := facebook.NewClient(pageAccessToken)
	client.SetAPIVersion("v23.0") // Use latest API version

	fmt.Println("📄 Facebook Pages API Example")
	fmt.Println("=============================")

	// 1. Get page information
	fmt.Println("\n🏢 Getting page information...")
	page, err := client.GetPage(pageID)
	if err != nil {
		log.Printf("Error getting page info: %v", err)
		return
	}

	fmt.Printf("Page Name: %s\n", page.Name)
	fmt.Printf("Page ID: %s\n", page.ID)
	fmt.Printf("Category: %s\n", page.Category)
	if page.About != "" {
		fmt.Printf("About: %s\n", page.About)
	}
	if page.Website != "" {
		fmt.Printf("Website: %s\n", page.Website)
	}
	if page.Phone != "" {
		fmt.Printf("Phone: %s\n", page.Phone)
	}
	fmt.Printf("Fans: %d\n", page.FanCount)
	fmt.Printf("Followers: %d\n", page.FollowersCount)
	fmt.Printf("Talking About: %d\n", page.TalkingAboutCount)
	fmt.Printf("Verified: %t\n", page.IsVerified)
	fmt.Printf("Published: %t\n", page.IsPublished)
	fmt.Printf("Can Post: %t\n", page.CanPost)

	// 2. Get managed pages (requires user token, not page token)
	fmt.Println("\n📋 Getting managed pages...")
	pages, err := client.GetPages()
	if err != nil {
		fmt.Printf("Note: Cannot get managed pages with page access token (requires user token): %v\n", err)
		fmt.Println("To get managed pages, you need a user access token instead of a page access token.")
	} else {
		fmt.Printf("Found %d managed pages:\n", len(pages))
		for i, p := range pages {
			fmt.Printf("  %d. %s (ID: %s, Category: %s)\n", i+1, p.Name, p.ID, p.Category)
		}
	}

	// 3. Get recent posts
	fmt.Println("\n📰 Getting recent posts...")
	posts, err := client.GetPosts(pageID, 5)
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return
	}

	fmt.Printf("Found %d posts:\n", len(posts.Data))
	for i, post := range posts.Data {
		fmt.Printf("\n--- Post %d ---\n", i+1)
		fmt.Printf("ID: %s\n", post.ID)
		fmt.Printf("Type: %s\n", post.Type)
		fmt.Printf("Created: %s\n", post.CreatedTime.Time.Format("2006-01-02 15:04:05"))
		
		if post.Message != "" {
			message := post.Message
			if len(message) > 150 {
				message = message[:150] + "..."
			}
			fmt.Printf("Message: %s\n", message)
		}
		
		if post.Story != "" {
			fmt.Printf("Story: %s\n", post.Story)
		}
		
		if post.Link != "" {
			fmt.Printf("Link: %s\n", post.Link)
		}
		
		if post.Permalink != "" {
			fmt.Printf("Permalink: %s\n", post.Permalink)
		}
		
		fmt.Printf("Published: %t\n", post.IsPublished)
		fmt.Printf("Hidden: %t\n", post.IsHidden)

		// Get comments for this post
		fmt.Printf("\n💬 Getting comments...\n")
		comments, err := client.GetPostComments(post.ID, 3, "reverse_chronological")
		if err != nil {
			log.Printf("Error getting comments: %v", err)
		} else {
			if comments.Summary.TotalCount > 0 {
				fmt.Printf("Total comments: %d (showing %d)\n", comments.Summary.TotalCount, len(comments.Data))
				for j, comment := range comments.Data {
					fmt.Printf("  Comment %d: %s - %s\n", j+1, comment.From.Name, comment.Message)
					
					// Get replies if available
					if comment.CommentCount > 0 {
						replies, err := client.GetCommentReplies(comment.ID, 2)
						if err != nil {
							log.Printf("    Error getting replies: %v", err)
						} else {
							fmt.Printf("    %d replies:\n", len(replies.Data))
							for k, reply := range replies.Data {
								fmt.Printf("      Reply %d: %s - %s\n", k+1, reply.From.Name, reply.Message)
							}
						}
					}
				}
			} else {
				fmt.Printf("No comments on this post\n")
			}
		}

		// Only show first 2 posts for demo
		if i >= 1 {
			fmt.Printf("\n... (showing only first 2 posts for demo)\n")
			break
		}
	}

	// 4. Demonstrate getting specific page fields
	fmt.Println("\n🔍 Getting specific page fields...")
	pageFields, err := client.GetPage(pageID, "id", "name", "fan_count", "followers_count", "picture{url}")
	if err != nil {
		log.Printf("Error getting page fields: %v", err)
	} else {
		fmt.Printf("Selected fields:\n")
		fmt.Printf("  Name: %s\n", pageFields.Name)
		fmt.Printf("  Fans: %d\n", pageFields.FanCount)
		fmt.Printf("  Followers: %d\n", pageFields.FollowersCount)
		if pageFields.Picture.Data.URL != "" {
			fmt.Printf("  Picture URL: %s\n", pageFields.Picture.Data.URL)
		}
	}

	// 5. Demonstrate getting specific post fields
	if len(posts.Data) > 0 {
		firstPostID := posts.Data[0].ID
		fmt.Println("\n📝 Getting specific post fields...")
		specificPosts, err := client.GetPosts(pageID, 1, "id", "message", "created_time")
		if err != nil {
			log.Printf("Error getting specific post fields: %v", err)
		} else if len(specificPosts.Data) > 0 {
			post := specificPosts.Data[0]
			fmt.Printf("Post with selected fields:\n")
			fmt.Printf("  ID: %s\n", post.ID)
			fmt.Printf("  Created: %s\n", post.CreatedTime.Time.Format("2006-01-02 15:04:05"))
			if post.Message != "" {
				fmt.Printf("  Message: %s\n", post.Message)
			}
		}

		// 6. Get a specific comment
		fmt.Println("\n💭 Getting a specific comment...")
		comments, err := client.GetPostComments(firstPostID, 1, "reverse_chronological")
		if err != nil {
			log.Printf("Error getting comments: %v", err)
		} else if len(comments.Data) > 0 {
			commentID := comments.Data[0].ID
			comment, err := client.GetComment(commentID)
			if err != nil {
				log.Printf("Error getting specific comment: %v", err)
			} else {
				fmt.Printf("Specific comment details:\n")
				fmt.Printf("  ID: %s\n", comment.ID)
				fmt.Printf("  From: %s\n", comment.From.Name)
				fmt.Printf("  Message: %s\n", comment.Message)
				fmt.Printf("  Created: %s\n", comment.CreatedTime.Time.Format("2006-01-02 15:04:05"))
				fmt.Printf("  Likes: %d\n", comment.LikeCount)
				fmt.Printf("  Replies: %d\n", comment.CommentCount)
			}
		}
	}

	fmt.Println("\n🎉 Pages API example completed!")
	fmt.Println("\n📚 Available Page API operations:")
	fmt.Println("- GetPage(pageID, fields...) - Get page information")
	fmt.Println("- GetPages() - Get managed pages")
	fmt.Println("- GetPosts(pageID, limit, fields...) - Get page posts")
	fmt.Println("- GetPostComments(postID, limit, order, fields...) - Get post comments")
	fmt.Println("- GetCommentReplies(commentID, limit, fields...) - Get comment replies")
	fmt.Println("- GetComment(commentID, fields...) - Get specific comment")
}
