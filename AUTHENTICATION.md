# Facebook Pages API - Authentication Guide

This document explains the flexible authentication options available in the Facebook Pages API v2.

## Authentication Methods

The API now supports three different authentication methods, in order of priority:

### 1. Query Parameter (Highest Priority)
Pass the access token as a query parameter in the URL:
```
GET /api/pages/123456?access_token=YOUR_TOKEN_HERE
```

### 2. Authorization Header (Medium Priority)
Pass the access token in the Authorization header using Bearer token format:
```
Authorization: Bearer YOUR_TOKEN_HERE
```

### 3. Environment Variable (Fallback)
Set the `PAGE_ACCESS_TOKEN` environment variable on the server:
```bash
export PAGE_ACCESS_TOKEN=YOUR_TOKEN_HERE
./server
```

## Examples

### Using Query Parameter
```bash
curl "http://localhost:8080/api/pages/471958375999732?access_token=EAABs..."
```

### Using Authorization Header
```bash
curl -H "Authorization: Bearer EAABs..." \
     "http://localhost:8080/api/pages/471958375999732"
```

### Using Environment Variable
```bash
export PAGE_ACCESS_TOKEN=EAABs...
curl "http://localhost:8080/api/pages/471958375999732"
```

## Postman Collection

The project includes a comprehensive Postman collection:

**Facebook_Pages_API_v2.postman_collection.json** - Complete collection with authentication examples

### Collection Features:
- Examples for all three authentication methods
- Organized by authentication type
- Pre-request and test scripts
- Environment variables for easy token management
- Mixed authentication examples showing different approaches

## Environment Variables

Use the provided environment file:
- `Facebook_Pages_API_Environment_v2.postman_environment.json`

Update the `accessToken` variable with your actual Facebook Page Access Token.

## API Endpoints

All endpoints support flexible authentication:

- `GET /health` - Health check (no auth required)
- `GET /api/pages` - List all pages
- `GET /api/pages/{pageId}` - Get specific page
- `GET /api/pages/{pageId}/posts` - Get page posts
- `GET /api/posts/{postId}/comments` - Get post comments
- `GET /api/comments/{commentId}` - Get specific comment
- `GET /api/comments/{commentId}/replies` - Get comment replies

## Error Handling

If no valid access token is provided via any method, the API returns:

```json
{
  "error": "Authentication error: no access token provided - use access_token query parameter, Authorization header, or PAGE_ACCESS_TOKEN environment variable"
}
```

## Security Considerations

1. **Query Parameters**: Convenient but visible in logs and browser history
2. **Authorization Headers**: More secure, not logged in most web servers
3. **Environment Variables**: Most secure for single-token deployments

Choose the authentication method that best fits your security requirements and use case.
