package domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

type GetCampaignMemberInput struct {
	CampaignID  uuid.UUID `json:"campaign_id"`
	MemberID    uuid.UUID `json:"member_id"`
	RequesterID uuid.UUID `json:"requester_id"`
}

func NewGetCampaignMemberInput(campaignID, memberID, requesterID uuid.UUID) *GetCampaignMemberInput {
	return &GetCampaignMemberInput{
		CampaignID:  campaignID,
		MemberID:    memberID,
		RequesterID: requesterID,
	}
}

func (c *GetCampaignMemberInput) Validate() error {
	return nil
}

func (c *GetCampaignMemberInput) PrepareToInsert() sqlc.GetCampaignMemberParams {
	return sqlc.GetCampaignMemberParams{
		CampaignID: pgtype.UUID{
			Bytes: c.CampaignID,
			Valid: true,
		},
		MemberID: pgtype.UUID{
			Bytes: c.MemberID,
			Valid: true,
		},
		RequesterID: pgtype.UUID{
			Bytes: c.RequesterID,
			Valid: true,
		},
	}
}

type CampaignMember struct {
	ID             uuid.UUID       `json:"id"`
	CampaignID     uuid.UUID       `json:"campaign_id"`
	UserID         uuid.UUID       `json:"user_id"`
	Role           sqlc.MemberRole `json:"role"`
	JoinedAt       time.Time       `json:"joined_at"`
	LastAccessedAt time.Time       `json:"last_accessed_at"`
}

func NewCampaignMemberFromSqlc(campaignMember sqlc.CampaignMember) CampaignMember {
	return CampaignMember{
		ID:             campaignMember.ID.Bytes,
		CampaignID:     campaignMember.CampaignID.Bytes,
		UserID:         campaignMember.UserID.Bytes,
		Role:           campaignMember.Role,
		JoinedAt:       campaignMember.JoinedAt.Time,
		LastAccessedAt: campaignMember.LastAccessed.Time,
	}
}
