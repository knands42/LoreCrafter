package usecases

import (
	"context"

	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

// AICampaignUseCaseImpl implements AICampaignUseCase
type AICampaignUseCase struct {
	ctx  context.Context
	repo sqlc.Querier
}

// NewAICampaignUseCase creates a new instance of AICampaignUseCase
func NewAICampaignUseCase(ctx context.Context, repo sqlc.Querier) *AICampaignUseCase {
	return &AICampaignUseCase{ctx: ctx, repo: repo}
}

// GenerateCampaignSettings generates campaign settings using AI
func (uc *AICampaignUseCase) GenerateCampaignSettings(input domain.CampaignCreationInput) (domain.CampaignCreationInput, error) {
	return domain.CampaignCreationInput{}, nil
}
