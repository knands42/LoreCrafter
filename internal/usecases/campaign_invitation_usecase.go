package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"log"

	"github.com/google/uuid"
	"github.com/knands42/lorecrafter/internal/domain"
	"github.com/knands42/lorecrafter/internal/utils"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

var (
	ErrUserIsAlreadyAMember          = errors.New("user is already a member of this campaign")
	ErrGeneratingTheToken            = errors.New("could not generate a token")
	ErrCreatingTheCampaignInvitation = errors.New("could not create the campaign invitation")

	ErrInviteIsNoLongerValid   = errors.New("the campaign invite is no longer valid")
	ErrAddingTheCampaignMember = errors.New("fail to add the new member to the campaign")
)

type CampaignInvitationUseCase struct {
	ctx                    context.Context
	repo                   sqlc.Querier
	campaignMembersUseCase *CampaignMembersUseCase
}

func NewCampaignInvitationUseCase(
	ctx context.Context,
	repo sqlc.Querier,
	campaignMembersUseCase *CampaignMembersUseCase,
) *CampaignInvitationUseCase {
	return &CampaignInvitationUseCase{
		ctx:                    ctx,
		repo:                   repo,
		campaignMembersUseCase: campaignMembersUseCase,
	}
}

func (c *CampaignInvitationUseCase) CreateAnInvite(input domain.CreateCampaignInvitationInput) (domain.CampaignInvitation, error) {
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

	return *domain.NewCampaignInvitationFromSqlc(createdCampaignInvitation), nil
}

func (c *CampaignInvitationUseCase) ReceiveAnInvite(userId uuid.UUID, input domain.ReceiveCampaignInvite) error {
	err := input.Validate()
	if err != nil {
		return err
	}

	// update & check invitation token
	var params sqlc.UpdateCampaignInviteStatusParams
	if input.Accepted {
		params = input.PrepareToInsert(userId, sqlc.InvitationStatusAccepted)
	} else {
		params = input.PrepareToInsert(userId, sqlc.InvitationStatusRejected)
	}

	updatedInviteStatus, err := c.repo.UpdateCampaignInviteStatus(c.ctx, params)
	if err != nil && err.Error() == "no rows in result set" {
		return ErrInviteIsNoLongerValid
	}

	// don't add the new member to the campaign if not accepted the invite
	if !input.Accepted {
		return err
	}

	// add player as a member of the campaign
	_, err = c.campaignMembersUseCase.AddCampaignPlayerMember(updatedInviteStatus.InvitedBy.Bytes, domain.CreateCampaignMemberInput{
		UserID:     userId,
		CampaignID: updatedInviteStatus.CampaignID.Bytes,
		Role:       sqlc.MemberRolePlayer,
	})

	go func() {
		err = c.repo.InvalidateAllCampaignInvitations(c.ctx, pgtype.UUID{
			Bytes: userId,
			Valid: true,
		})
		if err != nil && err.Error() != "no rows in result set" {
			fmt.Printf("Error invalidating invites for %v user: %v", userId, err)
		}
	}()

	return err
}
