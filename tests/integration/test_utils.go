package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/stretchr/testify/require"
)

// DeleteUser deletes a user from the database by username
func DeleteUser(t *testing.T, username string) {
	// Use the TestDB connection to delete the user
	_, err := TestDB.Exec(context.Background(), "DELETE FROM users WHERE username = $1", username)
	if err != nil {
		t.Logf("Error deleting user %s: %v", username, err)
	}
}

// CreateTestUser creates a test user with random credentials
func CreateTestUser(t *testing.T) (domain.User, http.Cookie) {
	// Generate random username and email
	username := fmt.Sprintf("testuser_%s", uuid.New().String()[:8])
	email := fmt.Sprintf("%s@example.com", username)
	password := "Password123!"

	// Create user registration input
	input := domain.UserCreationInput{
		Username: username,
		Email:    email,
		Password: password,
	}

	// Register the user
	var createdUser domain.User
	statusCode, _ := RegisterUser(t, input, &createdUser)
	_, cookies := LoginUser(t, domain.LoginInput{
		Username: username,
		Password: password,
	}, nil)

	// Assert that the user was created successfully
	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, createdUser.ID)
	require.Equal(t, username, createdUser.Username)
	require.Equal(t, email, createdUser.Email)

	for _, cookie := range cookies {
		if cookie.Name == "auth_token" {
			return createdUser, *cookie
		}
	}

	return domain.User{}, http.Cookie{}
}

// RegisterUser registers a new user
func RegisterUser(t *testing.T, input domain.UserCreationInput, output interface{}) (int, []*http.Cookie) {
	return SendRequest(t, "POST", "/api/auth/register", input, output)
}

// LoginUser logs in a user
func LoginUser(t *testing.T, input domain.LoginInput, output interface{}) (int, []*http.Cookie) {
	return SendRequest(t, "POST", "/api/auth/login", input, output)
}

// CreateCampaign creates a new campaign
func CreateCampaign(t *testing.T, cookie http.Cookie, input domain.CampaignCreationInput, output interface{}) (int, []*http.Cookie) {
	return SendAuthenticatedRequest(t, "POST", "/api/campaigns", cookie, input, output)
}

// GetCampaign gets a campaign by ID
func GetCampaign(t *testing.T, cookie http.Cookie, campaignID uuid.UUID, output interface{}) (int, []*http.Cookie) {
	return SendAuthenticatedRequest(t, "GET", fmt.Sprintf("/api/campaigns/%s", campaignID), cookie, nil, output)
}

// UpdateCampaign updates a campaign
func UpdateCampaign(t *testing.T, cookie http.Cookie, campaignID uuid.UUID, input domain.UpdateCampaignInput, output interface{}) (int, []*http.Cookie) {
	return SendAuthenticatedRequest(t, "PUT", fmt.Sprintf("/api/campaigns/%s", campaignID), cookie, input, output)
}

// DeleteCampaign deletes a campaign
func DeleteCampaign(t *testing.T, cookie http.Cookie, campaignID uuid.UUID) (int, []*http.Cookie) {
	return SendAuthenticatedRequest(t, "DELETE", fmt.Sprintf("/api/campaigns/%s", campaignID), cookie, nil, nil)
}

// ListUserCampaigns lists all campaigns a user is a member of
func ListUserCampaigns(t *testing.T, cookie http.Cookie, output interface{}) (int, []*http.Cookie) {
	return SendAuthenticatedRequest(t, "GET", "/api/campaigns", cookie, nil, output)
}

// SendRequest sends an HTTP request to the test server
func SendRequest(t *testing.T, method, path string, body interface{}, output interface{}) (int, []*http.Cookie) {
	// Create request body
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		require.NoError(t, err)
	}

	// Create request
	req, err := http.NewRequest(method, TestServer.URL+path, bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := TestClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Parse response if output is provided
	if output != nil && resp.StatusCode < 300 {
		err = json.NewDecoder(resp.Body).Decode(output)
		require.NoError(t, err)
	}

	return resp.StatusCode, resp.Cookies()
}

// SendAuthenticatedRequest sends an authenticated HTTP request to the test server
func SendAuthenticatedRequest(t *testing.T, method string, path string, cookie http.Cookie, body interface{}, output interface{}) (int, []*http.Cookie) {
	// Create request body
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		require.NoError(t, err)
	}

	// Create request
	req, err := http.NewRequest(method, TestServer.URL+path, bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&cookie)

	// Send request
	resp, err := TestClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Parse response if output is provided
	if output != nil && resp.StatusCode < 300 {
		err = json.NewDecoder(resp.Body).Decode(output)
		require.NoError(t, err)
	}

	return resp.StatusCode, resp.Cookies()
}
