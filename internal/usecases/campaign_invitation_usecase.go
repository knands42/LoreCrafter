package usecases

import (
	"context"
	"errors"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
)

var (
	ErrUserIsAlreadyAMember          = errors.New("user is already a member of this campaign")
	ErrGeneratingTheToken            = errors.New("could not generate a token")
	ErrCreatingTheCampaignInvitation = errors.New("could not create the campaign invitation")
)

type CampaignInvitationUseCase struct {
	ctx  context.Context
	repo sqlc.Querier
}

func NewCampaignInvitationUseCase(
	ctx context.Context,
	repo sqlc.Querier) *CampaignInvitationUseCase {
	return &CampaignInvitationUseCase{
		ctx:  ctx,
		repo: repo,
	}
}

func (c *CampaignInvitationUseCase) CreateAnInvite(input domain.CampaignInvitationInput) (domain.CampaignInvitation, error) {
	if err := input.Validate(); err != nil {
		return domain.CampaignInvitation{}, err
	}

	token, err := utils.GenerateRandomToken(32)
	if err != nil {
		return domain.CampaignInvitation{}, ErrGeneratingTheToken
	}

	params, err := input.PrepareToInsert(token)
	if err != nil {
		log.Printf("error preparing invite %v", err)
		return domain.CampaignInvitation{}, ErrCreatingTheCampaignInvitation
	}

	createdCampaignInvitation, err := c.repo.CreateCampaignInvitation(c.ctx, params)
	if err != nil && err.Error() == "no rows in result set" {
		log.Printf("error creating invite %v", err)
		return domain.CampaignInvitation{}, ErrUserIsAlreadyAMember
	} else if err != nil {
		log.Printf("error creating invite %v", err)
		return domain.CampaignInvitation{}, ErrCreatingTheCampaignInvitation
	}

	return *domain.NewCampaignInvitation(createdCampaignInvitation), nil
}

func (c *CampaignInvitationUseCase) ReceiveAnInvite() {}
