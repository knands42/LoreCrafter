package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

type CreateCampaignInvitationInput struct {
	CampaignID uuid.UUID `json:"-"`
	Username   string    `json:"username" example:"johndoe2"`
	InvitedBy  uuid.UUID `json:"-"`
}

func (c *CreateCampaignInvitationInput) Validate() error {
	var validationErrors []string

	if strings.TrimSpace(c.Username) == "" {
		validationErrors = append(validationErrors, "username is required")
	}

	if len(validationErrors) > 0 {
		return &utils.ValidationError{Errors: validationErrors}
	}

	return nil
}

func (c *CreateCampaignInvitationInput) PrepareToInsert(token string) (sqlc.CreateCampaignInvitationParams, error) {
	id, err := utils.GeneratePGUUID()
	if err != nil {
		return sqlc.CreateCampaignInvitationParams{}, err
	}

	return sqlc.CreateCampaignInvitationParams{
		ID: id,
		CampaignID: pgtype.UUID{
			Bytes: c.CampaignID,
			Valid: true,
		},
		Username: c.Username,
		InvitedBy: pgtype.UUID{
			Bytes: c.InvitedBy,
			Valid: true,
		},
		Token: token,
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(24 * time.Hour),
			Valid: true,
		},
	}, nil
}

type ReceiveCampaignInvite struct {
	Token    string `json:"-"`
	Accepted bool   `json:"accepted"`
}

func (c *ReceiveCampaignInvite) Validate() error {
	return nil
}

func (c *ReceiveCampaignInvite) PrepareToInsert(userId uuid.UUID, status sqlc.InvitationStatus) sqlc.UpdateCampaignInviteStatusParams {
	return sqlc.UpdateCampaignInviteStatusParams{
		UserID: pgtype.UUID{
			Bytes: userId,
			Valid: true,
		},
		Token:  c.Token,
		Status: status,
	}
}

type CampaignInvitation struct {
	ID         uuid.UUID             `json:"id"`
	CampaignID uuid.UUID             `json:"campaign_id"`
	UserID     uuid.UUID             `json:"user_id"`
	InvitedBy  uuid.UUID             `json:"invited_by"`
	Token      string                `json:"token"`
	Status     sqlc.InvitationStatus `json:"status"`
	ExpiresAt  time.Time             `json:"expires_at"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
}

func NewCampaignInvitationFromSqlc(sqlcData sqlc.CampaignInvitation) *CampaignInvitation {
	return &CampaignInvitation{
		ID:         sqlcData.ID.Bytes,
		CampaignID: sqlcData.CampaignID.Bytes,
		UserID:     sqlcData.UserID.Bytes,
		InvitedBy:  sqlcData.InvitedBy.Bytes,
		Token:      sqlcData.Token,
		Status:     sqlcData.Status,
		ExpiresAt:  sqlcData.ExpiresAt.Time,
		CreatedAt:  sqlcData.CreatedAt.Time,
		UpdatedAt:  sqlcData.UpdatedAt.Time,
	}
}
