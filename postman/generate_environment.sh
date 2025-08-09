#!/bin/bash

# Facebook Pages API - Postman Environment Generator
# This script helps generate Postman environment variables based on your actual Facebook page data

echo "🔧 Facebook Pages API - Postman Environment Generator"
echo "==================================================="

# Check if server is running
if ! curl -s http://localhost:8080/health > /dev/null; then
    echo "❌ Error: API server is not running"
    echo "💡 Please start the server first:"
    echo "   export PAGE_ACCESS_TOKEN='your_token'"
    echo "   ./start_server.sh"
    exit 1
fi

echo "✅ API server is running"

# Get page ID from user
read -p "📝 Enter your Facebook Page ID: " PAGE_ID

if [ -z "$PAGE_ID" ]; then
    echo "❌ Page ID is required"
    exit 1
fi

echo "🔍 Fetching page information..."

# Get page info
PAGE_INFO=$(curl -s "http://localhost:8080/api/pages/$PAGE_ID")

if echo "$PAGE_INFO" | grep -q "error"; then
    echo "❌ Error getting page info:"
    echo "$PAGE_INFO" | jq -r '.error' 2>/dev/null || echo "$PAGE_INFO"
    exit 1
fi

PAGE_NAME=$(echo "$PAGE_INFO" | jq -r '.name' 2>/dev/null)
echo "✅ Found page: $PAGE_NAME"

echo "🔍 Fetching recent posts..."

# Get posts to find a post ID
POSTS=$(curl -s "http://localhost:8080/api/pages/$PAGE_ID/posts?limit=1")

if echo "$POSTS" | grep -q "error"; then
    echo "⚠️  Warning: Could not get posts"
    POST_ID="$PAGE_ID\_example_post_id"
else
    POST_ID=$(echo "$POSTS" | jq -r '.data[0].id' 2>/dev/null)
    if [ "$POST_ID" = "null" ] || [ -z "$POST_ID" ]; then
        echo "⚠️  Warning: No posts found"
        POST_ID="$PAGE_ID\_example_post_id"
    else
        echo "✅ Found post: $POST_ID"
    fi
fi

echo "🔍 Fetching comments..."

# Get comments to find a comment ID
COMMENT_ID="example_comment_id"
if [ "$POST_ID" != "$PAGE_ID\_example_post_id" ]; then
    COMMENTS=$(curl -s "http://localhost:8080/api/posts/$POST_ID/comments?limit=1")
    
    if ! echo "$COMMENTS" | grep -q "error"; then
        TEMP_COMMENT_ID=$(echo "$COMMENTS" | jq -r '.data[0].id' 2>/dev/null)
        if [ "$TEMP_COMMENT_ID" != "null" ] && [ -n "$TEMP_COMMENT_ID" ]; then
            COMMENT_ID="$TEMP_COMMENT_ID"
            echo "✅ Found comment: $COMMENT_ID"
        else
            echo "⚠️  Warning: No comments found"
        fi
    else
        echo "⚠️  Warning: Could not get comments"
    fi
else
    echo "⚠️  Warning: Using example comment ID"
fi

# Generate updated environment file
ENV_FILE="postman/Facebook_Pages_API_Custom.postman_environment.json"

cat > "$ENV_FILE" << EOF
{
  "id": "facebook-pages-api-custom-env",
  "name": "Facebook Pages API Custom Environment",
  "values": [
    {
      "key": "baseUrl",
      "value": "http://localhost:8080",
      "type": "default",
      "enabled": true,
      "description": "Base URL of the Facebook Pages API server"
    },
    {
      "key": "pageId",
      "value": "$PAGE_ID",
      "type": "default",
      "enabled": true,
      "description": "Your Facebook Page ID"
    },
    {
      "key": "pageName",
      "value": "$PAGE_NAME",
      "type": "default",
      "enabled": true,
      "description": "Your Facebook Page Name"
    },
    {
      "key": "postId",
      "value": "$POST_ID",
      "type": "default",
      "enabled": true,
      "description": "A recent post ID from your page"
    },
    {
      "key": "commentId",
      "value": "$COMMENT_ID",
      "type": "default",
      "enabled": true,
      "description": "A comment ID from your page"
    },
    {
      "key": "productionUrl",
      "value": "https://your-domain.com",
      "type": "default",
      "enabled": false,
      "description": "Production server URL (update when deployed)"
    }
  ],
  "_postman_variable_scope": "environment"
}
EOF

echo ""
echo "🎉 Custom Postman environment generated!"
echo "📁 File: $ENV_FILE"
echo ""
echo "📋 Your environment variables:"
echo "   Page ID: $PAGE_ID"
echo "   Page Name: $PAGE_NAME"
echo "   Post ID: $POST_ID"
echo "   Comment ID: $COMMENT_ID"
echo ""
echo "📚 Next steps:"
echo "   1. Import $ENV_FILE into Postman"
echo "   2. Select 'Facebook Pages API Custom Environment'"
echo "   3. Start testing your API!"
echo ""
echo "💡 Tip: Re-run this script if you want to update with fresh post/comment IDs"
EOF
