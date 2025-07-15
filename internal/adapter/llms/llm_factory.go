package llms

import (
	"context"

	"github.com/tmc/langchaingo/llms"
)

type LlmFactoryInterface interface {
	GenerateFromSinglePrompt(ctx context.Context, prompt string, options ...llms.CallOption) (string, error)
	GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error)
}

type LlmFactory struct {
	llm     llms.Model
	options []llms.CallOption
}

func NewLlmFactory(
	llm llms.Model,
	options ...llms.CallOption,
) *LlmFactory {
	llmOptions := []llms.CallOption{
		llms.WithMaxTokens(2000),
		llms.WithTemperature(0.8),
	}
	llmOptions = append(llmOptions, options...)

	return &LlmFactory{
		llm:     llm,
		options: llmOptions,
	}
}

func (llmFactory *LlmFactory) GenerateFromSinglePrompt(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, llmFactory.llm, prompt, append(llmFactory.options, options...)...)
}

func (llmFactory *LlmFactory) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {
	return llmFactory.llm.GenerateContent(ctx, messages, append(llmFactory.options, options...)...)
}
