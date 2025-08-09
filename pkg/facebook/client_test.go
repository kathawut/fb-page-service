package facebook

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	token := "test_token"
	client := NewClient(token)
	
	if client.AccessToken != token {
		t.Errorf("Expected access token to be %s, got %s", token, client.AccessToken)
	}
	
	if client.APIVersion != DefaultAPIVersion {
		t.Errorf("Expected API version to be %s, got %s", DefaultAPIVersion, client.APIVersion)
	}
	
	if client.BaseURL != BaseURL {
		t.Errorf("Expected base URL to be %s, got %s", BaseURL, client.BaseURL)
	}
	
	if client.HTTPClient.Timeout != 30*time.Second {
		t.Errorf("Expected timeout to be 30s, got %v", client.HTTPClient.Timeout)
	}
}

func TestSetAPIVersion(t *testing.T) {
	client := NewClient("test_token")
	newVersion := "v19.0"
	
	client.SetAPIVersion(newVersion)
	
	if client.APIVersion != newVersion {
		t.Errorf("Expected API version to be %s, got %s", newVersion, client.APIVersion)
	}
}

func TestBuildURL(t *testing.T) {
	client := NewClient("test_token")
	endpoint := "me"
	expectedURL := "https://graph.facebook.com/v23.0/me"
	
	actualURL := client.buildURL(endpoint)
	
	if actualURL != expectedURL {
		t.Errorf("Expected URL to be %s, got %s", expectedURL, actualURL)
	}
}

func TestBuildURLWithCustomVersion(t *testing.T) {
	client := NewClient("test_token")
	client.SetAPIVersion("v19.0")
	endpoint := "12345/posts"
	expectedURL := "https://graph.facebook.com/v19.0/12345/posts"
	
	actualURL := client.buildURL(endpoint)
	
	if actualURL != expectedURL {
		t.Errorf("Expected URL to be %s, got %s", expectedURL, actualURL)
	}
}
