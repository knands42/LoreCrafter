package usecases

import (
	"context"
	"fmt"
	llms2 "github.com/knands42/lorecrafter/internal/adapter/llms"
	"github.com/knands42/lorecrafter/internal/prompts/dnd_5e"
	"strings"

	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
)

// AICampaignUseCase handles AI-powered campaign generation
type AICampaignUseCase struct {
	ctx  context.Context
	repo sqlc.Querier
	llm  llms2.LlmFactoryInterface
}

// NewAICampaignUseCase creates a new instance of AICampaignUseCase
func NewAICampaignUseCase(ctx context.Context, repo sqlc.Querier, llm llms2.LlmFactoryInterface) *AICampaignUseCase {
	return &AICampaignUseCase{
		ctx:  ctx,
		repo: repo,
		llm:  llm,
	}
}

// GenerateCampaignSettings generates campaign settings using AI
func (uc *AICampaignUseCase) GenerateCampaignSettings(input domain.CampaignCreationInput) (domain.CampaignCreationInput, error) {
	// Prepare the prompt with the user's input
	prompt := dnd_5e.NewDND5ECreateCampaignPromptData(input.WorldTheme, input.WrittenTone, input.Setting).CreateCampaignSettingPrompt()

	// Call the OpenAI API
	generatedCampaignSettings, err := uc.llm.GenerateFromSinglePrompt(uc.ctx, prompt)
	if err != nil {
		return input, fmt.Errorf("failed to generate campaign: %w", err)
	}
	input.Setting = generatedCampaignSettings

	return input, nil
}

// extractJSON extracts the first JSON object from a string
func extractJSON(s string) string {
	start := strings.Index(s, "{")
	if start == -1 {
		return ""
	}

	// Find the matching closing brace
	braceCount := 0
	inString := false
	for i := start; i < len(s); i++ {
		c := s[i]
		if c == '"' && (i == 0 || s[i-1] != '\\') {
			inString = !inString
		} else if !inString {
			if c == '{' {
				braceCount++
			} else if c == '}' {
				braceCount--
				if braceCount == 0 {
					return s[start : i+1]
				}
			}
		}
	}

	return ""
}
