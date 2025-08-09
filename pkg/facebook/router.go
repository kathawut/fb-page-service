package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Router handles HTTP routes for Facebook Pages API
type Router struct {
	defaultClient *Client
}

// NewRouter creates a new router with a default Facebook client
func NewRouter(defaultAccessToken string) *Router {
	var client *Client
	if defaultAccessToken != "" {
		client = NewClient(defaultAccessToken)
	}
	return &Router{
		defaultClient: client,
	}
}

// getClientFromRequest creates a client from request parameters or uses default
func (r *Router) getClientFromRequest(req *http.Request) (*Client, error) {
	// Try to get access token from query parameter first
	accessToken := req.URL.Query().Get("access_token")
	
	// If not in query, try Authorization header
	if accessToken == "" {
		authHeader := req.Header.Get("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			accessToken = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}
	
	// If still not found, use default client
	if accessToken == "" {
		if r.defaultClient == nil {
			return nil, fmt.Errorf("no access token provided and no default token configured")
		}
		return r.defaultClient, nil
	}
	
	// Create new client with provided token
	return NewClient(accessToken), nil
}

// SetupRoutes configures all the API routes
func (r *Router) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	
	// Page routes
	router.HandleFunc("/api/pages/{pageId}", r.getPage).Methods("GET")
	router.HandleFunc("/api/pages", r.getPages).Methods("GET")
	
	// Post routes
	router.HandleFunc("/api/pages/{pageId}/posts", r.getPosts).Methods("GET")
	
	// Comment routes
	router.HandleFunc("/api/posts/{postId}/comments", r.getPostComments).Methods("GET")
	router.HandleFunc("/api/comments/{commentId}", r.getComment).Methods("GET")
	router.HandleFunc("/api/comments/{commentId}/replies", r.getCommentReplies).Methods("GET")
	
	// Health check
	router.HandleFunc("/health", r.healthCheck).Methods("GET")
	
	return router
}

// getPage handles GET /api/pages/{pageId}
func (r *Router) getPage(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pageID := vars["pageId"]
	
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
func (r *Router) getPages(w http.ResponseWriter, req *http.Request) {
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
func (r *Router) getPosts(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pageID := vars["pageId"]
	
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
func (r *Router) getPostComments(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	postID := vars["postId"]
	
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
func (r *Router) getComment(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	commentID := vars["commentId"]
	
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
func (r *Router) getCommentReplies(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	commentID := vars["commentId"]
	
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
func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
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
func (r *Router) writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// writeError writes an error response
func (r *Router) writeError(w http.ResponseWriter, statusCode int, message string) {
	r.writeJSON(w, statusCode, map[string]interface{}{
		"error": message,
		"code":  statusCode,
	})
}
