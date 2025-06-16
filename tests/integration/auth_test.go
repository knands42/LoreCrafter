package integration

import (
	"net/http"
	"testing"

	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegister_Success(t *testing.T) {
	// Given a valid user registration input
	username := "testuser_register_success"
	email := "testuser_register_success@example.com"
	password := "Password123!"

	input := domain.UserCreationInput{
		Username: username,
		Email:    email,
		Password: password,
	}

	// Register cleanup to delete the user after the test
	t.Cleanup(func() {
		DeleteUser(t, username)
	})

	// When registering the user
	var authOutput domain.AuthOutput
	statusCode := RegisterUser(t, input, &authOutput)

	// Then the user should be registered successfully
	assert.Equal(t, http.StatusCreated, statusCode)
	assert.NotEmpty(t, authOutput.Token)
	assert.NotEmpty(t, authOutput.User.ID)
	assert.Equal(t, username, authOutput.User.Username)
	assert.Equal(t, email, authOutput.User.Email)
}

func TestRegister_Failure_UserAlreadyExists(t *testing.T) {
	// Given an existing user
	username := "testuser_register_duplicate"
	email := "testuser_register_duplicate@example.com"
	password := "Password123!"

	input := domain.UserCreationInput{
		Username: username,
		Email:    email,
		Password: password,
	}

	// Register cleanup to delete the user after the test
	t.Cleanup(func() {
		DeleteUser(t, username)
	})

	// Register the user first
	var authOutput domain.AuthOutput
	statusCode := RegisterUser(t, input, &authOutput)
	require.Equal(t, http.StatusCreated, statusCode)

	// When trying to register the same user again
	statusCode = RegisterUser(t, input, nil)

	// Then it should fail with a conflict status
	assert.Equal(t, http.StatusConflict, statusCode)
}

func TestRegister_Failure_InvalidInput(t *testing.T) {
	// Given invalid user registration inputs
	testCases := []struct {
		name     string
		input    domain.UserCreationInput
		expected int
	}{
		{
			name: "Empty username",
			input: domain.UserCreationInput{
				Username: "",
				Email:    "test@example.com",
				Password: "Password123!",
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "Empty email",
			input: domain.UserCreationInput{
				Username: "testuser",
				Email:    "",
				Password: "Password123!",
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "Empty password",
			input: domain.UserCreationInput{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "",
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "Invalid email format",
			input: domain.UserCreationInput{
				Username: "testuser",
				Email:    "invalid-email",
				Password: "Password123!",
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "Password too short",
			input: domain.UserCreationInput{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "short",
			},
			expected: http.StatusBadRequest,
		},
	}

	// When registering with invalid inputs
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode := RegisterUser(t, tc.input, nil)

			// Then it should fail with the expected status code
			assert.Equal(t, tc.expected, statusCode)
		})
	}
}

func TestLogin_Success(t *testing.T) {
	// Given a registered user
	username := "testuser_login_success"
	email := "testuser_login_success@example.com"
	password := "Password123!"

	// Register cleanup to delete the user after the test
	t.Cleanup(func() {
		DeleteUser(t, username)
	})

	// Create the user
	input := domain.UserCreationInput{
		Username: username,
		Email:    email,
		Password: password,
	}

	var registerOutput domain.AuthOutput
	statusCode := RegisterUser(t, input, &registerOutput)
	require.Equal(t, http.StatusCreated, statusCode)

	// When logging in with valid credentials
	loginInput := domain.LoginInput{
		Username: username,
		Password: password,
	}

	var loginOutput domain.AuthOutput
	statusCode = LoginUser(t, loginInput, &loginOutput)

	// Then the login should be successful
	assert.Equal(t, http.StatusOK, statusCode)
	assert.NotEmpty(t, loginOutput.Token)
	assert.NotEmpty(t, loginOutput.User.ID)
	assert.Equal(t, username, loginOutput.User.Username)
	assert.Equal(t, email, loginOutput.User.Email)
}

func TestLogin_Failure_InvalidCredentials(t *testing.T) {
	// Given a registered user
	username := "testuser_login_failure"
	email := "testuser_login_failure@example.com"
	password := "Password123!"

	// Register cleanup to delete the user after the test
	t.Cleanup(func() {
		DeleteUser(t, username)
	})

	// Create the user
	input := domain.UserCreationInput{
		Username: username,
		Email:    email,
		Password: password,
	}

	var registerOutput domain.AuthOutput
	statusCode := RegisterUser(t, input, &registerOutput)
	require.Equal(t, http.StatusCreated, statusCode)

	// When logging in with invalid credentials
	testCases := []struct {
		name     string
		input    domain.LoginInput
		expected int
	}{
		{
			name: "Wrong username",
			input: domain.LoginInput{
				Username: "wrong_username",
				Password: password,
			},
			expected: http.StatusUnauthorized,
		},
		{
			name: "Wrong password",
			input: domain.LoginInput{
				Username: username,
				Password: "wrong_password",
			},
			expected: http.StatusUnauthorized,
		},
		{
			name: "Empty username",
			input: domain.LoginInput{
				Username: "",
				Password: password,
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "Empty password",
			input: domain.LoginInput{
				Username: username,
				Password: "",
			},
			expected: http.StatusBadRequest,
		},
	}

	// Then the login should fail with the expected status code
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode := LoginUser(t, tc.input, nil)
			assert.Equal(t, tc.expected, statusCode)
		})
	}
}

func TestMe_Success(t *testing.T) {
	// Given a registered and authenticated user
	user := CreateTestUser(t)

	// When getting the user's profile
	var responseBody interface{}
	statusCode := SendAuthenticatedRequest(t, "GET", "/api/me", user.Token, nil, &responseBody)

	// Then the profile should be retrieved successfully
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Contains(t, responseBody, "Authenticated")
	assert.Contains(t, responseBody, "User ID")
}

func TestMe_Failure_Unauthorized(t *testing.T) {
	// Given no authentication token
	// When getting the user's profile without a token
	statusCode := SendAuthenticatedRequest(t, "GET", "/api/me", "", nil, nil)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)

	// Given an invalid authentication token
	// When getting the user's profile with an invalid token
	statusCode = SendAuthenticatedRequest(t, "GET", "/api/me", "invalid_token", nil, nil)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}
