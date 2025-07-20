package integration

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/knands42/lorecrafter/internal/utils"

	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func campaignInvitationSetup(t *testing.T) (
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
		false,
		domain.SettingsMetadata{},
		domain.SettingsAIMetadata{},
	)

	statusCode, _ := CreateCampaign(t, nil, userGMCookie, *campaignInput, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)

	return userGM, userGMCookie, userToBeInvited, userToBeInvitedCookie, campaign
}

func addCampaignToAPlayer(t *testing.T, playerCookie http.Cookie) {
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
		false,
		domain.SettingsMetadata{},
		domain.SettingsAIMetadata{},
	)

	var campaign *sqlc.Campaign
	statusCode, _ := CreateCampaign(t, nil, playerCookie, *campaignInput, &campaign)
	require.Equal(t, http.StatusCreated, statusCode)
}

func TestCreateCampaignInvitation_Success(t *testing.T) {
	// setup user & campaign
	userGM, userGMCookie, userToBeInvited, _, campaign := campaignInvitationSetup(t)

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

func TestCreateCampaignInvitationForAnotherPlayerInAnotherGame_Success(t *testing.T) {
	// setup user & campaign
	userGM, userGMCookie, userToBeInvited, userToBeInvitedCookie, campaign := campaignInvitationSetup(t)

	// add player2 to another campaign
	addCampaignToAPlayer(t, userToBeInvitedCookie)

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
	_, _, _, _, campaign := campaignInvitationSetup(t)

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
	_, userGMCookie, _, _, campaign := campaignInvitationSetup(t)

	// And an invalid invitation input (empty username)
	invitationInput := domain.CreateCampaignInvitationInput{
		Username: "",
	}

	// When creating a campaign invitation with invalid input
	statusCode, _ := CreateCampaignInvitation(t, userGMCookie, campaign.ID.Bytes, invitationInput, nil)

	// Then it should fail with a bad request status
	assert.Equal(t, http.StatusBadRequest, statusCode)
}

func TestCreateCampaignInvitation_Failure_UserAlreadyMember(t *testing.T) {
	// setup user & campaign
	userGM, userGMCookie, userToBeInvited, _, campaign := campaignInvitationSetup(t)

	// add user 2 to campaign
	_, err := campaignMembersUseCase.AddCampaignPlayerMember(userGM.ID, domain.CreateCampaignMemberInput{
		UserID:     userToBeInvited.ID,
		CampaignID: campaign.ID.Bytes,
		Role:       sqlc.MemberRolePlayer,
	})
	require.NoError(t, err)

	// create the campaign invitation
	invitationInput := domain.CreateCampaignInvitationInput{
		Username: userToBeInvited.Username,
	}
	var createdInvitation domain.CampaignInvitation
	statusCode, _ := CreateCampaignInvitation(t, userGMCookie, campaign.ID.Bytes, invitationInput, &createdInvitation)
	require.Equal(t, http.StatusConflict, statusCode)
}

func TestReceivingACampaignInvitation_Success(t *testing.T) {
	// setup user & campaign
	_, userGMCookie, userToBeInvited, userToBeInvitedCookie, campaign := campaignInvitationSetup(t)

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

func TestShouldInvalidatePreviousInvitationsByAcceptingOrRejecting(t *testing.T) {
	// setup user & campaign
	_, userGMCookie, userToBeInvited, userToBeInvitedCookie, campaign := campaignInvitationSetup(t)

	// create the campaign invitation
	invitationInput := domain.CreateCampaignInvitationInput{
		Username: userToBeInvited.Username,
	}
	var createdInvitation domain.CampaignInvitation
	statusCode, _ := CreateCampaignInvitation(t, userGMCookie, campaign.ID.Bytes, invitationInput, &createdInvitation)
	require.Equal(t, http.StatusCreated, statusCode)

	// create the second campaign invitation
	invitationInput = domain.CreateCampaignInvitationInput{
		Username: userToBeInvited.Username,
	}
	statusCode, _ = CreateCampaignInvitation(t, userGMCookie, campaign.ID.Bytes, invitationInput, &createdInvitation)
	require.Equal(t, http.StatusCreated, statusCode)

	// check if it has two invitations
	var campaignInvitations []struct {
		UserID uuid.UUID `json:"user_id"`
	}
	rows, err := TestDB.Query(
		context.Background(),
		`SELECT user_id FROM campaign_invitations AS ci WHERE ci.campaign_id = $1 AND ci.user_id = $2`,
		campaign.ID.Bytes, userToBeInvited.ID,
	)
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		var invitation struct {
			UserID uuid.UUID `json:"user_id"`
		}
		err := rows.Scan(&invitation.UserID)
		require.NoError(t, err)
		campaignInvitations = append(campaignInvitations, invitation)
	}
	require.NoError(t, rows.Err())
	require.Equal(t, 2, len(campaignInvitations))

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

	// check if there is no pending state invites
	time.Sleep(2 * time.Second)
	var userId any
	row := TestDB.QueryRow(
		context.Background(),
		`SELECT user_id FROM campaign_invitations AS ci WHERE ci.campaign_id = $1 AND ci.user_id = $2 AND ci.status = 'pending'`,
		campaign.ID.Bytes, userToBeInvited.ID,
	)
	err = row.Scan(&userId)
	require.Equal(t, "no rows in result set", err.Error())
}
