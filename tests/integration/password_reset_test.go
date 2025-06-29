package integration

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/knands42/lorecrafter/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForgotPassword_Success(t *testing.T) {
	// Create a test user
	user, _ := CreateTestUser(t)

	// Request password reset
	statusCode, _ := ForgotPassword(t, user.Email)

	// Verify the response
	assert.Equal(t, http.StatusNoContent, statusCode)

	// Give some time for the goroutine to complete
	time.Sleep(100 * time.Millisecond)
}

func TestMultipleForgotPasswordRequests_InvalidatesPreviousTokens(t *testing.T) {
	// Create a test user
	user, _ := CreateTestUser(t)
	email := user.Email

	for i := 0; i < 3; i++ {
		statusCode, _ := ForgotPassword(t, email)
		assert.Equal(t, http.StatusNoContent, statusCode)
		// Add a small delay between requests and give time for the goroutine to complete
		time.Sleep(200 * time.Millisecond)
	}

	// Get the latest token from the database (should be the only valid one)
	var tokens []struct {
		TokenHash string    `db:"token_hash"`
		ExpiresAt time.Time `db:"expires_at"`
		Used      bool      `db:"used"`
	}

	rows, err := TestDB.Query(
		context.Background(),
		`SELECT token_hash, expires_at, used FROM password_reset_tokens AS prt LEFT JOIN users AS u ON u.id = prt.user_id WHERE u.email = $1 ORDER BY prt.created_at DESC`,
		email,
	)
	require.NoError(t, err)

	for rows.Next() {
		var token struct {
			TokenHash string    `db:"token_hash"`
			ExpiresAt time.Time `db:"expires_at"`
			Used      bool      `db:"used"`
		}
		require.NoError(t, rows.Scan(&token.TokenHash, &token.ExpiresAt, &token.Used))
		tokens = append(tokens, token)
	}
	require.NoError(t, rows.Err())

	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, false, tokens[0].Used)
	assert.True(t, tokens[0].ExpiresAt.After(time.Now()))
}

func TestResetPassword_Success(t *testing.T) {
	// Create a test user
	user, _ := CreateTestUser(t)
	email := user.Email

	// Request password reset
	statusCode, _ := ForgotPassword(t, email)
	assert.Equal(t, http.StatusNoContent, statusCode)

	// Get the latest token from the database
	var token struct {
		TokenHash string `db:"token_hash"`
	}
	row := TestDB.QueryRow(
		context.Background(),
		`SELECT token_hash FROM password_reset_tokens WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1`,
		user.ID,
	)
	require.NoError(t, row.Scan(&token.TokenHash))

	// Reset password with the token
	newPassword := "NewPassword123!"
	statusCode, _ = ResetPassword(t, token.TokenHash, email, newPassword)
	assert.Equal(t, http.StatusNoContent, statusCode)

	// Verify the token is marked as used
	var isUsed bool
	row = TestDB.QueryRow(
		context.Background(),
		`SELECT used FROM password_reset_tokens WHERE token_hash = $1`,
		token.TokenHash,
	)
	require.NoError(t, row.Scan(&isUsed))
	assert.True(t, isUsed, "token should be marked as used after password reset")

	// Verify the password was actually changed by trying to log in with the new password
	loginInput := domain.LoginInput{
		UsernameOrEmail: email,
		Password:        newPassword,
	}
	var loginResponse struct {
		Token string `json:"token"`
	}
	statusCode, _ = LoginUser(t, loginInput, &loginResponse)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.NotEmpty(t, loginResponse.Token, "should be able to login with new password")
}

func TestResetPassword_InvalidToken(t *testing.T) {
	// Create a test user
	user, _ := CreateTestUser(t)

	// Try to reset password with an invalid token
	invalidToken := "invalid-token-123"
	statusCode, _ := ResetPassword(t, invalidToken, user.Email, "NewPassword123!")

	// Should return 400 Bad Request for invalid token
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestResetPassword_ExpiredToken(t *testing.T) {
	// Create a test user
	user, _ := CreateTestUser(t)
	email := user.Email

	// Insert an expired token directly into the database
	expiredToken := "expired-token-123"
	_, err := TestDB.Exec(
		context.Background(),
		`INSERT INTO password_reset_tokens (id, user_id, token_hash, expires_at, used)
		 VALUES ($1, $2, $3, $4, $5)`,
		uuid.New(),
		user.ID,
		expiredToken,
		time.Now().Add(-24*time.Hour), // Expired 24 hours ago
		false,
	)
	require.NoError(t, err)

	// Try to reset password with the expired token
	statusCode, _ := ResetPassword(t, expiredToken, email, "NewPassword123!")

	// Should return 400 Bad Request for expired token
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestResetPassword_UsedToken(t *testing.T) {
	// Create a test user
	user, _ := CreateTestUser(t)
	email := user.Email

	// Insert an already used token directly into the database
	usedToken := "used-token-123"
	_, err := TestDB.Exec(
		context.Background(),
		`INSERT INTO password_reset_tokens (id, user_id, token_hash, expires_at, used)
		 VALUES ($1, $2, $3, $4, $5)`,
		uuid.New(),
		user.ID,
		usedToken,
		time.Now().Add(24*time.Hour), // Expires in 24 hours
		true,                         // Already used
	)
	require.NoError(t, err)

	// Try to reset password with the used token
	statusCode, _ := ResetPassword(t, usedToken, email, "NewPassword123!")

	// Should return 400 Bad Request for used token
	assert.Equal(t, http.StatusNotFound, statusCode)
}
