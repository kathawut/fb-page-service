# Facebook Pages API Server

A REST API server that provides HTTP endpoints for accessing Facebook Pages API functionality.

## Quick Start

### Environment Setup
```bash
export PAGE_ACCESS_TOKEN="your_page_access_token"
export PORT="8080"  # Optional, defaults to 8080
```

### Run the Server
```bash
# From project root
go run cmd/server/main.go

# Or build and run
go build -o server cmd/server/main.go
./server
```

## API Endpoints

### Page Operations

#### Get Page Information
```
GET /api/pages/{pageId}
```

**Query Parameters:**
- `fields` - Comma-separated list of fields to retrieve

**Example:**
```bash
curl "http://localhost:8080/api/pages/471958375999732?fields=id,name,category,fan_count"
```

**Response:**
```json
{
  "id": "471958375999732",
  "name": "Demo Oms",
  "category": "Software",
  "fan_count": 2
}
```

#### Get Managed Pages
```
GET /api/pages
```

**Example:**
```bash
curl "http://localhost:8080/api/pages"
```

**Response:**
```json
{
  "data": [
    {
      "id": "123456789",
      "name": "My Page",
      "category": "Business"
    }
  ]
}
```

### Post Operations

#### Get Page Posts
```
GET /api/pages/{pageId}/posts
```

**Query Parameters:**
- `limit` - Number of posts to retrieve (default: 10)
- `fields` - Comma-separated list of fields to retrieve

**Example:**
```bash
curl "http://localhost:8080/api/pages/471958375999732/posts?limit=5&fields=id,message,created_time"
```

**Response:**
```json
{
  "data": [
    {
      "id": "471958375999732_122109395630612603",
      "message": "Hello 😜",
      "created_time": "2024-11-26T04:54:25+0000"
    }
  ],
  "paging": {
    "next": "https://graph.facebook.com/v23.0/..."
  }
}
```

### Comment Operations

#### Get Post Comments
```
GET /api/posts/{postId}/comments
```

**Query Parameters:**
- `limit` - Number of comments to retrieve (default: 10)
- `order` - Comment order: `chronological` or `reverse_chronological` (default)
- `fields` - Comma-separated list of fields to retrieve

**Example:**
```bash
curl "http://localhost:8080/api/posts/471958375999732_122109395630612603/comments?limit=3&order=chronological"
```

**Response:**
```json
{
  "data": [
    {
      "id": "122109395630612603_433721156285059",
      "message": "Great post!",
      "created_time": "2024-11-28T07:03:02+0000",
      "from": {
        "name": "User Name",
        "id": "123456789"
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

#### Get Specific Comment
```
GET /api/comments/{commentId}
```

**Example:**
```bash
curl "http://localhost:8080/api/comments/122109395630612603_433721156285059"
```

#### Get Comment Replies
```
GET /api/comments/{commentId}/replies
```

**Query Parameters:**
- `limit` - Number of replies to retrieve (default: 10)
- `fields` - Comma-separated list of fields to retrieve

**Example:**
```bash
curl "http://localhost:8080/api/comments/122109395630612603_433721156285059/replies?limit=5"
```

### Utility Endpoints

#### Health Check
```
GET /health
```

**Response:**
```json
{
  "status": "ok",
  "service": "facebook-pages-api",
  "version": "v23.0"
}
```

## Available Fields

### Page Fields
- `id`, `name`, `category`, `about`, `description`
- `website`, `phone`, `username`, `link`
- `fan_count`, `followers_count`, `talking_about_count`
- `picture`, `cover`, `location`, `hours`
- `is_published`, `is_verified`, `can_post`

### Post Fields
- `id`, `message`, `created_time`, `updated_time`
- `type`, `status_type`, `is_published`, `is_hidden`
- `story`, `link`, `name`, `caption`, `description`
- `picture`, `permalink_url`

### Comment Fields
- `id`, `message`, `created_time`
- `from{id,name,picture}` - Nested user info
- `like_count`, `comment_count`, `attachment`
- `can_like`, `can_comment`, `can_remove`, `can_hide`
- `is_hidden`, `is_private`, `parent`

## Error Handling

All endpoints return standardized error responses:

```json
{
  "error": "Error message description",
  "code": 400
}
```

**Common HTTP Status Codes:**
- `200` - Success
- `400` - Bad Request (missing required parameters)
- `500` - Internal Server Error (Facebook API errors)

## Authentication

The server uses a single Facebook Page Access Token set via the `PAGE_ACCESS_TOKEN` environment variable. This token is used for all API calls.

**Note:** Some endpoints like `/api/pages` require a User Access Token instead of a Page Access Token.

## Rate Limiting

Facebook has rate limits on API calls. The server doesn't implement rate limiting, so consider:
- Implementing caching for frequently requested data
- Adding rate limiting middleware
- Using field selection to minimize API calls

## Docker Support

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

Build and run:
```bash
docker build -t facebook-pages-api .
docker run -p 8080:8080 -e PAGE_ACCESS_TOKEN="your_token" facebook-pages-api
```
