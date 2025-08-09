package facebook

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
)

// UploadPhoto uploads a photo to a Facebook page
func (c *Client) UploadPhoto(pageID string, imagePath string, message string, published bool) (*PhotoResponse, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("opening image file: %w", err)
	}
	defer file.Close()

	return c.UploadPhotoFromReader(pageID, file, message, published)
}

// UploadPhotoFromReader uploads a photo from an io.Reader to a Facebook page
func (c *Client) UploadPhotoFromReader(pageID string, reader io.Reader, message string, published bool) (*PhotoResponse, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Add the image file
	part, err := writer.CreateFormFile("source", "image.jpg")
	if err != nil {
		return nil, fmt.Errorf("creating form file: %w", err)
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, fmt.Errorf("copying image data: %w", err)
	}

	// Add message if provided
	if message != "" {
		err = writer.WriteField("message", message)
		if err != nil {
			return nil, fmt.Errorf("writing message field: %w", err)
		}
	}

	// Add published status
	publishedStr := "true"
	if !published {
		publishedStr = "false"
	}
	err = writer.WriteField("published", publishedStr)
	if err != nil {
		return nil, fmt.Errorf("writing published field: %w", err)
	}

	// Add access token
	err = writer.WriteField("access_token", c.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("writing access token field: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("closing multipart writer: %w", err)
	}

	// Create the request
	endpoint := fmt.Sprintf("%s/photos", pageID)
	apiURL := c.buildURL(endpoint)

	req, err := c.HTTPClient.Post(apiURL, writer.FormDataContentType(), &body)
	if err != nil {
		return nil, fmt.Errorf("making photo upload request: %w", err)
	}

	var photoResp PhotoResponse
	if err := c.handleResponse(req, &photoResp); err != nil {
		return nil, err
	}

	return &photoResp, nil
}

// UploadPhotoByURL uploads a photo from a URL to a Facebook page
func (c *Client) UploadPhotoByURL(pageID string, imageURL string, message string, published bool) (*PhotoResponse, error) {
	params := url.Values{}
	params.Set("url", imageURL)
	params.Set("access_token", c.AccessToken)
	
	if message != "" {
		params.Set("message", message)
	}
	
	publishedStr := "true"
	if !published {
		publishedStr = "false"
	}
	params.Set("published", publishedStr)

	endpoint := fmt.Sprintf("%s/photos", pageID)
	resp, err := c.makeRequest("POST", endpoint, params, nil)
	if err != nil {
		return nil, fmt.Errorf("uploading photo by URL: %w", err)
	}

	var photoResp PhotoResponse
	if err := c.handleResponse(resp, &photoResp); err != nil {
		return nil, err
	}

	return &photoResp, nil
}

// GetPhotos retrieves photos from a Facebook page
func (c *Client) GetPhotos(pageID string, limit int) ([]Photo, error) {
	params := url.Values{}
	
	if limit > 0 {
		params.Set("limit", fmt.Sprintf("%d", limit))
	}
	
	params.Set("fields", "id,name,picture,source,created_time,updated_time,link")

	endpoint := fmt.Sprintf("%s/photos", pageID)
	resp, err := c.makeRequest("GET", endpoint, params, nil)
	if err != nil {
		return nil, fmt.Errorf("getting photos: %w", err)
	}

	var photosResp struct {
		Data []Photo `json:"data"`
	}
	
	if err := c.handleResponse(resp, &photosResp); err != nil {
		return nil, err
	}

	return photosResp.Data, nil
}

// DeletePhoto deletes a photo from a Facebook page
func (c *Client) DeletePhoto(photoID string) error {
	resp, err := c.makeRequest("DELETE", photoID, nil, nil)
	if err != nil {
		return fmt.Errorf("deleting photo: %w", err)
	}

	var result struct {
		Success bool `json:"success"`
	}
	
	if err := c.handleResponse(resp, &result); err != nil {
		return err
	}

	if !result.Success {
		return fmt.Errorf("failed to delete photo")
	}

	return nil
}
