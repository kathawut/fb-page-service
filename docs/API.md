# Facebook Pages API Go Client - API Documentation

## Overview

This Go client library provides a comprehensive interface for interacting with the Facebook Pages API (Graph API). It supports the latest API version (v18.0) and includes functionality for page management, posting, photo uploads, and insights analytics.

## Table of Contents

- [Quick Start](#quick-start)
- [Authentication](#authentication)
- [API Methods](#api-methods)
- [Data Types](#data-types)
- [Examples](#examples)
- [Error Handling](#error-handling)

## Quick Start

```go
package main

import (
    "facebook-pages-api-go/pkg/facebook"
    "fmt"
    "log"
)

func main() {
    // Create a new client
    client := facebook.NewClient("your_page_access_token")
    
    // Get page information
    page, err := client.GetPage("your_page_id")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Page: %s (ID: %s)\n", page.Name, page.ID)
}
```

## Authentication

### Access Token Requirements

You need a Facebook Page Access Token to use this API. Here's how to obtain one:

1. **Create a Facebook App**: Go to [Facebook Developers](https://developers.facebook.com/) and create a new app
2. **Get User Access Token**: Use Facebook Login to get a user access token with `pages_manage_posts` and `pages_read_engagement` permissions
3. **Exchange for Page Token**: Use the user token to get page access tokens via the `/me/accounts` endpoint

### Required Permissions

- `pages_manage_posts` - To create, edit, and delete posts
- `pages_read_engagement` - To read page insights and analytics
- `pages_manage_metadata` - To manage page information
- `pages_show_list` - To get list of managed pages

## API Methods

### Client Management

#### `NewClient(accessToken string) *Client`
Creates a new Facebook Pages API client.

```go
client := facebook.NewClient("your_access_token")
```

#### `SetAPIVersion(version string)`
Sets the API version to use (default: v18.0).

```go
client.SetAPIVersion("v19.0")
```

### Page Operations

#### `GetPage(pageID string, fields ...string) (*Page, error)`
Retrieves information about a Facebook page.

```go
// Get page with default fields
page, err := client.GetPage("123456789")

// Get page with specific fields
page, err := client.GetPage("123456789", "id", "name", "fan_count", "website")
```

#### `GetPages() ([]Page, error)`
Retrieves a list of pages that the user manages.

```go
pages, err := client.GetPages()
```

### Post Operations

#### `CreatePost(pageID string, post CreatePostRequest) (*PostResponse, error)`
Creates a new post on the specified page.

```go
post := facebook.CreatePostRequest{
    Message:   "Hello, world!",
    Link:      "https://example.com",
    Published: true,
}
response, err := client.CreatePost("page_id", post)
```

#### `GetPosts(pageID string, limit int, fields ...string) (*PostsResponse, error)`
Retrieves posts from a Facebook page.

```go
// Get last 10 posts
posts, err := client.GetPosts("page_id", 10)

// Get posts with specific fields
posts, err := client.GetPosts("page_id", 5, "id", "message", "created_time")
```

#### `GetPost(postID string, fields ...string) (*Post, error)`
Retrieves a specific post.

```go
post, err := client.GetPost("post_id")
```

#### `DeletePost(postID string) error`
Deletes a post from a Facebook page.

```go
err := client.DeletePost("post_id")
```

### Photo Operations

#### `UploadPhoto(pageID, imagePath, message string, published bool) (*PhotoResponse, error)`
Uploads a photo from a local file.

```go
response, err := client.UploadPhoto("page_id", "/path/to/image.jpg", "Caption", true)
```

#### `UploadPhotoFromReader(pageID string, reader io.Reader, message string, published bool) (*PhotoResponse, error)`
Uploads a photo from an io.Reader.

```go
file, _ := os.Open("image.jpg")
defer file.Close()
response, err := client.UploadPhotoFromReader("page_id", file, "Caption", true)
```

#### `UploadPhotoByURL(pageID, imageURL, message string, published bool) (*PhotoResponse, error)`
Uploads a photo from a URL.

```go
response, err := client.UploadPhotoByURL("page_id", "https://example.com/image.jpg", "Caption", true)
```

#### `GetPhotos(pageID string, limit int) ([]Photo, error)`
Retrieves photos from a Facebook page.

```go
photos, err := client.GetPhotos("page_id", 20)
```

#### `DeletePhoto(photoID string) error`
Deletes a photo from a Facebook page.

```go
err := client.DeletePhoto("photo_id")
```

### Insights and Analytics

#### `GetPageInsights(pageID string, metrics []string, period string, since, until *time.Time) (*InsightsResponse, error)`
Retrieves insights data for a Facebook page.

```go
since := time.Now().AddDate(0, 0, -7) // 7 days ago
until := time.Now()
metrics := []string{"page_fans", "page_impressions", "page_engaged_users"}

insights, err := client.GetPageInsights("page_id", metrics, "day", &since, &until)
```

#### `GetPostInsights(postID string, metrics []string) (*InsightsResponse, error)`
Retrieves insights data for a specific post.

```go
metrics := []string{"post_impressions", "post_engaged_users"}
insights, err := client.GetPostInsights("post_id", metrics)
```

#### `GetAvailableMetrics(pageID string) ([]string, error)`
Returns a list of available metrics for insights.

```go
metrics, err := client.GetAvailableMetrics("page_id")
```

### Utility Functions

#### `ValidateAccessToken() error`
Validates if the access token is valid.

```go
err := client.ValidateAccessToken()
```

#### `GetTokenInfo() (*TokenInfo, error)`
Gets information about the current access token.

```go
tokenInfo, err := client.GetTokenInfo()
```

#### `GetUserInfo() (*User, error)`
Gets information about the current user.

```go
user, err := client.GetUserInfo()
```

## Data Types

### Core Types

- `Page` - Represents a Facebook page
- `Post` - Represents a Facebook post
- `Photo` - Represents a Facebook photo
- `User` - Represents a Facebook user
- `Insight` - Represents analytics data

### Request Types

- `CreatePostRequest` - Parameters for creating a post
- `PhotoResponse` - Response from photo upload
- `PostResponse` - Response from post creation

### Response Types

- `PostsResponse` - Paginated list of posts
- `InsightsResponse` - Analytics data response
- `ErrorResponse` - API error response

## Examples

### Basic Page Information

```go
client := facebook.NewClient("your_token")
page, err := client.GetPage("page_id")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Page: %s\n", page.Name)
fmt.Printf("Fans: %d\n", page.FanCount)
fmt.Printf("Verified: %t\n", page.IsVerified)
```

### Creating Posts

```go
// Text post
textPost := facebook.CreatePostRequest{
    Message:   "Hello from Go!",
    Published: true,
}
response, err := client.CreatePost("page_id", textPost)

// Link post
linkPost := facebook.CreatePostRequest{
    Message:     "Check out this awesome link!",
    Link:        "https://golang.org",
    Name:        "The Go Programming Language",
    Description: "Go is an open source programming language.",
    Published:   true,
}
response, err := client.CreatePost("page_id", linkPost)
```

### Photo Upload

```go
// Upload from local file
photoResp, err := client.UploadPhoto("page_id", "photo.jpg", "My photo!", true)

// Upload from URL
photoResp, err := client.UploadPhotoByURL("page_id", "https://example.com/photo.jpg", "Remote photo!", true)
```

### Analytics and Insights

```go
// Page insights for last 30 days
since := time.Now().AddDate(0, 0, -30)
until := time.Now()
insights, err := client.GetPageInsights("page_id", nil, "day", &since, &until)

for _, insight := range insights.Data {
    fmt.Printf("Metric: %s\n", insight.Name)
    if len(insight.Values) > 0 {
        latest := insight.Values[len(insight.Values)-1]
        fmt.Printf("Latest value: %v\n", latest.Value)
    }
}
```

## Error Handling

The client returns structured errors that include Facebook API error details:

```go
page, err := client.GetPage("invalid_id")
if err != nil {
    // Error contains detailed information
    fmt.Printf("Error: %v\n", err)
    
    // You can also check for specific error types
    if strings.Contains(err.Error(), "invalid") {
        // Handle invalid ID error
    }
}
```

### Common Error Codes

- `190` - Access token issues
- `100` - Invalid parameter
- `200` - Permission denied
- `803` - Rate limiting

## Rate Limiting

Facebook has rate limits on API calls. The client doesn't implement automatic retry logic, so you should:

1. Monitor your API usage
2. Implement exponential backoff for retries
3. Cache responses when possible
4. Use batch requests for bulk operations

## Best Practices

1. **Token Security**: Never commit access tokens to version control
2. **Error Handling**: Always check errors and handle them appropriately
3. **Rate Limiting**: Respect Facebook's rate limits
4. **Pagination**: Use pagination for large datasets
5. **Field Selection**: Only request fields you need to reduce response size
6. **Caching**: Cache responses when appropriate to reduce API calls

## API Versions

This client supports Facebook Graph API v18.0 by default. You can change the version:

```go
client := facebook.NewClient("token")
client.SetAPIVersion("v19.0")
```

**Note**: Different API versions may have different features and field availability. Always refer to Facebook's API documentation for version-specific changes.

## Permissions Reference

| Permission | Required For |
|------------|-------------|
| `pages_manage_posts` | Creating, editing, deleting posts |
| `pages_read_engagement` | Reading insights and analytics |
| `pages_show_list` | Getting list of managed pages |
| `pages_manage_metadata` | Managing page information |
| `publish_pages` | Publishing content to pages |

## Contributing

When contributing to this project:

1. Follow Go coding standards
2. Add tests for new functionality
3. Update documentation
4. Handle errors appropriately
5. Maintain backward compatibility when possible
