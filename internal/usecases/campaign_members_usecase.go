package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

var (
	ErrFailedToCreateCampaignMember = errors.New("failed to create campaign member")
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

// AddCampaignPlayerMember adds a user to a campaign if the requester has GM permissions
func (uc *CampaignMembersUseCase) AddCampaignPlayerMember(requester uuid.UUID, input domain.CreateCampaignMemberInput) (domain.CampaignMember, error) {
	// validate the input
	err := input.Validate()
	if err != nil {
		return domain.CampaignMember{}, err
	}

	// create the user member
	params, err := input.PrepareToInsertPlayer(requester)
	if err != nil {
		log.Printf("failed to prepare to insert a campaign member: %v", err)
		return domain.CampaignMember{}, ErrFailedToCreateCampaignMember
	}

	createdCampaignMember, err := uc.repo.CreateCampaignPlayerMember(uc.ctx, params)
	if err != nil {
		log.Printf("Error saving campaign member: %v", err)
		return domain.CampaignMember{}, ErrAddingTheCampaignMember
	}

	return domain.NewCampaignMemberFromSqlc(createdCampaignMember), err
}

// CreateGMMember adds the GM as the first member to a new campaign
func (uc *CampaignMembersUseCase) CreateGMMember(input domain.CreateCampaignMemberInput) (domain.CampaignMember, error) {
	// validate the input
	err := input.Validate()
	if err != nil {
		return domain.CampaignMember{}, err
	}

	// create the user member
	params, err := input.PrepareToInsertFirstGM()
	if err != nil {
		log.Printf("failed to prepare to insert a campaign member: %v", err)
		return domain.CampaignMember{}, ErrFailedToCreateCampaignMember
	}

	createdCampaignMember, err := uc.repo.CreateFirstCampaignMember(uc.ctx, params)
	if err != nil {
		log.Printf("Error saving campaign member: %v", err)
		return domain.CampaignMember{}, ErrAddingTheCampaignMember
	}

	return domain.NewCampaignMemberFromSqlc(createdCampaignMember), err
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
func (uc *CampaignMembersUseCase) GetCampaignMembers(input domain.GetCampaignMembersInput) ([]domain.CampaignMember, error) {
	// validate
	err := input.Validate()
	if err != nil {
		return []domain.CampaignMember{}, err
	}

	// persist
	params := input.PrepareToInsert()
	campaignMembers, err := uc.repo.GetCampaignMembers(uc.ctx, params)
	if err != nil {
		return []domain.CampaignMember{}, ErrFailedToRetrieveAMember
	}

	members := make([]domain.CampaignMember, len(campaignMembers))
	for i, member := range campaignMembers {
		members[i] = domain.NewCampaignMemberFromSqlc(member)
	}
	return members, nil
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
