package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

var (
	ErrFailedToRetrieveAMember = errors.New("failed to retrieve a member")
)

type CampaignMembersUseCase struct {
	ctx  context.Context
	repo sqlc.Querier
}

func NewCampaignMembersUseCase(ctx context.Context, repo sqlc.Querier) *CampaignMembersUseCase {
	return &CampaignMembersUseCase{
		ctx:  ctx,
		repo: repo,
	}
}

// AddCampaignMember adds a user to a campaign if the requester has GM permissions
func (uc *CampaignMembersUseCase) AddCampaignMember(campaignID, userID, requesterID uuid.UUID, role string) error {
	return nil
}

// RemoveCampaignMember removes a user from a campaign if the requester has GM permissions
func (uc *CampaignMembersUseCase) RemoveCampaignMember(campaignID, userID, requesterID uuid.UUID) error {
	return nil

}

// LeaveCampaign allows a user to leave a campaign
func (uc *CampaignMembersUseCase) LeaveCampaign(campaignID, userID uuid.UUID) error {
	return nil
}

// GetCampaignMembers lists all members of a campaign if the user has access
func (uc *CampaignMembersUseCase) GetCampaignMembers(campaignID, userID uuid.UUID) ([]sqlc.CampaignMember, error) {
	return []sqlc.CampaignMember{}, nil
}

// GetCampaignMember lists all members of a campaign if the user has access
func (uc *CampaignMembersUseCase) GetCampaignMember(input domain.GetCampaignMemberInput) (domain.CampaignMember, error) {
	// validate
	err := input.Validate()
	if err != nil {
		return domain.CampaignMember{}, err
	}

	// persist
	params := input.PrepareToInsert()
	campaignMember, err := uc.repo.GetCampaignMember(uc.ctx, params)
	if err != nil {
		return domain.CampaignMember{}, ErrFailedToRetrieveAMember
	}

	return domain.NewCampaignMemberFromSqlc(campaignMember), nil
}
