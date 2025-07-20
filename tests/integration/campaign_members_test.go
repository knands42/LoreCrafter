package integration

import (
	"net/http"
	"testing"

	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"github.com/stretchr/testify/require"
)

func campaignMemberSetup(t *testing.T) (
	userGM domain.User,
	userGMCookie http.Cookie,
	campaign *sqlc.Campaign,
) {
	// create a user that will be a GM
	userGM, userGMCookie = CreateTestUser(t)

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

	return userGM, userGMCookie, campaign
}

func TestGetACampaignMember_Success(t *testing.T) {
	// setup user & campaign
	userGM, userCookies, campaign := campaignMemberSetup(t)
	user2, user2Cookies := CreateTestUser(t)

	// add a new member
	_, err := campaignMembersUseCase.AddCampaignPlayerMember(userGM.ID, domain.CreateCampaignMemberInput{
		CampaignID: campaign.ID.Bytes,
		UserID:     user2.ID,
		Role:       sqlc.MemberRolePlayer,
	})
	require.NoError(t, err)

	// get the new member logged as the GM
	var campaignMember domain.CampaignMember
	statusCode, _ := GetCampaignMember(t, userCookies, campaign.ID.Bytes, user2.ID, &campaignMember)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, user2.ID, campaignMember.UserID)

	// get the new member logged as the member itself
	statusCode, _ = GetCampaignMember(t, user2Cookies, campaign.ID.Bytes, user2.ID, &campaignMember)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, user2.ID, campaignMember.UserID)
}

func TestGetCampaignMembers_Success(t *testing.T) {
	// setup user & campaign
	userGM, userCookies, campaign := campaignMemberSetup(t)
	user2, user2Cookies := CreateTestUser(t)

	// add a new member
	_, err := campaignMembersUseCase.AddCampaignPlayerMember(userGM.ID, domain.CreateCampaignMemberInput{
		CampaignID: campaign.ID.Bytes,
		UserID:     user2.ID,
		Role:       sqlc.MemberRolePlayer,
	})
	require.NoError(t, err)

	// get all members logged as the GM
	var campaignMembers []domain.CampaignMember
	statusCode, _ := GetCampaignMembers(t, userCookies, campaign.ID.Bytes, &campaignMembers)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, 2, len(campaignMembers))
	require.Equal(t, userGM.ID, campaignMembers[0].UserID)
	require.Equal(t, user2.ID, campaignMembers[1].UserID)

	// get all members logged as the member itself
	statusCode, _ = GetCampaignMembers(t, user2Cookies, campaign.ID.Bytes, &campaignMembers)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, 2, len(campaignMembers))
	require.Equal(t, userGM.ID, campaignMembers[0].UserID)
	require.Equal(t, user2.ID, campaignMembers[1].UserID)
}
