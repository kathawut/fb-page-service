# Facebook Pages API Go Client

A comprehensive Go client library for interacting with the Facebook Pages API (Graph API v18.0). This library provides easy-to-use methods for managing Facebook pages, creating posts, uploading photos, and retrieving analytics.

## 🚀 Features

- ✅ **Page Management** - Get page information and manage multiple pages
- ✅ **Post Creation** - Create text, link, and photo posts
- ✅ **Photo Uploads** - Upload photos from files or URLs
- ✅ **Analytics & Insights** - Retrieve page and post analytics
- ✅ **Access Token Validation** - Validate and get token information
- ✅ **Error Handling** - Detailed error messages and proper error handling
- ✅ **Configurable API Versions** - Support for different Facebook API versions
- ✅ **No External Dependencies** - Uses only Go standard library

## 📦 Installation

```bash
git clone https://github.com/your-username/facebook-pages-api-go.git
cd facebook-pages-api-go
go mod tidy
```

## 🔧 Quick Setup

Run the setup script for automated setup and testing:

```bash
./setup.sh
```

Or manually:

```bash
go build -v
go test ./...
```

## 🔑 Authentication

You need a Facebook Page Access Token. Here's how to get one:

1. **Create a Facebook App** at [Facebook Developers](https://developers.facebook.com/)
2. **Get User Access Token** with required permissions:
   - `pages_manage_posts` - To create/edit/delete posts
   - `pages_read_engagement` - To read insights
   - `pages_show_list` - To get managed pages
3. **Exchange for Page Token** using the `/me/accounts` endpoint

## 🏃‍♂️ Quick Start

```go
package main

import (
    "facebook-pages-api-go/pkg/facebook"
    "fmt"
    "log"
    "os"
)

func main() {
    // Get credentials from environment variables
    pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
    pageID := os.Getenv("PAGE_ID")
    
    // Create client
    client := facebook.NewClient(pageAccessToken)
    
    // Get page information
    page, err := client.GetPage(pageID)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Page: %s (ID: %s)\n", page.Name, page.ID)
    fmt.Printf("Fans: %d\n", page.FanCount)
    
    // Create a post
    post := facebook.CreatePostRequest{
        Message:   "Hello from Go! 🚀",
        Published: true,
    }
    
    response, err := client.CreatePost(pageID, post)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Post created: %s\n", response.ID)
}
```

## 📚 Examples

### Set Environment Variables

```bash
export PAGE_ACCESS_TOKEN="your_page_access_token"
export PAGE_ID="your_page_id"
export API_VERSION="v18.0"  # Optional, defaults to v18.0
```

### Run Examples

```bash
# Basic example - Page info and posting
go run main.go

# Or run specific examples
go run examples/basic/main.go      # Comprehensive demo
go run examples/photos/main.go     # Photo upload demo
go run examples/insights/main.go   # Analytics demo
```

## 🔗 API Methods

### Page Operations
- `GetPage(pageID, ...fields)` - Get page information
- `GetPages()` - Get managed pages list

### Post Operations  
- `CreatePost(pageID, post)` - Create new posts
- `GetPosts(pageID, limit, ...fields)` - Get page posts
- `GetPost(postID, ...fields)` - Get specific post
- `DeletePost(postID)` - Delete posts

### Photo Operations
- `UploadPhoto(pageID, imagePath, message, published)` - Upload from file
- `UploadPhotoFromReader(pageID, reader, message, published)` - Upload from reader
- `UploadPhotoByURL(pageID, imageURL, message, published)` - Upload from URL
- `GetPhotos(pageID, limit)` - Get page photos
- `DeletePhoto(photoID)` - Delete photos

### Analytics & Insights
- `GetPageInsights(pageID, metrics, period, since, until)` - Get page analytics
- `GetPostInsights(postID, metrics)` - Get post analytics  
- `GetAvailableMetrics(pageID)` - List available metrics

### Utilities
- `ValidateAccessToken()` - Validate token
- `GetTokenInfo()` - Get token details
- `GetUserInfo()` - Get user information
- `SetAPIVersion(version)` - Change API version

## 📊 Supported Insights Metrics

### Page Metrics
- `page_fans` - Total page likes
- `page_fan_adds` - New page likes
- `page_impressions` - Page impressions
- `page_engaged_users` - People who engaged
- `page_views_total` - Total page views
- And 30+ more metrics...

### Post Metrics  
- `post_impressions` - Post impressions
- `post_engaged_users` - People who engaged
- `post_clicks` - Total clicks
- `post_reactions_by_type_total` - Reactions breakdown
- And more...

## 🛠️ Configuration

### API Version
```go
client := facebook.NewClient("token")
client.SetAPIVersion("v19.0")  // Use different API version
```

### Custom HTTP Client
```go
client := facebook.NewClient("token")
client.HTTPClient = &http.Client{
    Timeout: 60 * time.Second,
}
```

## 📁 Project Structure

```
facebook-pages-api-go/
├── pkg/facebook/          # Core library
│   ├── client.go         # Main client
│   ├── types.go          # Data types
│   ├── pages.go          # Page operations  
│   ├── photos.go         # Photo operations
│   ├── insights.go       # Analytics
│   └── utils.go          # Utilities
├── examples/             # Usage examples
│   ├── basic/           # Basic example
│   ├── photos/          # Photo upload example
│   └── insights/        # Analytics example
├── docs/                # Documentation
│   └── API.md           # Complete API docs
├── main.go              # Simple demo
├── setup.sh             # Setup script
└── README.md            # This file
```

## 🚨 Error Handling

The client provides detailed error information:

```go
page, err := client.GetPage("invalid_id")
if err != nil {
    // Detailed error with Facebook API response
    fmt.Printf("Error: %v\n", err)
}
```

Common error codes:
- `190` - Invalid access token
- `100` - Invalid parameter  
- `200` - Permission denied
- `803` - Rate limit exceeded

## ⚡ Performance Tips

1. **Use Field Selection** - Only request needed fields
2. **Implement Caching** - Cache responses when appropriate  
3. **Batch Operations** - Use bulk operations when possible
4. **Rate Limiting** - Respect Facebook's rate limits
5. **Error Retry** - Implement exponential backoff

## 🔒 Security

- ⚠️ **Never commit access tokens** to version control
- 🔄 **Rotate tokens regularly** 
- 🛡️ **Use environment variables** for credentials
- 📝 **Log API calls** for debugging (without tokens)

## 📝 Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality  
4. Update documentation
5. Submit a pull request

## 📋 Requirements

- Go 1.21 or higher
- Facebook Page Access Token
- Valid Facebook Page ID

## 📖 Documentation

- [Complete API Documentation](docs/API.md)
- [Facebook Pages API Official Docs](https://developers.facebook.com/docs/pages-api)
- [Facebook Graph API Reference](https://developers.facebook.com/docs/graph-api)

## 🎯 Use Cases

- **Social Media Management** - Automate posting and engagement
- **Analytics Dashboards** - Build custom analytics tools
- **Content Publishing** - Automated content distribution
- **Community Management** - Manage multiple pages efficiently
- **Marketing Automation** - Integrate with marketing workflows

## 📞 Support

- 📧 Create an issue for bug reports
- 💡 Feature requests welcome
- 📚 Check documentation first
- 🔍 Search existing issues

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Facebook Graph API team for the comprehensive API
- Go community for excellent HTTP and JSON libraries
- Contributors and users of this library

---

**Happy coding! 🎉**
