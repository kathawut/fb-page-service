package facebook

import (
	"fmt"
	"net/url"
	"time"
)

// GetPageInsights retrieves insights data for a Facebook page
func (c *Client) GetPageInsights(pageID string, metrics []string, period string, since, until *time.Time) (*InsightsResponse, error) {
	params := url.Values{}
	
	// Set metrics
	if len(metrics) == 0 {
		metrics = []string{
			"page_fans",
			"page_fan_adds",
			"page_fan_removes",
			"page_views_total",
			"page_impressions",
			"page_posts_impressions",
			"page_engaged_users",
		}
	}
	
	metricsStr := ""
	for i, metric := range metrics {
		if i > 0 {
			metricsStr += ","
		}
		metricsStr += metric
	}
	params.Set("metric", metricsStr)
	
	// Set period (day, week, days_28)
	if period != "" {
		params.Set("period", period)
	} else {
		params.Set("period", "day")
	}
	
	// Set date range
	if since != nil {
		params.Set("since", since.Format("2006-01-02"))
	}
	if until != nil {
		params.Set("until", until.Format("2006-01-02"))
	}

	endpoint := fmt.Sprintf("%s/insights", pageID)
	resp, err := c.makeRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting page insights: %w", err)
	}

	var insightsResp InsightsResponse
	if err := c.handleResponse(resp, &insightsResp); err != nil {
		return nil, err
	}

	return &insightsResp, nil
}

// GetPostInsights retrieves insights data for a specific post
func (c *Client) GetPostInsights(postID string, metrics []string) (*InsightsResponse, error) {
	params := url.Values{}
	
	// Set metrics
	if len(metrics) == 0 {
		metrics = []string{
			"post_impressions",
			"post_impressions_unique",
			"post_engaged_users",
			"post_reactions_by_type_total",
			"post_clicks",
			"post_video_views",
		}
	}
	
	metricsStr := ""
	for i, metric := range metrics {
		if i > 0 {
			metricsStr += ","
		}
		metricsStr += metric
	}
	params.Set("metric", metricsStr)

	endpoint := fmt.Sprintf("%s/insights", postID)
	resp, err := c.makeRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting post insights: %w", err)
	}

	var insightsResp InsightsResponse
	if err := c.handleResponse(resp, &insightsResp); err != nil {
		return nil, err
	}

	return &insightsResp, nil
}

// GetAvailableMetrics retrieves available metrics for insights
func (c *Client) GetAvailableMetrics(pageID string) ([]string, error) {
	// Common page metrics for v21.0
	pageMetrics := []string{
		"page_fans",
		"page_fan_adds",
		"page_fan_removes", 
		"page_fan_adds_by_likes_source",
		"page_fan_adds_by_unlikes_source",
		"page_views_total",
		"page_views_logged_in_total",
		"page_views_logged_in_unique",
		"page_views_external_referrals",
		"page_impressions",
		"page_impressions_unique",
		"page_impressions_paid",
		"page_impressions_paid_unique",
		"page_impressions_organic",
		"page_impressions_organic_unique",
		"page_impressions_viral",
		"page_impressions_viral_unique",
		"page_posts_impressions",
		"page_posts_impressions_unique",
		"page_posts_impressions_paid",
		"page_posts_impressions_paid_unique",
		"page_posts_impressions_organic",
		"page_posts_impressions_organic_unique",
		"page_posts_impressions_viral",
		"page_posts_impressions_viral_unique",
		"page_engaged_users",
		"page_consumptions",
		"page_consumptions_unique",
		"page_places_checkin_total",
		"page_places_checkin_total_unique",
		"page_places_checkin_mobile",
		"page_places_checkin_mobile_unique",
		"page_negative_feedback",
		"page_negative_feedback_unique",
		"page_positive_feedback_by_type",
		"page_positive_feedback_by_type_unique",
	}
	
	return pageMetrics, nil
}