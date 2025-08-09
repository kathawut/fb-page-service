package facebook

import (
	"fmt"
	"net/url"
)

// ValidateAccessToken validates if the access token is valid
func (c *Client) ValidateAccessToken() error {
	resp, err := c.makeRequest("GET", "me", nil, nil)
	if err != nil {
		return fmt.Errorf("validating access token: %w", err)
	}

	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	
	if err := c.handleResponse(resp, &result); err != nil {
		return fmt.Errorf("invalid access token: %w", err)
	}

	return nil
}

// GetTokenInfo gets information about the current access token
func (c *Client) GetTokenInfo() (*TokenInfo, error) {
	params := url.Values{}
	params.Set("input_token", c.AccessToken)

	resp, err := c.makeRequest("GET", "debug_token", params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting token info: %w", err)
	}

	var tokenResp struct {
		Data TokenInfo `json:"data"`
	}
	
	if err := c.handleResponse(resp, &tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp.Data, nil
}

// GetUserInfo gets information about the current user
func (c *Client) GetUserInfo() (*User, error) {
	params := url.Values{}
	params.Set("fields", "id,name,email")

	resp, err := c.makeRequest("GET", "me", params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting user info: %w", err)
	}

	var user User
	if err := c.handleResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
