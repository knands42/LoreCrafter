package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/stretchr/testify/require"
)

// CreateTestUser creates a test user with random credentials
func CreateTestUser(t *testing.T) domain.AuthOutput {
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

	// Convert pgtype.UUID to uuid.UUID
	var userID uuid.UUID
	err := userID.Scan(authOutput.User.ID.Bytes)
	require.NoError(t, err)

	// Return the test user
	return authOutput
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
