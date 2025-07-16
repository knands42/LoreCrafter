package domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

type CampaignInvitationInput struct {
	CampaignID uuid.UUID `json:"-"`
	Username   string    `json:"username" example:"johndoe2"`
	InvitedBy  uuid.UUID `json:"-"`
}

func (c *CampaignInvitationInput) Validate() error {
	return nil
}

func (c *CampaignInvitationInput) PrepareToInsert(token string) (sqlc.CreateCampaignInvitationParams, error) {
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

type CampaignInvitation struct {
	sqlc.CampaignInvitation
}

func NewCampaignInvitation(dbData sqlc.CampaignInvitation) *CampaignInvitation {
	return &CampaignInvitation{
		CampaignInvitation: dbData,
	}
}
