package facebook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultAPIVersion = "v23.0"
	BaseURL          = "https://graph.facebook.com"
)

// Client represents a Facebook Pages API client
type Client struct {
	AccessToken string
	APIVersion  string
	HTTPClient  *http.Client
	BaseURL     string
}

// NewClient creates a new Facebook Pages API client
func NewClient(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
		APIVersion:  DefaultAPIVersion,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		BaseURL: BaseURL,
	}
}

// SetAPIVersion sets the API version to use
func (c *Client) SetAPIVersion(version string) {
	c.APIVersion = version
}

// buildURL constructs the full API URL
func (c *Client) buildURL(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", c.BaseURL, c.APIVersion, endpoint)
}

// makeRequest performs an HTTP request to the Facebook API
func (c *Client) makeRequest(method, endpoint string, params url.Values, body io.Reader) (*http.Response, error) {
	apiURL := c.buildURL(endpoint)
	
	req, err := http.NewRequest(method, apiURL, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Add access token to query parameters
	if params == nil {
		params = url.Values{}
	}
	params.Set("access_token", c.AccessToken)

	// Add query parameters to URL
	if len(params) > 0 {
		req.URL.RawQuery = params.Encode()
	}

	// Set content type for POST requests with body
	if method == "POST" && body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	return resp, nil
}

// handleResponse processes the API response and handles errors
func (c *Client) handleResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var errorResp ErrorResponse
		if err := json.Unmarshal(body, &errorResp); err != nil {
			return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("API error: %s (code: %d)", errorResp.Error.Message, errorResp.Error.Code)
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("unmarshaling response: %w", err)
		}
	}

	return nil
}
