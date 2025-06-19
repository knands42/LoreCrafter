package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

// Campaign permission errors
var (
	ErrNotCampaignMember      = errors.New("user is not a member of this campaign")
	ErrInsufficientPermission = errors.New("user does not have sufficient permissions for this action")
)

// CampaignCreationInput represents the input for creating a new campaign
type CampaignCreationInput struct {
	Title           string              `json:"title"`
	SettingSummary  string              `json:"setting_summary"`
	Setting         string              `json:"setting"`
	GameSystem      sqlc.GameSystemEnum `json:"game_system"`
	NumberOfPlayers int16               `json:"number_of_players"`
	ImageURL        string              `json:"image_url"`
	IsPublic        bool                `json:"is_public"`
}

func (campaign *CampaignCreationInput) Validate() error {
	var validationErrors []string

	if strings.TrimSpace(campaign.Title) == "" {
		validationErrors = append(validationErrors, "title is required")
	}

	if len(campaign.Title) > 100 {
		validationErrors = append(validationErrors, "title must be at most 100 characters")
	}

	if len(validationErrors) > 0 {
		return &utils.ValidationError{Errors: validationErrors}
	}

	return nil
}

func (campaign *CampaignCreationInput) PrepareToInsert(creatorID uuid.UUID) (sqlc.CreateCampaignParams, error) {
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
		NumberOfPlayers: pgtype.Int2{
			Int16: campaign.NumberOfPlayers,
			Valid: true,
		},
		GameSystem: campaign.GameSystem,
		Setting: pgtype.Text{
			String: campaign.Setting,
			Valid:  true,
		},
		Status:    sqlc.CampaignStatusEnumPLANNING,
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

type Campaign struct {
	ID              uuid.UUID               `json:"id"`
	Title           string                  `json:"title"`
	SettingSummary  string                  `json:"setting_summary"`
	Setting         string                  `json:"setting"`
	GameSystem      sqlc.GameSystemEnum     `json:"game_system"`
	NumberOfPlayers uint16                  `json:"number_of_players"`
	ImageUrl        string                  `json:"image_url"`
	IsPublic        bool                    `json:"is_public"`
	InviteCode      string                  `json:"invite_code"`
	Metadata        interface{}             `json:"metadata"`
	Status          sqlc.CampaignStatusEnum `json:"status"`
	CreatedBy       uuid.UUID               `json:"created_by"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
}

func FromSqlcCampaignToDomain(campaignSqlc sqlc.Campaign) Campaign {
	return Campaign{
		ID:              campaignSqlc.ID.Bytes,
		Title:           campaignSqlc.Title,
		SettingSummary:  campaignSqlc.SettingSummary.String,
		Setting:         campaignSqlc.Setting.String,
		GameSystem:      campaignSqlc.GameSystem,
		NumberOfPlayers: uint16(campaignSqlc.NumberOfPlayers.Int16),
		ImageUrl:        campaignSqlc.ImageUrl.String,
		IsPublic:        campaignSqlc.IsPublic,
		InviteCode:      campaignSqlc.InviteCode.String,
		Status:          campaignSqlc.Status,
		CreatedBy:       campaignSqlc.CreatedBy.Bytes,
		CreatedAt:       campaignSqlc.CreatedAt.Time,
		UpdatedAt:       campaignSqlc.UpdatedAt.Time,
	}
}
