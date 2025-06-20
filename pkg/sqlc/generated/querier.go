// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateCampaign(ctx context.Context, arg CreateCampaignParams) (Campaign, error)
	CreateCampaignMember(ctx context.Context, arg CreateCampaignMemberParams) (CampaignMember, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCampaign(ctx context.Context, arg DeleteCampaignParams) error
	DeleteCampaignMember(ctx context.Context, arg DeleteCampaignMemberParams) error
	GenerateInviteCode(ctx context.Context, arg GenerateInviteCodeParams) (Campaign, error)
	GetCampaignByID(ctx context.Context, arg GetCampaignByIDParams) (Campaign, error)
	GetCampaignByInviteCode(ctx context.Context, inviteCode pgtype.Text) (Campaign, error)
	GetCampaignMember(ctx context.Context, arg GetCampaignMemberParams) (CampaignMember, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByUsernameOrEmail(ctx context.Context, arg GetUserByUsernameOrEmailParams) (User, error)
	ListCampaignMembers(ctx context.Context, campaignID pgtype.UUID) ([]CampaignMember, error)
	ListCampaignsByUserID(ctx context.Context, userID pgtype.UUID) ([]Campaign, error)
	UpdateCampaign(ctx context.Context, arg UpdateCampaignParams) (Campaign, error)
	UpdateCampaignMember(ctx context.Context, arg UpdateCampaignMemberParams) (CampaignMember, error)
}

var _ Querier = (*Queries)(nil)
