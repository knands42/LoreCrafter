package integration

import (
	"net/http"
	"testing"

	"github.com/knands42/lorecrafter/internal/utils"

	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupUserAndACampaign(t *testing.T) (
	userGM domain.User,
	userGMCookie http.Cookie,
	userToBeInvited domain.User,
	userToBeInvitedCookie http.Cookie,
	campaign *sqlc.Campaign,
) {
	// create a user that will be a GM
	userGM, userGMCookie = CreateTestUser(t)
	userToBeInvited, userToBeInvitedCookie = CreateTestUser(t)

	// create a campaign
	title := "Test Campaign for Invitation"
	settingSummary := "This is a test campaign for invitation"
	settings := "This is a long test campaign for invitation"
	campaignInput := domain.NewCampaignCreationInput(
		title,
		settingSummary,
		settings,
		sqlc.GameSystemEnumDND5E,
		6,
		"",
		true,
		domain.SettingsMetadata{},
		domain.SettingsAIMetadata{},
	)

	statusCode, _ := CreateCampaign(t, nil, userGMCookie, *campaignInput, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	return userGM, userGMCookie, userToBeInvited, userToBeInvitedCookie, campaign
}

func TestCreateCampaignInvitation_Success(t *testing.T) {
	// setup user & campaign
	userGM, userGMCookie, userToBeInvited, _, campaign := setupUserAndACampaign(t)

	// create the campaign invitation
	invitationInput := domain.CreateCampaignInvitationInput{
		Username: userToBeInvited.Username,
	}
	var createdInvitation domain.CampaignInvitation
	statusCode, _ := CreateCampaignInvitation(t, userGMCookie, campaign.ID.Bytes, invitationInput, &createdInvitation)

	// Then the invitation should be created successfully
	campaignUUID, err := utils.FromPGTypeUUID(campaign.ID)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, statusCode)
	assert.NotEmpty(t, createdInvitation)
	assert.NotEmpty(t, createdInvitation.ID)
	assert.NotEmpty(t, createdInvitation.CreatedAt)
	assert.NotEmpty(t, createdInvitation.UpdatedAt)
	assert.Equal(t, campaignUUID, createdInvitation.CampaignID)
	assert.Equal(t, userToBeInvited.ID, createdInvitation.UserID)
	assert.Equal(t, userGM.ID, createdInvitation.InvitedBy)
}

func TestCreateCampaignInvitation_Failure_Unauthorized(t *testing.T) {
	// setup user & campaign
	_, _, _, _, campaign := setupUserAndACampaign(t)

	// And a valid invitation input
	invitationInput := domain.CreateCampaignInvitationInput{
		Username: "testuser_to_invite",
	}

	// When creating a campaign invitation without authentication
	statusCode, _ := CreateCampaignInvitation(t, http.Cookie{}, campaign.ID.Bytes, invitationInput, nil)

	// Then it should fail with an unauthorized status
	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func TestCreateCampaignInvitation_Failure_InvalidInput(t *testing.T) {
	// setup user & campaign
	_, userGMCookie, _, _, campaign := setupUserAndACampaign(t)

	// And an invalid invitation input (empty username)
	invitationInput := domain.CreateCampaignInvitationInput{
		Username: "",
	}

	// When creating a campaign invitation with invalid input
	statusCode, _ := CreateCampaignInvitation(t, userGMCookie, campaign.ID.Bytes, invitationInput, nil)

	// Then it should fail with a bad request status
	assert.Equal(t, http.StatusBadRequest, statusCode)
}

// TODO
func TestCreateCampaignInvitation_Failure_UserAlreadyMember(t *testing.T) {}

func TestReceivingACampaignInvitation_Success(t *testing.T) {
	// setup user & campaign
	_, userGMCookie, userToBeInvited, userToBeInvitedCookie, campaign := setupUserAndACampaign(t)

	// create the campaign invitation
	invitationInput := domain.CreateCampaignInvitationInput{
		Username: userToBeInvited.Username,
	}
	var createdInvitation domain.CampaignInvitation
	statusCode, _ := CreateCampaignInvitation(t, userGMCookie, campaign.ID.Bytes, invitationInput, &createdInvitation)
	require.Equal(t, http.StatusCreated, statusCode)

	// receive the campaign invitation
	receiveInvitationInput := domain.ReceiveCampaignInvite{
		Token:    createdInvitation.Token,
		Accepted: true,
	}
	statusCode, _ = ReceiveCampaignInvitation(t, userToBeInvitedCookie, createdInvitation.Token, receiveInvitationInput, nil)
	require.Equal(t, http.StatusNoContent, statusCode)

	// check if the user is a member of the campaign
	var campaignMember domain.CampaignMember
	statusCode, _ = GetCampaignMember(t, userToBeInvitedCookie, campaign.ID.Bytes, userToBeInvited.ID, &campaignMember)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, userToBeInvited.ID, campaignMember.UserID)
}

// TODO
func TestShouldInvalidatePreviousInvitationsByAcceptingOrRejecting(t *testing.T) {}
