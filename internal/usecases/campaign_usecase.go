package usecases

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
)

var (
	ErrCampaignCreation        = errors.New("error creating campaign")
	ErrCampaignNotFound        = errors.New("campaign not found")
	ErrCampaignMemberCreation  = errors.New("error creating campaign member")
	ErrInsufficientPermissions = errors.New("insufficient permissions")
)

// CampaignUseCase implements the campaign business logic
type CampaignUseCase struct {
	ctx  context.Context
	repo sqlc.Querier
}

// NewCampaignUseCase creates a new campaign use case
func NewCampaignUseCase(
	ctx context.Context,
	repo sqlc.Querier,
) *CampaignUseCase {
	return &CampaignUseCase{
		ctx:  ctx,
		repo: repo,
	}
}

// CreateCampaign creates a new campaign and adds the creator as a GM
func (uc *CampaignUseCase) CreateCampaign(input domain.CampaignCreationInput, creatorID uuid.UUID) (domain.Campaign, error) {
	// Validate the input
	err := input.Validate()
	if err != nil {
		return domain.Campaign{}, err
	}

	// Save the campaign
	createCampaignParams, err := input.ToSqlcParams(creatorID)
	if err != nil {
		return domain.Campaign{}, err
	}
	createdCampaign, err := uc.repo.CreateCampaign(uc.ctx, createCampaignParams)
	if err != nil {
		log.Printf("Error saving campaign: %v", err)
		return domain.Campaign{}, ErrCampaignCreation
	}

	// Add the creator as a GM
	newUUUIDV7, err := utils.GeneratePGUUID()
	if err != nil {
		return domain.Campaign{}, err
	}
	createCampaignMemberParams := sqlc.CreateCampaignMemberParams{
		ID:         newUUUIDV7,
		CampaignID: createdCampaign.ID,
		UserID:     createdCampaign.CreatedBy,
		Role:       sqlc.MemberRoleGm,
	}
	if _, err := uc.repo.CreateCampaignMember(uc.ctx, createCampaignMemberParams); err != nil {
		log.Printf("Error saving campaign member: %v", err)
		return domain.Campaign{}, ErrCampaignMemberCreation
	}

	return domain.FromSqlcCampaignToDomain(createdCampaign), nil
}

// GetCampaign retrieves a campaign by ID if the user has access
func (uc *CampaignUseCase) GetCampaign(input domain.GetCampaignInput) (sqlc.Campaign, error) {
	// Get the campaign
	getCampaignParams, err := input.ToSqlcParams()
	if err != nil {
		return sqlc.Campaign{}, err
	}
	campaign, err := uc.repo.GetCampaignByID(uc.ctx, getCampaignParams)
	if err != nil {
		return sqlc.Campaign{}, ErrCampaignNotFound
	}

	return campaign, nil
}

// UpdateCampaign updates a campaign if the user has GM permissions
func (uc *CampaignUseCase) UpdateCampaign(campaign domain.UpdateCampaignInput) error {
	updateCampaignParams, err := campaign.ToSqlcParams()
	_, err = uc.repo.UpdateCampaign(uc.ctx, updateCampaignParams)
	if err != nil && err.Error() == "no rows in result set" {
		return ErrCampaignNotFound
	}

	return err
}

// DeleteCampaign deletes a campaign if the user has GM permissions
func (uc *CampaignUseCase) DeleteCampaign(input domain.DeleteCampaignInput) error {
	// Delete the campaign
	deleteCampaignParams, err := input.ToSqlcParams()
	if err != nil {
		return err
	}

	return uc.repo.DeleteCampaign(uc.ctx, deleteCampaignParams)
}

// ListUserCampaigns lists all campaigns a user is a member of
func (uc *CampaignUseCase) ListUserCampaigns(userID uuid.UUID) ([]sqlc.Campaign, error) {
	userPGUUID, err := utils.GeneratePGUUIDFromCustomId(userID)
	if err != nil {
		return nil, err
	}

	return uc.repo.ListCampaignsByUserID(uc.ctx, userPGUUID)
}

// AddCampaignMember adds a user to a campaign if the requester has GM permissions
func (uc *CampaignUseCase) AddCampaignMember(campaignID, userID, requesterID uuid.UUID, role string) error {
	return nil
}

// RemoveCampaignMember removes a user from a campaign if the requester has GM permissions
func (uc *CampaignUseCase) RemoveCampaignMember(campaignID, userID, requesterID uuid.UUID) error {
	return nil

}

// LeaveCampaign allows a user to leave a campaign
func (uc *CampaignUseCase) LeaveCampaign(campaignID, userID uuid.UUID) error {
	return nil
}

// GetCampaignMembers lists all members of a campaign if the user has access
func (uc *CampaignUseCase) GetCampaignMembers(campaignID, userID uuid.UUID) ([]sqlc.CampaignMember, error) {
	return []sqlc.CampaignMember{}, nil
}
