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

	if pageAccessToken == "" || pageID == "" {
		log.Fatal("Please set PAGE_ACCESS_TOKEN and PAGE_ID environment variables")
	}

	client := facebook.NewClient(pageAccessToken)
	client.SetAPIVersion("v21.0") // Use latest API version

	fmt.Println("🗨️  Facebook Page Comments Example")
	fmt.Println("==================================")

	// First, get some posts to work with
	fmt.Println("\n📰 Getting recent posts...")
	posts, err := client.GetPosts(pageID, 5, "id", "message", "created_time")
	if err != nil {
		log.Printf("Error getting posts: %v", err)
		return
	}

	if len(posts.Data) == 0 {
		fmt.Println("No posts found on this page")
		return
	}

	fmt.Printf("Found %d posts\n", len(posts.Data))

	// Get comments for each post
	for i, post := range posts.Data {
		fmt.Printf("\n--- Post %d ---\n", i+1)
		fmt.Printf("Post ID: %s\n", post.ID)
		if post.Message != "" {
			message := post.Message
			if len(message) > 100 {
				message = message[:100] + "..."
			}
			fmt.Printf("Message: %s\n", message)
		}
		fmt.Printf("Created: %s\n", post.CreatedTime.Format("2006-01-02 15:04:05"))

		// Get comments for this post
		fmt.Printf("\n💬 Getting comments for post %s...\n", post.ID)
		comments, err := client.GetPostComments(post.ID, 10, "chronological")
		if err != nil {
			log.Printf("Error getting comments for post %s: %v", post.ID, err)
			continue
		}

		if comments.Summary.TotalCount > 0 {
			fmt.Printf("Total comments: %d\n", comments.Summary.TotalCount)
			fmt.Printf("Showing %d comments:\n", len(comments.Data))

			for j, comment := range comments.Data {
				fmt.Printf("\n  Comment %d:\n", j+1)
				fmt.Printf("  ID: %s\n", comment.ID)
				fmt.Printf("  From: %s (ID: %s)\n", comment.From.Name, comment.From.ID)
				fmt.Printf("  Created: %s\n", comment.CreatedTime.Format("2006-01-02 15:04:05"))
				
				if comment.Message != "" {
					message := comment.Message
					if len(message) > 200 {
						message = message[:200] + "..."
					}
					fmt.Printf("  Message: %s\n", message)
				}
				
				fmt.Printf("  Likes: %d\n", comment.LikeCount)
				fmt.Printf("  Replies: %d\n", comment.CommentCount)
				fmt.Printf("  Can remove: %t\n", comment.CanRemove)
				fmt.Printf("  Is hidden: %t\n", comment.IsHidden)

				// Get replies if this comment has any
				// if comment.CommentCount > 0 {
				// 	fmt.Printf("\n  📝 Getting replies for comment %s...\n", comment.ID)
				// 	replies, err := client.GetCommentReplies(comment.ID, 5)
				// 	if err != nil {
				// 		log.Printf("    Error getting replies: %v", err)
				// 	} else {
				// 		fmt.Printf("    Found %d replies:\n", len(replies.Data))
				// 		for k, reply := range replies.Data {
				// 			fmt.Printf("    Reply %d:\n", k+1)
				// 			fmt.Printf("      From: %s\n", reply.From.Name)
				// 			if reply.Message != "" {
				// 				replyMsg := reply.Message
				// 				if len(replyMsg) > 100 {
				// 					replyMsg = replyMsg[:100] + "..."
				// 				}
				// 				fmt.Printf("      Message: %s\n", replyMsg)
				// 			}
				// 			fmt.Printf("      Created: %s\n", reply.CreatedTime.Format("2006-01-02 15:04:05"))
				// 		}
				// 	}
				// }
			}

			// Show pagination info if available
			if comments.Paging.Next != "" {
				fmt.Printf("\n  📄 More comments available (next page available)\n")
			}
		} else {
			fmt.Printf("No comments found for this post\n")
		}

		// Only process first 2 posts to keep output manageable
		if i >= 1 {
			fmt.Printf("\n... (showing only first 2 posts for demo)\n")
			break
		}
	}

	fmt.Println("\n🎉 Comments example completed!")
	fmt.Println("\n📋 Available comment operations:")
	fmt.Println("- GetPostComments() - Get all comments for a post")
	fmt.Println("- GetComment() - Get specific comment details")
	fmt.Println("- GetCommentReplies() - Get replies to a comment")
	fmt.Println("\n📚 Comment fields available:")
	fmt.Println("- id, message, created_time, from")
	fmt.Println("- like_count, comment_count, user_likes")
	fmt.Println("- can_like, can_comment, can_remove, can_hide")
	fmt.Println("- is_hidden, is_private, parent")
	fmt.Println("- attachment, message_tags")
}
