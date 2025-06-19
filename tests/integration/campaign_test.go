package integration

import (
	"github.com/knands42/lorecrafter/internal/utils"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateCampaign_Success(t *testing.T) {
	// Given a registered and authenticated user
	user, cookie := CreateTestUser(t)

	// And a valid campaign creation input
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	// When creating a campaign
	var campaign sqlc.Campaign
	statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)

	// Then the campaign should be created successfully
	assert.Equal(t, http.StatusCreated, statusCode)
	assert.NotEmpty(t, campaign.ID.Bytes)
	assert.Equal(t, title, campaign.Title)
	assert.Equal(t, settingSummary, campaign.SettingSummary.String)
	assert.True(t, campaign.IsPublic)

	createdBy, _ := utils.FromPGTypeUUID(campaign.CreatedBy)
	assert.Equal(t, user.ID, createdBy)
}

func TestCreateCampaign_Failure_Unauthorized(t *testing.T) {
	// Given no authentication token
	// And a valid campaign creation input
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	// When creating a campaign without a token
	statusCode, _ := CreateCampaign(t, http.Cookie{}, *input, nil)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func TestCreateCampaign_Failure_InvalidInput(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And invalid campaign creation inputs
	testCases := []struct {
		name     string
		input    domain.CampaignCreationInput
		expected int
	}{
		{
			name: "Empty title",
			input: *domain.NewCampaignCreationInput(
				"",
				"This is a test campaign",
				"",
				sqlc.GameSystemEnumDND5E,
				6,
				"",
				true,
			),
			expected: http.StatusBadRequest,
		},
		{
			name: "Title too long",
			input: *domain.NewCampaignCreationInput(
				string(make([]byte, 256)),
				"This is a test campaign",
				"",
				sqlc.GameSystemEnumDND5E,
				6,
				"",
				true,
			),
			expected: http.StatusBadRequest,
		},
	}

	// When creating campaigns with invalid inputs
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			statusCode, _ := CreateCampaign(t, cookie, tc.input, nil)

			// Then it should fail with the expected status code
			assert.Equal(t, tc.expected, statusCode)
		})
	}
}

func TestGetCampaign_Success(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a campaign created by the user
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	var campaign sqlc.Campaign
	statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	// When getting the campaign
	var retrievedCampaign sqlc.Campaign
	statusCode, _ = GetCampaign(t, cookie, campaign.ID.Bytes, &retrievedCampaign)

	// Then the campaign should be retrieved successfully
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, campaign.ID.Bytes, retrievedCampaign.ID.Bytes)
	assert.Equal(t, title, retrievedCampaign.Title)
	assert.Equal(t, settingSummary, retrievedCampaign.SettingSummary.String)
	assert.Equal(t, campaign.IsPublic, retrievedCampaign.IsPublic)
}

func TestGetCampaign_Failure_NotFound(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a non-existent campaign ID
	nonExistentID := uuid.New()

	// When getting a non-existent campaign
	statusCode, _ := GetCampaign(t, cookie, nonExistentID, nil)

	// Then it should fail with a not found status
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestGetCampaign_Failure_Unauthorized(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a campaign created by the user
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	var campaign sqlc.Campaign
	statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	// When getting the campaign without a token
	statusCode, _ = GetCampaign(t, http.Cookie{}, campaign.ID.Bytes, nil)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func TestUpdateCampaign_Success(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a campaign created by the user
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	var campaign sqlc.Campaign
	statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	// When updating the campaign
	newTitle := "Updated Test Campaign"
	newSettingSummary := "This is an updated test campaign"
	updateInput := domain.NewUpdateCampaignInput(
		campaign.ID.Bytes,
		newTitle,
		newSettingSummary,
		"New settings",
		"no image",
		false,
		sqlc.GameSystemEnumDND5E,
		4,
		sqlc.CampaignStatusEnumACTIVE,
	)
	statusCode, _ = UpdateCampaign(t, cookie, campaign.ID.Bytes, *updateInput, nil)

	// Then the campaign should be updated successfully
	assert.Equal(t, http.StatusOK, statusCode)

	// And when getting the updated campaign
	var updatedCampaign sqlc.Campaign
	statusCode, _ = GetCampaign(t, cookie, campaign.ID.Bytes, &updatedCampaign)

	// Then it should have the updated values
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, newTitle, updatedCampaign.Title)
	assert.Equal(t, newSettingSummary, updatedCampaign.SettingSummary.String)
	assert.False(t, updatedCampaign.IsPublic)
}

func TestUpdateCampaign_Failure_NotFound(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a non-existent campaign ID
	nonExistentID := uuid.New()

	// When updating a non-existent campaign
	updateInput := domain.NewUpdateCampaignInput(
		nonExistentID,
		"New title",
		"New summary",
		"New settings",
		"no image",
		false,
		sqlc.GameSystemEnumDND5E,
		4,
		sqlc.CampaignStatusEnumACTIVE,
	)

	statusCode, _ := UpdateCampaign(t, cookie, nonExistentID, *updateInput, nil)

	// Then it should fail with a not found status
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestUpdateCampaign_Failure_Unauthorized(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a campaign created by the user
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	var campaign sqlc.Campaign
	statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	// When updating the campaign without a token
	updateInput := domain.UpdateCampaignInput{
		Title:          "Updated Test Campaign",
		SettingSummary: "This is an updated test campaign",
		IsPublic:       false,
	}

	statusCode, _ = UpdateCampaign(t, http.Cookie{}, campaign.ID.Bytes, updateInput, nil)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func TestDeleteCampaign_Success(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a campaign created by the user
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	var campaign sqlc.Campaign
	statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	// When deleting the campaign
	statusCode, _ = DeleteCampaign(t, cookie, campaign.ID.Bytes)

	// Then the campaign should be deleted successfully
	assert.Equal(t, http.StatusNoContent, statusCode)

	// And when trying to get the deleted campaign
	statusCode, _ = GetCampaign(t, cookie, campaign.ID.Bytes, nil)

	// Then it should not be found
	assert.Equal(t, http.StatusNotFound, statusCode)
}

func TestDeleteCampaign_Success_NotFound(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a non-existent campaign ID
	nonExistentID := uuid.New()

	// When deleting a non-existent campaign
	statusCode, _ := DeleteCampaign(t, cookie, nonExistentID)

	// Then it should fail with a not found status
	assert.Equal(t, http.StatusNoContent, statusCode)
}

func TestDeleteCampaign_Failure_Unauthorized(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And a campaign created by the user
	title := "Test Campaign"
	settingSummary := "This is a test campaign"
	settings := "This is a long test campaign"
	input := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
	)

	var campaign sqlc.Campaign
	statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	// When deleting the campaign without a token
	statusCode, _ = DeleteCampaign(t, http.Cookie{}, campaign.ID.Bytes)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func TestListUserCampaigns_Success(t *testing.T) {
	// Given a registered and authenticated user
	_, cookie := CreateTestUser(t)

	// And multiple campaigns created by the user
	campaignTitles := []string{
		"Test Campaign 1",
		"Test Campaign 2",
		"Test Campaign 3",
	}

	for _, title := range campaignTitles {
		settingSummary := "This is a test campaign"
		settings := "This is a long test campaign"
		input := domain.NewCampaignCreationInput(
			title,
			settingSummary,
			settings,
			sqlc.GameSystemEnumDND5E,
			6,
			"",
			true,
		)

		var campaign sqlc.Campaign
		statusCode, _ := CreateCampaign(t, cookie, *input, &campaign)
		require.Equal(t, http.StatusCreated, statusCode)
	}

	// When listing the user's campaigns
	var campaigns []sqlc.Campaign
	statusCode, _ := ListUserCampaigns(t, cookie, &campaigns)

	// Then the campaigns should be listed successfully
	assert.Equal(t, http.StatusOK, statusCode)
	assert.GreaterOrEqual(t, len(campaigns), len(campaignTitles))

	// And the list should contain the created campaigns
	titles := make(map[string]bool)
	for _, campaign := range campaigns {
		titles[campaign.Title] = true
	}

	for _, title := range campaignTitles {
		assert.True(t, titles[title], "Campaign with title %s not found", title)
	}
}

func TestListUserCampaigns_Failure_Unauthorized(t *testing.T) {
	// When listing campaigns without a token
	statusCode, _ := ListUserCampaigns(t, http.Cookie{}, nil)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}
