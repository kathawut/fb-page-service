# Facebook Pages API Go Client - Usage Guide

## Updated Features

✅ **Latest Facebook API Version**: v23.0  
✅ **Command Line Parameters**: Pass PAGE_ACCESS_TOKEN and PAGE_ID as arguments  
✅ **Environment Variables**: Still supported as fallback  
✅ **Comment Retrieval**: Get comments from posts  
✅ **No Post Creation**: Removed create/delete post functionality as requested  

## 🚀 Quick Start

### Method 1: Command Line Parameters (Recommended)

```bash
# Basic usage
go run main.go "<PAGE_ACCESS_TOKEN>" "<PAGE_ID>"

# With custom API version
go run main.go "<PAGE_ACCESS_TOKEN>" "<PAGE_ID>" "v23.0"

# Example with your credentials
go run main.go "EAAZAiedZCs92AB..." "471958375999732" "v23.0"
```

### Method 2: Environment Variables (Fallback)

```bash
export PAGE_ACCESS_TOKEN="your_access_token"
export PAGE_ID="your_page_id"
export API_VERSION="v23.0"  # Optional, defaults to v23.0
go run main.go
```

## 📋 Available Examples

### 1. Basic Demo
```bash
go run examples/basic/main.go "<TOKEN>" "<PAGE_ID>" "v23.0"
```

### 2. Comments Demo
```bash
go run examples/comments/main.go
```

### 3. Insights Demo
```bash
go run examples/insights/main.go
```

### 4. Photos Demo
```bash
go run examples/photos/main.go
```

## 🔧 API Methods Available

### Page Operations
- `GetPage(pageID, ...fields)` - Get page information
- `GetPages()` - Get managed pages list

### Post Operations (Read-Only)
- `GetPosts(pageID, limit, ...fields)` - Get page posts
- `GetPost(postID, ...fields)` - Get specific post

### Comment Operations (New!)
- `GetPostComments(postID, limit, order, ...fields)` - Get post comments
- `GetCommentReplies(commentID, limit, ...fields)` - Get comment replies
- `GetComment(commentID, ...fields)` - Get specific comment

### Photo Operations
- `UploadPhoto()` - Upload photos from file
- `UploadPhotoByURL()` - Upload photos from URL
- `GetPhotos()` - Get page photos

### Analytics & Insights
- `GetPageInsights()` - Get page analytics
- `GetPostInsights()` - Get post analytics

### Utilities
- `ValidateAccessToken()` - Validate token
- `GetTokenInfo()` - Get token details
- `GetUserInfo()` - Get user information

## 🆕 What's New

1. **Latest API Version**: Updated to Facebook Graph API v23.0
2. **Parameter Support**: Accept tokens via command line arguments
3. **Comment Functionality**: Full comment retrieval support
4. **Removed Features**: No post creation/deletion (as requested)
5. **Better Error Handling**: Updated for v23.0 compatibility

## 📊 Example Output

```
📝 Using parameters from command line arguments
🔧 Using API version: v23.0

🚀 Facebook Pages API Go Client Demo
=====================================
📄 Page ID: 471958375999732
🔑 Token: EAAZAiedZC...XSE10Ks4jO

📄 Getting page information...
Page Name: Demo Oms
Page ID: 471958375999732
Category: Software
Fan Count: 2
Is Verified: false

📰 Getting recent posts...
Found 3 recent posts:
1. Post ID: 471958375999732_1234567890
   Message: Hello from our page...
   Created: 2025-08-09 11:20:00
   
   Getting comments...
   Total comments: 5
   Comment 1: John Doe - Great post!
   Comment 2: Jane Smith - Thanks for sharing!

🎉 Demo completed successfully!
```

## 🔒 Security Notes

- ⚠️ **Never commit access tokens** to version control
- 🔄 **Rotate tokens regularly**
- 🛡️ **Use command line parameters** for better security
- 📝 **Be careful with terminal history** when using CLI params

## 📚 Next Steps

1. Try the different examples
2. Explore comment functionality
3. Use insights for analytics
4. Build your own integrations

Happy coding! 🎉
