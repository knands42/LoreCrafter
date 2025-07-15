package interfaces

import (
	"context"
	"github.com/tmc/langchaingo/llms"
)

type LlmFactoryInterface interface {
	GenerateFromSinglePrompt(ctx context.Context, prompt string, options ...llms.CallOption) (string, error)
	GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error)
}
