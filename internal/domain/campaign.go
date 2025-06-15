package domain

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"strings"
)

// Campaign permission errors
var (
	ErrNotCampaignMember      = errors.New("user is not a member of this campaign")
	ErrInsufficientPermission = errors.New("user does not have sufficient permissions for this action")
)

// CampaignCreationInput represents the input for creating a new campaign
type CampaignCreationInput struct {
	Title          string `json:"title"`
	SettingSummary string `json:"setting_summary"`
	Setting        string `json:"setting"`
	ImageURL       string `json:"image_url"`
	IsPublic       bool   `json:"is_public"`
}

func (campaign *CampaignCreationInput) Validate() error {
	var validationErrors []string

	if strings.TrimSpace(campaign.Title) == "" {
		validationErrors = append(validationErrors, "title is required")
	}

	if len(validationErrors) > 0 {
		return &utils.ValidationError{Errors: validationErrors}
	}

	return nil
}

func (campaign *CampaignCreationInput) ToSqlcParams(creatorID uuid.UUID) (sqlc.CreateCampaignParams, error) {
	newUUUIDV7, err := utils.GeneratePGUUID()
	creatorUUUIDV7, err := utils.GeneratePGUUIDFromCustomId(creatorID)
	if err != nil {
		return sqlc.CreateCampaignParams{}, err
	}

	return sqlc.CreateCampaignParams{
		ID:    newUUUIDV7,
		Title: campaign.Title,
		SettingSummary: pgtype.Text{
			String: campaign.SettingSummary,
			Valid:  true,
		},
		Setting: pgtype.Text{
			String: campaign.Setting,
			Valid:  true,
		},
		IsPublic:  campaign.IsPublic,
		CreatedBy: creatorUUUIDV7,
	}, nil
}

func HasPermission(memberRole sqlc.MemberRole, requiredRole sqlc.MemberRole) bool {
	if memberRole == sqlc.MemberRoleGm {
		return true
	}
	if requiredRole == sqlc.MemberRolePlayer && memberRole == sqlc.MemberRolePlayer {
		return true
	}

	return false
}

type GetCampaignInput struct {
	UserId     uuid.UUID `json:"user_id"`
	CampaignId uuid.UUID `json:"campaign_id"`
}

func (campaign *GetCampaignInput) ToSqlcParams() (sqlc.GetCampaignByIDParams, error) {
	campaignPGUUID, err := utils.GeneratePGUUIDFromCustomId(campaign.CampaignId)
	if err != nil {
		return sqlc.GetCampaignByIDParams{}, err
	}
	userPGUUID, err := utils.GeneratePGUUIDFromCustomId(campaign.UserId)
	if err != nil {
		return sqlc.GetCampaignByIDParams{}, err
	}

	return sqlc.GetCampaignByIDParams{
		UserID: userPGUUID,
		ID:     campaignPGUUID,
	}, nil
}

type UpdateCampaignInput struct {
	ID             uuid.UUID `json:"campaign_id"`
	UserId         uuid.UUID `json:"user_id"`
	Title          string    `json:"title"`
	SettingSummary string    `json:"setting_summary"`
	Setting        string    `json:"setting"`
	ImageURL       string    `json:"image_url"`
	IsPublic       bool      `json:"is_public"`
}

func (campaign *UpdateCampaignInput) ToSqlcParams() (sqlc.UpdateCampaignParams, error) {
	campaignPGUUID, err := utils.GeneratePGUUIDFromCustomId(campaign.ID)
	if err != nil {
		return sqlc.UpdateCampaignParams{}, err
	}
	userPGUUID, err := utils.GeneratePGUUIDFromCustomId(campaign.UserId)
	if err != nil {
		return sqlc.UpdateCampaignParams{}, err
	}

	return sqlc.UpdateCampaignParams{
		ID:     campaignPGUUID,
		UserID: userPGUUID,
		Title:  campaign.Title,
		SettingSummary: pgtype.Text{
			String: campaign.SettingSummary,
			Valid:  true,
		},
		Setting: pgtype.Text{
			String: campaign.Setting,
			Valid:  true,
		},
		ImageUrl: pgtype.Text{
			String: campaign.ImageURL,
			Valid:  true,
		},
		IsPublic: campaign.IsPublic,
	}, nil
}

type DeleteCampaignInput struct {
	ID     uuid.UUID `json:"campaign_id"`
	UserID uuid.UUID `json:"user_id"`
}

func (campaign *DeleteCampaignInput) ToSqlcParams() (sqlc.DeleteCampaignParams, error) {
	campaignPGUUID, err := utils.GeneratePGUUIDFromCustomId(campaign.ID)
	if err != nil {
		return sqlc.DeleteCampaignParams{}, err
	}
	userPGUUID, err := utils.GeneratePGUUIDFromCustomId(campaign.UserID)
	if err != nil {
		return sqlc.DeleteCampaignParams{}, err
	}

	return sqlc.DeleteCampaignParams{
		ID:     campaignPGUUID,
		UserID: userPGUUID,
	}, nil
}
