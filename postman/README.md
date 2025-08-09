# Facebook Pages API - Postman Collection

This directory contains Postman collection and environment files for testing the Facebook Pages API server with flexible authentication options.

## Files

- **`Facebook_Pages_API_v2.postman_collection.json`** - Complete Postman collection with authentication examples
- **`Facebook_Pages_API_Environment_v2.postman_environment.json`** - Environment variables for testing
- **`generate_environment.sh`** - Script to generate environment file
- **`README.md`** - This documentation file

## Quick Setup

### 1. Import into Postman

1. Open Postman
2. Click **Import** button
3. Select both JSON files:
   - `Facebook_Pages_API_v2.postman_collection.json`
   - `Facebook_Pages_API_Environment_v2.postman_environment.json`
4. Click **Import**

### 2. Set Environment

1. Select **Facebook Pages API Environment v2** from the environment dropdown
2. Update the variables with your actual values:
   - `accessToken` - Your Facebook Page Access Token
   - `pageId` - Your Facebook Page ID
   - `postId` - A valid post ID from your page
   - `commentId` - A valid comment ID

### 3. Start the Server

The server can be started with or without environment variables:

```bash
# Option 1: With environment variable (fallback authentication)
export PAGE_ACCESS_TOKEN="your_page_access_token"
./start_server.sh

# Option 2: Without environment variable (use request-level authentication)
./start_server.sh
```

## Authentication Methods

The collection demonstrates three authentication methods:

1. **Query Parameter**: `?access_token=YOUR_TOKEN`
2. **Authorization Header**: `Authorization: Bearer YOUR_TOKEN`
3. **Environment Variable**: Server-side `PAGE_ACCESS_TOKEN`

## Available Requests

### 🔍 Health Check
- **GET** `/health`
- Tests if the server is running
- No authentication required

### 📄 Page Operations

#### Get Page Information
- **GET** `/api/pages/{pageId}`
- Query Parameters:
  - `fields` - Select specific fields (optional)
- Example: `?fields=id,name,category,fan_count`

#### Get Managed Pages
- **GET** `/api/pages`
- Requires user access token (not page token)
- Returns list of pages managed by the user

### 📰 Post Operations

#### Get Page Posts
- **GET** `/api/pages/{pageId}/posts`
- Query Parameters:
  - `limit` - Number of posts (default: 10)
  - `fields` - Select specific fields (optional)
- Example: `?limit=5&fields=id,message,created_time`

### 💬 Comment Operations

#### Get Post Comments
- **GET** `/api/posts/{postId}/comments`
- Query Parameters:
  - `limit` - Number of comments (default: 10)
  - `order` - `chronological` or `reverse_chronological`
  - `fields` - Select specific fields (optional)

#### Get Specific Comment
- **GET** `/api/comments/{commentId}`
- Query Parameters:
  - `fields` - Select specific fields (optional)

#### Get Comment Replies
- **GET** `/api/comments/{commentId}/replies`
- Query Parameters:
  - `limit` - Number of replies (default: 10)
  - `fields` - Select specific fields (optional)

## Environment Variables

| Variable | Default Value | Description |
|----------|---------------|-------------|
| `baseUrl` | `http://localhost:8080` | API server base URL |
| `pageId` | `471958375999732` | Facebook Page ID for testing |
| `postId` | `471958375999732_122109395630612603` | Post ID for comment testing |
| `commentId` | `122109395630612603_433721156285059` | Comment ID for replies testing |

## Field Options

### Page Fields
```
id, name, category, about, description, website, phone, username, link,
fan_count, followers_count, talking_about_count, picture, cover,
location, hours, is_published, is_verified, can_post, access_token
```

### Post Fields
```
id, message, created_time, updated_time, type, status_type,
is_published, is_hidden, story, link, name, caption, description,
picture, permalink_url
```

### Comment Fields
```
id, message, created_time, from{id,name,picture}, like_count,
comment_count, attachment, can_like, can_comment, can_remove,
can_hide, is_hidden, is_private, parent
```

## Sample Requests

### Get Page with Specific Fields
```
GET {{baseUrl}}/api/pages/{{pageId}}?fields=id,name,category,fan_count,is_verified
```

### Get Recent Posts
```
GET {{baseUrl}}/api/pages/{{pageId}}/posts?limit=5&fields=id,message,created_time
```

### Get Comments in Chronological Order
```
GET {{baseUrl}}/api/posts/{{postId}}/comments?order=chronological&limit=10
```

## Expected Responses

### Success Response (Page Info)
```json
{
  "id": "471958375999732",
  "name": "Demo Oms",
  "category": "Software",
  "fan_count": 2,
  "is_verified": false
}
```

### Success Response (Comments)
```json
{
  "data": [
    {
      "id": "comment_id",
      "message": "Great post!",
      "created_time": "2024-11-28T07:03:02+0000",
      "from": {
        "name": "User Name",
        "id": "user_id"
      },
      "like_count": 5,
      "comment_count": 2
    }
  ],
  "summary": {
    "total_count": 8
  }
}
```

### Error Response
```json
{
  "error": "Error message description",
  "code": 500
}
```

## Testing Tips

### 1. Update Variables
Before testing, update the environment variables with your actual:
- Page ID
- Post ID (get from page posts first)
- Comment ID (get from post comments first)

### 2. Test Flow
Recommended testing order:
1. Health check
2. Get page information
3. Get page posts
4. Get post comments (using post ID from step 3)
5. Get specific comment (using comment ID from step 4)
6. Get comment replies

### 3. Authentication Errors
If you get authentication errors:
- Ensure your PAGE_ACCESS_TOKEN is valid
- Check token permissions
- Some endpoints require user tokens instead of page tokens

### 4. Field Selection
Use field selection to optimize API calls:
```
?fields=id,name,message  # Only get these fields
```

## Automated Testing

The collection includes automated tests for:
- Response time validation
- Content-Type header validation
- Status code validation

### Custom Tests
You can add custom tests in the **Tests** tab of each request:

```javascript
pm.test("Page has valid ID", function () {
    const jsonData = pm.response.json();
    pm.expect(jsonData.id).to.be.a('string');
    pm.expect(jsonData.id).to.have.length.above(0);
});

pm.test("Fan count is a number", function () {
    const jsonData = pm.response.json();
    pm.expect(jsonData.fan_count).to.be.a('number');
});
```

## Production Setup

When deploying to production:

1. Update `baseUrl` to your production domain
2. Use HTTPS URLs
3. Consider adding API key authentication
4. Update rate limiting if needed

Example production environment:
```json
{
  "baseUrl": "https://api.yourdomain.com",
  "pageId": "your_production_page_id"
}
```

## Troubleshooting

### Common Issues

1. **Server not running**
   - Error: Connection refused
   - Solution: Start the API server first

2. **Invalid Page ID**
   - Error: Page not found
   - Solution: Use a valid, accessible page ID

3. **Token issues**
   - Error: Authentication failed
   - Solution: Check PAGE_ACCESS_TOKEN environment variable

4. **Field not found**
   - Error: Field doesn't exist
   - Solution: Check available fields for Facebook API v23.0

### Debug Mode

Enable Postman Console to see detailed request/response logs:
1. View → Show Postman Console
2. Make requests and check console output
