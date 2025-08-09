#!/bin/bash

# Facebook Pages API - Quick Test Script

echo "🧪 Facebook Pages API - Quick Test"
echo "=================================="

# Configuration
BASE_URL=${1:-"http://localhost:8080"}
PAGE_ID=${PAGE_ID:-"471958375999732"}

echo "🌐 Testing server: $BASE_URL"
echo "📄 Using Page ID: $PAGE_ID"
echo ""

# Test 1: Health Check
echo "1. 🔍 Health Check"
echo "   GET $BASE_URL/health"
HEALTH_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/health")

if [ "$HEALTH_RESPONSE" = "200" ]; then
    echo "   ✅ Server is healthy"
    curl -s "$BASE_URL/health" | jq '.'
else
    echo "   ❌ Server health check failed (HTTP $HEALTH_RESPONSE)"
    echo "   💡 Make sure the server is running: ./start_server.sh"
    exit 1
fi

echo ""

# Test 2: Page Information
echo "2. 📄 Page Information"
echo "   GET $BASE_URL/api/pages/$PAGE_ID"
PAGE_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/pages/$PAGE_ID")

if [ "$PAGE_RESPONSE" = "200" ]; then
    echo "   ✅ Page information retrieved"
    curl -s "$BASE_URL/api/pages/$PAGE_ID?fields=id,name,category,fan_count" | jq '.'
else
    echo "   ❌ Failed to get page information (HTTP $PAGE_RESPONSE)"
    echo "   💡 Check your PAGE_ACCESS_TOKEN and PAGE_ID"
fi

echo ""

# Test 3: Page Posts
echo "3. 📰 Page Posts"
echo "   GET $BASE_URL/api/pages/$PAGE_ID/posts"
POSTS_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/pages/$PAGE_ID/posts")

if [ "$POSTS_RESPONSE" = "200" ]; then
    echo "   ✅ Posts retrieved"
    POSTS_DATA=$(curl -s "$BASE_URL/api/pages/$PAGE_ID/posts?limit=1&fields=id,message,created_time")
    echo "$POSTS_DATA" | jq '.'
    
    # Extract first post ID for comments test
    FIRST_POST_ID=$(echo "$POSTS_DATA" | jq -r '.data[0].id // empty')
    
    if [ -n "$FIRST_POST_ID" ] && [ "$FIRST_POST_ID" != "null" ]; then
        echo ""
        
        # Test 4: Post Comments
        echo "4. 💬 Post Comments"
        echo "   GET $BASE_URL/api/posts/$FIRST_POST_ID/comments"
        COMMENTS_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/posts/$FIRST_POST_ID/comments")
        
        if [ "$COMMENTS_RESPONSE" = "200" ]; then
            echo "   ✅ Comments retrieved"
            curl -s "$BASE_URL/api/posts/$FIRST_POST_ID/comments?limit=2" | jq '.'
        else
            echo "   ⚠️  Comments request failed (HTTP $COMMENTS_RESPONSE)"
            echo "   💡 This is normal if the post has no comments"
        fi
    else
        echo ""
        echo "4. 💬 Post Comments"
        echo "   ⚠️  No posts found, skipping comments test"
    fi
else
    echo "   ❌ Failed to get posts (HTTP $POSTS_RESPONSE)"
    echo "   💡 Check your access token permissions"
fi

echo ""

# Test 5: Managed Pages (will likely fail with page token)
echo "5. 📋 Managed Pages"
echo "   GET $BASE_URL/api/pages"
PAGES_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/api/pages")

if [ "$PAGES_RESPONSE" = "200" ]; then
    echo "   ✅ Managed pages retrieved"
    curl -s "$BASE_URL/api/pages" | jq '.'
else
    echo "   ⚠️  Managed pages failed (HTTP $PAGES_RESPONSE)"
    echo "   💡 This requires a user access token, not a page access token"
fi

echo ""
echo "🎉 API Test Complete!"
echo ""

# Summary
echo "📊 Test Summary:"
if [ "$HEALTH_RESPONSE" = "200" ]; then
    echo "   ✅ Health Check: PASS"
else
    echo "   ❌ Health Check: FAIL"
fi

if [ "$PAGE_RESPONSE" = "200" ]; then
    echo "   ✅ Page Info: PASS"
else
    echo "   ❌ Page Info: FAIL"
fi

if [ "$POSTS_RESPONSE" = "200" ]; then
    echo "   ✅ Page Posts: PASS"
else
    echo "   ❌ Page Posts: FAIL"
fi

echo ""
echo "🔧 Available Endpoints:"
echo "   GET  $BASE_URL/health"
echo "   GET  $BASE_URL/api/pages/{pageId}"
echo "   GET  $BASE_URL/api/pages"
echo "   GET  $BASE_URL/api/pages/{pageId}/posts"
echo "   GET  $BASE_URL/api/posts/{postId}/comments"
echo "   GET  $BASE_URL/api/comments/{commentId}"
echo "   GET  $BASE_URL/api/comments/{commentId}/replies"
echo ""
echo "📚 For more testing, use:"
echo "   make test-api              # Built-in test client"
echo "   make postman               # Generate Postman environment"
echo "   curl examples above        # Manual testing"
