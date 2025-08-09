# Facebook Pages API Example

This example demonstrates how to use the Facebook Pages API client to interact with Facebook pages.

## Features Demonstrated

- **Page Information**: Get detailed information about a Facebook page
- **Managed Pages**: List all pages that the user manages (requires user access token)
- **Posts**: Retrieve recent posts from a page
- **Comments**: Get comments on posts and replies to comments
- **Field Selection**: Demonstrate how to request specific fields to optimize API calls

## Usage

### Environment Variables
```bash
export PAGE_ID="your_page_id"
export PAGE_ACCESS_TOKEN="your_page_access_token"
go run main.go
```

### Command Line Arguments
```bash
go run main.go <PAGE_ID> <PAGE_ACCESS_TOKEN>
```

### Build and Run
```bash
# Build the example
go build -o pages main.go

# Run with environment variables
./pages

# Run with command line arguments
./pages 471958375999732 EAAYour_Access_Token_Here
```

## API Operations Demonstrated

### 1. Page Operations
- `GetPage()` - Get comprehensive page information
- `GetPages()` - List managed pages (for user tokens)
- Custom field selection for optimized requests

### 2. Post Operations
- `GetPosts()` - Retrieve recent posts with pagination
- Field selection for posts (id, message, created_time, etc.)
- Post metadata (type, publication status, etc.)

### 3. Comment Operations
- `GetPostComments()` - Get comments on posts
- `GetCommentReplies()` - Get replies to comments
- `GetComment()` - Get specific comment details
- Comment ordering (chronological, reverse_chronological)

## Field Options

### Page Fields
- `id`, `name`, `category`, `about`, `description`
- `website`, `phone`, `username`, `link`
- `fan_count`, `followers_count`, `talking_about_count`
- `picture`, `cover`, `location`, `hours`
- `is_published`, `is_verified`, `can_post`

### Post Fields
- `id`, `message`, `story`, `link`, `name`
- `caption`, `description`, `picture`, `type`
- `status_type`, `created_time`, `updated_time`
- `permalink_url`, `is_published`, `is_hidden`

### Comment Fields
- `id`, `message`, `created_time`
- `from{id,name,picture}` - Nested user info
- `like_count`, `comment_count`, `attachment`
- `can_like`, `can_comment`, `can_remove`, `can_hide`
- `is_hidden`, `is_private`, `parent`

## API Version

This example uses Facebook Graph API **v23.0** (latest version).

## Authentication

- **Page Access Token**: Use for page-specific operations
- **User Access Token**: Required for `GetPages()` to list managed pages

## Error Handling

The example includes comprehensive error handling for:
- Network errors
- Authentication errors
- API limit errors
- Missing data scenarios

## Pagination

The Facebook API returns paginated results. The example shows:
- How to set limits on API calls
- How to check for additional pages of data
- Pagination information in responses

## Rate Limiting

Facebook has rate limits. Consider:
- Making fewer API calls by requesting multiple fields at once
- Using appropriate limits on data requests
- Implementing retry logic with exponential backoff for production use
