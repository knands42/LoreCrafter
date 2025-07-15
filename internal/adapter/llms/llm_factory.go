package llms

import (
	"context"
	"github.com/knands42/lorecrafter/internal/config"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"

	"github.com/tmc/langchaingo/llms"
)

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

func GetGeminiLlmFactory(ctx context.Context, cfg config.Config) (llms.Model, error) {
	return googleai.New(
		ctx,
		googleai.WithAPIKey(cfg.GoogleAPIKey),
		googleai.WithDefaultModel("gemini-2.0-flash-exp"),
	)
}

func GetOpenApiLlmFactory(cfg config.Config) (llms.Model, error) {
	return openai.New(
		openai.WithToken(cfg.OpenAIAPIKey),
		openai.WithModel("gpt-4-turbo-preview"),
	)
}

func (llmFactory *LlmFactory) GenerateFromSinglePrompt(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, llmFactory.llm, prompt, append(llmFactory.options, options...)...)
}

func (llmFactory *LlmFactory) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {
	return llmFactory.llm.GenerateContent(ctx, messages, append(llmFactory.options, options...)...)
}
