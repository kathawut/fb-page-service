package facebook

import (
	"fmt"
	"net/url"
)

// GetPage retrieves information about a Facebook page
func (c *Client) GetPage(pageID string, fields ...string) (*Page, error) {
	params := url.Values{}
	
	if len(fields) == 0 {
		fields = []string{
			"id", "name", "category", "about", "description",
			"website", "phone", "username", "link",
			"fan_count", "followers_count", 
			"talking_about_count", "picture", "cover",
			"location", "hours", "is_published", "is_verified",
			"can_post", "access_token",
		}
	}
	
	if len(fields) > 0 {
		fieldsStr := ""
		for i, field := range fields {
			if i > 0 {
				fieldsStr += ","
			}
			fieldsStr += field
		}
		params.Set("fields", fieldsStr)
	}

	resp, err := c.makeRequest("GET", pageID, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting page: %w", err)
	}

	var page Page
	if err := c.handleResponse(resp, &page); err != nil {
		return nil, err
	}

	return &page, nil
}

// GetPages retrieves a list of pages that the user manages
// Note: This requires a user access token, not a page access token
func (c *Client) GetPages() ([]Page, error) {
	params := url.Values{}
	params.Set("fields", "id,name,category,access_token,can_post")

	// Use "me/accounts" endpoint which requires user token
	resp, err := c.makeRequest("GET", "me/accounts", params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting pages (requires user access token): %w", err)
	}

	var pagesResp struct {
		Data []Page `json:"data"`
	}
	
	if err := c.handleResponse(resp, &pagesResp); err != nil {
		return nil, err
	}

	return pagesResp.Data, nil
}

// GetPosts retrieves posts from a Facebook page
func (c *Client) GetPosts(pageID string, limit int, fields ...string) (*PostsResponse, error) {
	params := url.Values{}
	
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	
	if len(fields) == 0 {
		fields = []string{
			"id", "message", "created_time",
		}
	}
	
	if len(fields) > 0 {
		fieldsStr := ""
		for i, field := range fields {
			if i > 0 {
				fieldsStr += ","
			}
			fieldsStr += field
		}
		params.Set("fields", fieldsStr)
	}

	endpoint := fmt.Sprintf("%s/posts", pageID)
	resp, err := c.makeRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting posts: %w", err)
	}

	var postsResp PostsResponse
	if err := c.handleResponse(resp, &postsResp); err != nil {
		return nil, err
	}

	return &postsResp, nil
}

// GetPostComments retrieves comments from a specific post
func (c *Client) GetPostComments(postID string, limit int, order string, fields ...string) (*CommentsResponse, error) {
	params := url.Values{}
	
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	
	if order != "" {
		params.Set("order", order)
	} else {
		params.Set("order", "reverse_chronological")
	}
	
	params.Set("summary", "true")
	
	if len(fields) == 0 {
		fields = []string{
			"id", "message", "created_time", "from{id,name,picture}",
			"like_count", "comment_count", "attachment",
		}
	}
	
	if len(fields) > 0 {
		fieldsStr := ""
		for i, field := range fields {
			if i > 0 {
				fieldsStr += ","
			}
			fieldsStr += field
		}
		params.Set("fields", fieldsStr)
	}

	endpoint := fmt.Sprintf("%s/comments", postID)
	resp, err := c.makeRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting comments: %w", err)
	}

	var commentsResp CommentsResponse
	if err := c.handleResponse(resp, &commentsResp); err != nil {
		return nil, err
	}

	return &commentsResp, nil
}

// GetCommentReplies retrieves replies to a specific comment
func (c *Client) GetCommentReplies(commentID string, limit int, fields ...string) (*CommentsResponse, error) {
	params := url.Values{}
	
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	
	if len(fields) == 0 {
		fields = []string{
			"id", "message", "created_time", "from{id,name,picture}",
			"like_count", "comment_count", "attachment",
		}
	}
	
	if len(fields) > 0 {
		fieldsStr := ""
		for i, field := range fields {
			if i > 0 {
				fieldsStr += ","
			}
			fieldsStr += field
		}
		params.Set("fields", fieldsStr)
	}

	endpoint := fmt.Sprintf("%s/comments", commentID)
	resp, err := c.makeRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting comment replies: %w", err)
	}

	var commentsResp CommentsResponse
	if err := c.handleResponse(resp, &commentsResp); err != nil {
		return nil, err
	}

	return &commentsResp, nil
}

// GetComment retrieves a specific comment by ID
func (c *Client) GetComment(commentID string, fields ...string) (*Comment, error) {
	params := url.Values{}
	
	if len(fields) == 0 {
		fields = []string{
			"id", "message", "created_time", "from{id,name,picture}",
			"like_count", "comment_count", "attachment", "parent",
		}
	}
	
	if len(fields) > 0 {
		fieldsStr := ""
		for i, field := range fields {
			if i > 0 {
				fieldsStr += ","
			}
			fieldsStr += field
		}
		params.Set("fields", fieldsStr)
	}

	resp, err := c.makeRequest("GET", commentID, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting comment: %w", err)
	}

	var comment Comment
	if err := c.handleResponse(resp, &comment); err != nil {
		return nil, err
	}

	return &comment, nil
}
