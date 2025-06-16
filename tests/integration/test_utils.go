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
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
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

// TestUser represents a test user
type TestUser struct {
	User     sqlc.User
	Token    string
	Username string
	Email    string
	Password string
}

// CreateTestUser creates a test user with random credentials
func CreateTestUser(t *testing.T) TestUser {
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
	var authOutput domain.AuthOutput
	statusCode := RegisterUser(t, input, &authOutput)

	// Assert that the user was created successfully
	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, authOutput.Token)
	require.NotEmpty(t, authOutput.User.ID)
	require.Equal(t, username, authOutput.User.Username)
	require.Equal(t, email, authOutput.User.Email)

	// Register cleanup function to delete the user after the test
	t.Cleanup(func() {
		DeleteUser(t, username)
	})

	// Return the test user
	return TestUser{
		User:     authOutput.User,
		Token:    authOutput.Token,
		Username: username,
		Email:    email,
		Password: password,
	}
}

// RegisterUser registers a new user
func RegisterUser(t *testing.T, input domain.UserCreationInput, output interface{}) int {
	return SendRequest(t, "POST", "/api/auth/register", input, output)
}

// LoginUser logs in a user
func LoginUser(t *testing.T, input domain.LoginInput, output interface{}) int {
	return SendRequest(t, "POST", "/api/auth/login", input, output)
}

// CreateCampaign creates a new campaign
func CreateCampaign(t *testing.T, token string, input domain.CampaignCreationInput, output interface{}) int {
	return SendAuthenticatedRequest(t, "POST", "/api/campaigns", token, input, output)
}

// GetCampaign gets a campaign by ID
func GetCampaign(t *testing.T, token string, campaignID uuid.UUID, output interface{}) int {
	return SendAuthenticatedRequest(t, "GET", fmt.Sprintf("/api/campaigns/%s", campaignID), token, nil, output)
}

// UpdateCampaign updates a campaign
func UpdateCampaign(t *testing.T, token string, campaignID uuid.UUID, input domain.UpdateCampaignInput, output interface{}) int {
	return SendAuthenticatedRequest(t, "PUT", fmt.Sprintf("/api/campaigns/%s", campaignID), token, input, output)
}

// DeleteCampaign deletes a campaign
func DeleteCampaign(t *testing.T, token string, campaignID uuid.UUID) int {
	return SendAuthenticatedRequest(t, "DELETE", fmt.Sprintf("/api/campaigns/%s", campaignID), token, nil, nil)
}

// ListUserCampaigns lists all campaigns a user is a member of
func ListUserCampaigns(t *testing.T, token string, output interface{}) int {
	return SendAuthenticatedRequest(t, "GET", "/api/campaigns", token, nil, output)
}

// SendRequest sends an HTTP request to the test server
func SendRequest(t *testing.T, method, path string, body interface{}, output interface{}) int {
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

	return resp.StatusCode
}

// SendAuthenticatedRequest sends an authenticated HTTP request to the test server
func SendAuthenticatedRequest(t *testing.T, method, path, token string, body interface{}, output interface{}) int {
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
	req.Header.Set("Authorization", "Bearer "+token)

	// Send request
	resp, err := TestClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Parse response if output is provided
	if output != nil && resp.StatusCode < 300 {
		err = json.NewDecoder(resp.Body).Decode(output)
		require.NoError(t, err)
	}

	return resp.StatusCode
}
