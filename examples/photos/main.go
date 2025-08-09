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

	// Upload a photo from file
	photoPath := "example.jpg" // Make sure this file exists
	if _, err := os.Stat(photoPath); os.IsNotExist(err) {
		fmt.Printf("Photo file %s does not exist. Using URL upload instead.\n", photoPath)
		
		// Upload photo from URL
		photoURL := "https://via.placeholder.com/600x400/ff6b6b/ffffff?text=Go+Facebook+API"
		photoResp, err := client.UploadPhotoByURL(pageID, photoURL, "Photo uploaded using Facebook Pages API Go client! 📸", true)
		if err != nil {
			log.Fatalf("Error uploading photo by URL: %v", err)
		}
		
		fmt.Printf("Photo uploaded successfully!\n")
		fmt.Printf("Photo ID: %s\n", photoResp.ID)
		if photoResp.PostID != "" {
			fmt.Printf("Post ID: %s\n", photoResp.PostID)
		}
	} else {
		photoResp, err := client.UploadPhoto(pageID, photoPath, "Photo uploaded from local file!", true)
		if err != nil {
			log.Fatalf("Error uploading photo: %v", err)
		}
		
		fmt.Printf("Photo uploaded successfully!\n")
		fmt.Printf("Photo ID: %s\n", photoResp.ID)
	}

	// Get photos
	photos, err := client.GetPhotos(pageID, 10)
	if err != nil {
		log.Printf("Error getting photos: %v", err)
	} else {
		fmt.Printf("\nFound %d photos:\n", len(photos))
		for i, photo := range photos {
			fmt.Printf("%d. Photo ID: %s\n", i+1, photo.ID)
			if photo.Name != "" {
				fmt.Printf("   Name: %s\n", photo.Name)
			}
			if photo.Source != "" {
				fmt.Printf("   Source: %s\n", photo.Source)
			}
			fmt.Printf("   Created: %s\n", photo.CreatedTime.Format("2006-01-02 15:04:05"))
		}
	}
}
