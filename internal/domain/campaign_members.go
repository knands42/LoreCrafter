package domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

type CreateCampaignMemberInput struct {
	CampaignID uuid.UUID       `json:"-"`
	UserID     uuid.UUID       `json:"user_id"`
	Role       sqlc.MemberRole `json:"role" example:"player"`
}

func (c *CreateCampaignMemberInput) Validate() error {
	return nil
}

func (c *CreateCampaignMemberInput) PrepareToInsertPlayer(requesterID uuid.UUID) (sqlc.CreateCampaignPlayerMemberParams, error) {
	newUUIDV7, err := utils.GeneratePGUUID()
	if err != nil {
		return sqlc.CreateCampaignPlayerMemberParams{}, err
	}
	createCampaignMemberParams := sqlc.CreateCampaignPlayerMemberParams{
		ID: newUUIDV7,
		CampaignID: pgtype.UUID{
			Bytes: c.CampaignID,
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: c.UserID,
			Valid: true,
		},
		RequesterID: pgtype.UUID{
			Bytes: requesterID,
			Valid: true,
		},
		Role: sqlc.MemberRolePlayer,
	}

	return createCampaignMemberParams, nil
}

func (c *CreateCampaignMemberInput) PrepareToInsertFirstGM() (sqlc.CreateFirstCampaignMemberParams, error) {
	newUUIDV7, err := utils.GeneratePGUUID()
	if err != nil {
		return sqlc.CreateFirstCampaignMemberParams{}, err
	}
	createCampaignMemberParams := sqlc.CreateFirstCampaignMemberParams{
		ID: newUUIDV7,
		CampaignID: pgtype.UUID{
			Bytes: c.CampaignID,
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: c.UserID,
			Valid: true,
		},
	}

	return createCampaignMemberParams, nil
}

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

type GetCampaignMembersInput struct {
	CampaignID  uuid.UUID `json:"campaign_id"`
	RequesterID uuid.UUID `json:"requester_id"`
}

func NewGetCampaignMembersInput(campaignID, requesterID uuid.UUID) *GetCampaignMemberInput {
	return &GetCampaignMemberInput{
		CampaignID:  campaignID,
		RequesterID: requesterID,
	}
}

func (c *GetCampaignMembersInput) Validate() error {
	return nil
}

func (c *GetCampaignMembersInput) PrepareToInsert() sqlc.GetCampaignMembersParams {
	return sqlc.GetCampaignMembersParams{
		CampaignID: pgtype.UUID{
			Bytes: c.CampaignID,
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
