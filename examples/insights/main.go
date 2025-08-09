package main

import (
	"facebook-pages-api-go/pkg/facebook"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	pageID := os.Getenv("PAGE_ID")

	if pageAccessToken == "" || pageID == "" {
		log.Fatal("Please set PAGE_ACCESS_TOKEN and PAGE_ID environment variables")
	}

	client := facebook.NewClient(pageAccessToken)

	// Get page insights for the last 30 days
	since := time.Now().AddDate(0, 0, -30) // 30 days ago
	until := time.Now()

	// Common page metrics
	metrics := []string{
		"page_fans",
		"page_fan_adds",
		"page_fan_removes",
		"page_views_total",
		"page_impressions",
		"page_engaged_users",
	}

	insights, err := client.GetPageInsights(pageID, metrics, "day", &since, &until)
	if err != nil {
		log.Fatalf("Error getting insights: %v", err)
	}

	fmt.Printf("Page Insights (Last 30 days)\n")
	fmt.Printf("=============================\n\n")

	for _, insight := range insights.Data {
		fmt.Printf("📊 %s\n", insight.Name)
		fmt.Printf("Period: %s\n", insight.Period)
		
		if len(insight.Values) > 0 {
			fmt.Printf("Data points: %d\n", len(insight.Values))
			
			// Show first and last values
			firstValue := insight.Values[0]
			lastValue := insight.Values[len(insight.Values)-1]
			
			fmt.Printf("First value: %v (Date: %s)\n", 
				firstValue.Value, 
				firstValue.EndTime.Format("2006-01-02"))
			
			fmt.Printf("Last value: %v (Date: %s)\n", 
				lastValue.Value, 
				lastValue.EndTime.Format("2006-01-02"))
		}
		
		fmt.Println()
	}

	// Get available metrics
	fmt.Println("Available Metrics:")
	fmt.Println("==================")
	availableMetrics, err := client.GetAvailableMetrics(pageID)
	if err != nil {
		log.Printf("Error getting available metrics: %v", err)
	} else {
		for i, metric := range availableMetrics {
			fmt.Printf("%d. %s\n", i+1, metric)
		}
	}
}
