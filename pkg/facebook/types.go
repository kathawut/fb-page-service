package facebook

import (
	"strings"
	"time"
)

// FacebookTime is a custom time type that handles Facebook's time format
type FacebookTime struct {
	time.Time
}

// UnmarshalJSON handles Facebook's time format: "2024-11-26T04:54:25+0000"
func (ft *FacebookTime) UnmarshalJSON(data []byte) error {
	timeStr := strings.Trim(string(data), `"`)
	if timeStr == "null" || timeStr == "" {
		return nil
	}

	// Facebook uses format: "2006-01-02T15:04:05+0000"
	parsedTime, err := time.Parse("2006-01-02T15:04:05-0700", timeStr)
	if err != nil {
		return err
	}

	ft.Time = parsedTime
	return nil
}

// MarshalJSON converts time back to Facebook format
func (ft FacebookTime) MarshalJSON() ([]byte, error) {
	if ft.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + ft.Time.Format("2006-01-02T15:04:05-0700") + `"`), nil
}

// ErrorResponse represents a Facebook API error response
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail contains the error details
type ErrorDetail struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Code      int    `json:"code"`
	ErrorCode int    `json:"error_code"`
	FBTraceID string `json:"fbtrace_id"`
}

// Page represents a Facebook page
type Page struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Category          string     `json:"category"`
	CategoryList      []Category `json:"category_list,omitempty"`
	About             string     `json:"about,omitempty"`
	Description       string     `json:"description,omitempty"`
	Website           string     `json:"website,omitempty"`
	Phone             string     `json:"phone,omitempty"`
	Email             string     `json:"email,omitempty"`
	Username          string     `json:"username,omitempty"`
	Link              string     `json:"link,omitempty"`
	FanCount          int        `json:"fan_count,omitempty"`
	FollowersCount    int        `json:"followers_count,omitempty"`
	CheckinsCount     int        `json:"checkins,omitempty"`
	TalkingAboutCount int        `json:"talking_about_count,omitempty"`
	Picture           Picture    `json:"picture,omitempty"`
	CoverPhoto        CoverPhoto `json:"cover,omitempty"`
	Location          Location   `json:"location,omitempty"`
	Hours             Hours      `json:"hours,omitempty"`
	IsPublished       bool       `json:"is_published,omitempty"`
	IsVerified        bool       `json:"is_verified,omitempty"`
	CanPost           bool       `json:"can_post,omitempty"`
	AccessToken       string     `json:"access_token,omitempty"`
}

// Category represents a page category
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Picture represents page profile picture
type Picture struct {
	Data PictureData `json:"data"`
}

// PictureData contains picture details
type PictureData struct {
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	IsSilhouette bool   `json:"is_silhouette"`
	URL          string `json:"url"`
}

// CoverPhoto represents page cover photo
type CoverPhoto struct {
	ID      string `json:"id"`
	Source  string `json:"source"`
	OffsetY int    `json:"offset_y"`
	OffsetX int    `json:"offset_x"`
}

// Location represents page location
type Location struct {
	Street    string  `json:"street,omitempty"`
	City      string  `json:"city,omitempty"`
	State     string  `json:"state,omitempty"`
	Country   string  `json:"country,omitempty"`
	Zip       string  `json:"zip,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// Hours represents page hours
type Hours struct {
	Monday    []string `json:"mon_1_open,omitempty"`
	Tuesday   []string `json:"tue_1_open,omitempty"`
	Wednesday []string `json:"wed_1_open,omitempty"`
	Thursday  []string `json:"thu_1_open,omitempty"`
	Friday    []string `json:"fri_1_open,omitempty"`
	Saturday  []string `json:"sat_1_open,omitempty"`
	Sunday    []string `json:"sun_1_open,omitempty"`
}

// Post represents a Facebook page post
type Post struct {
	ID          string       `json:"id"`
	Message     string       `json:"message,omitempty"`
	Story       string       `json:"story,omitempty"`
	Link        string       `json:"link,omitempty"`
	Name        string       `json:"name,omitempty"`
	Caption     string       `json:"caption,omitempty"`
	Description string       `json:"description,omitempty"`
	Picture     string       `json:"picture,omitempty"`
	Source      string       `json:"source,omitempty"`
	Type        string       `json:"type,omitempty"`
	StatusType  string       `json:"status_type,omitempty"`
	CreatedTime FacebookTime `json:"created_time"`
	UpdatedTime FacebookTime `json:"updated_time,omitempty"`
	Permalink   string       `json:"permalink_url,omitempty"`
	IsPublished bool         `json:"is_published,omitempty"`
	IsHidden    bool         `json:"is_hidden,omitempty"`
	Privacy     Privacy      `json:"privacy,omitempty"`
	Actions     []Action     `json:"actions,omitempty"`
}

// Privacy represents post privacy settings
type Privacy struct {
	Value       string `json:"value"`
	Description string `json:"description"`
	Friends     string `json:"friends,omitempty"`
	Networks    string `json:"networks,omitempty"`
	Allow       string `json:"allow,omitempty"`
	Deny        string `json:"deny,omitempty"`
}

// Action represents post actions
type Action struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// PostResponse represents the response when creating a post
type PostResponse struct {
	ID string `json:"id"`
}

// PhotoResponse represents the response when uploading a photo
type PhotoResponse struct {
	ID     string `json:"id"`
	PostID string `json:"post_id,omitempty"`
}

// Photo represents a Facebook photo
type Photo struct {
	ID          string       `json:"id"`
	Name        string       `json:"name,omitempty"`
	Picture     string       `json:"picture,omitempty"`
	Source      string       `json:"source,omitempty"`
	CreatedTime FacebookTime `json:"created_time"`
	UpdatedTime FacebookTime `json:"updated_time,omitempty"`
	Link        string       `json:"link,omitempty"`
	Album       Album        `json:"album,omitempty"`
	Width       int          `json:"width,omitempty"`
	Height      int          `json:"height,omitempty"`
}

// Album represents a photo album
type Album struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PagingData represents pagination information
type PagingData struct {
	Cursors  Cursors `json:"cursors,omitempty"`
	Previous string  `json:"previous,omitempty"`
	Next     string  `json:"next,omitempty"`
}

// Cursors represents pagination cursors
type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// PostsResponse represents a paginated list of posts
type PostsResponse struct {
	Data   []Post     `json:"data"`
	Paging PagingData `json:"paging,omitempty"`
}

// Insight represents page insights data
type Insight struct {
	Name        string  `json:"name"`
	Period      string  `json:"period"`
	Values      []Value `json:"values"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	ID          string  `json:"id,omitempty"`
}

// Value represents insight values
type Value struct {
	Value   interface{}  `json:"value"`
	EndTime FacebookTime `json:"end_time"`
}

// InsightsResponse represents page insights response
type InsightsResponse struct {
	Data   []Insight  `json:"data"`
	Paging PagingData `json:"paging,omitempty"`
}

// TokenInfo represents access token information
type TokenInfo struct {
	AppID       string   `json:"app_id"`
	Type        string   `json:"type"`
	Application string   `json:"application"`
	DataAccess  int64    `json:"data_access_expires_at"`
	ExpiresAt   int64    `json:"expires_at"`
	IsValid     bool     `json:"is_valid"`
	Scopes      []string `json:"scopes"`
	UserID      string   `json:"user_id"`
}

// User represents a Facebook user
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// Comment represents a Facebook comment
type Comment struct {
	ID           string       `json:"id"`
	Message      string       `json:"message,omitempty"`
	CreatedTime  FacebookTime `json:"created_time"`
	From         User         `json:"from"`
	LikeCount    int          `json:"like_count,omitempty"`
	CommentCount int          `json:"comment_count,omitempty"`
	UserLikes    bool         `json:"user_likes,omitempty"`
	CanLike      bool         `json:"can_like,omitempty"`
	CanComment   bool         `json:"can_comment,omitempty"`
	CanRemove    bool         `json:"can_remove,omitempty"`
	CanHide      bool         `json:"can_hide,omitempty"`
	IsHidden     bool         `json:"is_hidden,omitempty"`
	IsPrivate    bool         `json:"is_private,omitempty"`
	ParentID     string       `json:"parent,omitempty"`
	Attachment   Attachment   `json:"attachment,omitempty"`
	MessageTags  []MessageTag `json:"message_tags,omitempty"`
	PermalinkURL string       `json:"permalink_url,omitempty"`
}

// Attachment represents a comment attachment
type Attachment struct {
	Type        string `json:"type,omitempty"`
	URL         string `json:"url,omitempty"`
	Media       Media  `json:"media,omitempty"`
	Target      Target `json:"target,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// Media represents media in an attachment
type Media struct {
	Image Image `json:"image,omitempty"`
}

// Image represents an image
type Image struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	Src    string `json:"src,omitempty"`
}

// Target represents an attachment target
type Target struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

// MessageTag represents a tag in a message
type MessageTag struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

// CommentsResponse represents a paginated list of comments
type CommentsResponse struct {
	Data    []Comment      `json:"data"`
	Paging  PagingData     `json:"paging,omitempty"`
	Summary CommentSummary `json:"summary,omitempty"`
}

// CommentSummary represents comment summary information
type CommentSummary struct {
	Order      string `json:"order,omitempty"`
	TotalCount int    `json:"total_count,omitempty"`
	CanComment bool   `json:"can_comment,omitempty"`
}
