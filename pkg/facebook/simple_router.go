package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// SimpleRouter handles HTTP routes using standard library only
type SimpleRouter struct {
	defaultClient *Client // Default client for backward compatibility
}

// NewSimpleRouter creates a new router without external dependencies
func NewSimpleRouter(accessToken string) *SimpleRouter {
	var defaultClient *Client
	if accessToken != "" {
		defaultClient = NewClient(accessToken)
	}
	return &SimpleRouter{
		defaultClient: defaultClient,
	}
}

// getClientFromRequest resolves the Facebook client from request parameters or default
func (r *SimpleRouter) getClientFromRequest(req *http.Request) (*Client, error) {
	// Try to get access token from query parameter
	accessToken := req.URL.Query().Get("access_token")
	
	// If not in query, try Authorization header
	if accessToken == "" {
		authHeader := req.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			accessToken = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}
	
	// If token provided in request, create new client
	if accessToken != "" {
		return NewClient(accessToken), nil
	}
	
	// Fall back to default client
	if r.defaultClient != nil {
		return r.defaultClient, nil
	}
	
	// Last resort: try environment variable
	envToken := os.Getenv("PAGE_ACCESS_TOKEN")
	if envToken != "" {
		return NewClient(envToken), nil
	}
	
	return nil, fmt.Errorf("no access token provided - use access_token query parameter, Authorization header, or PAGE_ACCESS_TOKEN environment variable")
}

// ServeHTTP implements http.Handler interface
func (r *SimpleRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	path := req.URL.Path
	
	switch {
	case path == "/health":
		r.healthCheck(w, req)
	case strings.HasPrefix(path, "/api/pages/") && strings.HasSuffix(path, "/posts"):
		r.getPosts(w, req)
	case strings.HasPrefix(path, "/api/pages/") && !strings.Contains(path[11:], "/"):
		r.getPage(w, req)
	case path == "/api/pages":
		r.getPages(w, req)
	case strings.HasPrefix(path, "/api/posts/") && strings.HasSuffix(path, "/comments"):
		r.getPostComments(w, req)
	case strings.HasPrefix(path, "/api/comments/") && strings.HasSuffix(path, "/replies"):
		r.getCommentReplies(w, req)
	case strings.HasPrefix(path, "/api/comments/"):
		r.getComment(w, req)
	default:
		r.writeError(w, http.StatusNotFound, "Endpoint not found")
	}
}

// extractPathParam extracts parameter from URL path
func (r *SimpleRouter) extractPathParam(path, prefix, suffix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	
	param := path[len(prefix):]
	if suffix != "" && strings.HasSuffix(param, suffix) {
		param = param[:len(param)-len(suffix)]
	}
	
	return param
}

// getPage handles GET /api/pages/{pageId}
func (r *SimpleRouter) getPage(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		r.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	pageID := r.extractPathParam(req.URL.Path, "/api/pages/", "")
	if pageID == "" {
		r.writeError(w, http.StatusBadRequest, "Page ID is required")
		return
	}
	
	// Get client from request
	client, err := r.getClientFromRequest(req)
	if err != nil {
		r.writeError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication error: %v", err))
		return
	}
	
	// Parse fields parameter
	fieldsParam := req.URL.Query().Get("fields")
	var fields []string
	if fieldsParam != "" {
		fields = strings.Split(fieldsParam, ",")
	}
	
	page, err := client.GetPage(pageID, fields...)
	if err != nil {
		r.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting page: %v", err))
		return
	}
	
	r.writeJSON(w, http.StatusOK, page)
}

// getPages handles GET /api/pages
func (r *SimpleRouter) getPages(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		r.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	// Get client from request
	client, err := r.getClientFromRequest(req)
	if err != nil {
		r.writeError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication error: %v", err))
		return
	}
	
	pages, err := client.GetPages()
	if err != nil {
		r.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting pages: %v", err))
		return
	}
	
	r.writeJSON(w, http.StatusOK, map[string]interface{}{
		"data": pages,
	})
}

// getPosts handles GET /api/pages/{pageId}/posts
func (r *SimpleRouter) getPosts(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		r.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	pageID := r.extractPathParam(req.URL.Path, "/api/pages/", "/posts")
	if pageID == "" {
		r.writeError(w, http.StatusBadRequest, "Page ID is required")
		return
	}
	
	// Get client from request
	client, err := r.getClientFromRequest(req)
	if err != nil {
		r.writeError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication error: %v", err))
		return
	}
	
	// Parse limit parameter
	limitParam := req.URL.Query().Get("limit")
	limit := 10 // default
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 {
			limit = l
		}
	}
	
	// Parse fields parameter
	fieldsParam := req.URL.Query().Get("fields")
	var fields []string
	if fieldsParam != "" {
		fields = strings.Split(fieldsParam, ",")
	}
	
	posts, err := client.GetPosts(pageID, limit, fields...)
	if err != nil {
		r.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting posts: %v", err))
		return
	}
	
	r.writeJSON(w, http.StatusOK, posts)
}

// getPostComments handles GET /api/posts/{postId}/comments
func (r *SimpleRouter) getPostComments(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		r.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	postID := r.extractPathParam(req.URL.Path, "/api/posts/", "/comments")
	if postID == "" {
		r.writeError(w, http.StatusBadRequest, "Post ID is required")
		return
	}
	
	// Get client from request
	client, err := r.getClientFromRequest(req)
	if err != nil {
		r.writeError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication error: %v", err))
		return
	}
	
	// Parse limit parameter
	limitParam := req.URL.Query().Get("limit")
	limit := 10 // default
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 {
			limit = l
		}
	}
	
	// Parse order parameter
	order := req.URL.Query().Get("order")
	if order == "" {
		order = "reverse_chronological"
	}
	
	// Parse fields parameter
	fieldsParam := req.URL.Query().Get("fields")
	var fields []string
	if fieldsParam != "" {
		fields = strings.Split(fieldsParam, ",")
	}
	
	comments, err := client.GetPostComments(postID, limit, order, fields...)
	if err != nil {
		r.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting comments: %v", err))
		return
	}
	
	r.writeJSON(w, http.StatusOK, comments)
}

// getComment handles GET /api/comments/{commentId}
func (r *SimpleRouter) getComment(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		r.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	commentID := r.extractPathParam(req.URL.Path, "/api/comments/", "")
	if commentID == "" || strings.Contains(commentID, "/") {
		r.writeError(w, http.StatusBadRequest, "Comment ID is required")
		return
	}
	
	// Get client from request
	client, err := r.getClientFromRequest(req)
	if err != nil {
		r.writeError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication error: %v", err))
		return
	}
	
	// Parse fields parameter
	fieldsParam := req.URL.Query().Get("fields")
	var fields []string
	if fieldsParam != "" {
		fields = strings.Split(fieldsParam, ",")
	}
	
	comment, err := client.GetComment(commentID, fields...)
	if err != nil {
		r.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting comment: %v", err))
		return
	}
	
	r.writeJSON(w, http.StatusOK, comment)
}

// getCommentReplies handles GET /api/comments/{commentId}/replies
func (r *SimpleRouter) getCommentReplies(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		r.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	commentID := r.extractPathParam(req.URL.Path, "/api/comments/", "/replies")
	if commentID == "" {
		r.writeError(w, http.StatusBadRequest, "Comment ID is required")
		return
	}
	
	// Get client from request
	client, err := r.getClientFromRequest(req)
	if err != nil {
		r.writeError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication error: %v", err))
		return
	}
	
	// Parse limit parameter
	limitParam := req.URL.Query().Get("limit")
	limit := 10 // default
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 {
			limit = l
		}
	}
	
	// Parse fields parameter
	fieldsParam := req.URL.Query().Get("fields")
	var fields []string
	if fieldsParam != "" {
		fields = strings.Split(fieldsParam, ",")
	}
	
	replies, err := client.GetCommentReplies(commentID, limit, fields...)
	if err != nil {
		r.writeError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting comment replies: %v", err))
		return
	}
	
	r.writeJSON(w, http.StatusOK, replies)
}

// healthCheck handles GET /health
func (r *SimpleRouter) healthCheck(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		r.writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	version := "v23.0" // Default API version
	if r.defaultClient != nil {
		version = r.defaultClient.APIVersion
	}
	
	r.writeJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"service": "facebook-pages-api",
		"version": version,
	})
}

// writeJSON writes a JSON response
func (r *SimpleRouter) writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// writeError writes an error response
func (r *SimpleRouter) writeError(w http.ResponseWriter, statusCode int, message string) {
	r.writeJSON(w, statusCode, map[string]interface{}{
		"error": message,
		"code":  statusCode,
	})
}
