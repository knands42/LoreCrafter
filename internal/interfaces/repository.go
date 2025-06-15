package interfaces

import (
	"context"
	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
)

// CampaignRepository defines the interface for campaign data operations
type CampaignRepository interface {
	CreateCampaign(ctx context.Context, campaign domain.Campaign) error
	GetCampaignByID(ctx context.Context, id uuid.UUID) (*domain.Campaign, error)
	UpdateCampaign(ctx context.Context, campaign domain.Campaign) error
	DeleteCampaign(ctx context.Context, id uuid.UUID) error
	ListCampaignsByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Campaign, error)
}

// CampaignMemberRepository defines the interface for campaign member data operations
type CampaignMemberRepository interface {
	CreateCampaignMember(ctx context.Context, member domain.CampaignMember) error
	GetCampaignMember(ctx context.Context, campaignID, userID uuid.UUID) (*domain.CampaignMember, error)
	UpdateCampaignMember(ctx context.Context, member domain.CampaignMember) error
	DeleteCampaignMember(ctx context.Context, campaignID, userID uuid.UUID) error
	ListCampaignMembers(ctx context.Context, campaignID uuid.UUID) ([]domain.CampaignMember, error)
}
