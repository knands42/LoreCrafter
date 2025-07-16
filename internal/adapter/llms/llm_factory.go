package llms

import (
	"context"
	"fmt"
	"github.com/knands42/lorecrafter/internal/config"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

type LlmFactory struct {
	llms    []llms.Model
	options []llms.CallOption
}

func NewLlmFactory(
	ctx context.Context,
	cfg config.Config,
	options ...llms.CallOption,
) (*LlmFactory, error) {
	var llmList []llms.Model

	openaiLlm, err := GetOpenApiLlmFactory(cfg)
	if err == nil {
		llmList = append(llmList, openaiLlm)
	}

	geminiLlm, err := GetGeminiLlmFactory(ctx, cfg)
	if err == nil {
		llmList = append(llmList, geminiLlm)
	}

	anthropicLlm, err := GetAnthropicLlmFactory(ctx, cfg)
	if err == nil {
		llmList = append(llmList, anthropicLlm)
	}

	if len(llmList) == 0 {
		return nil, fmt.Errorf("no valid LLMs initialized")
	}

	llmOptions := []llms.CallOption{
		llms.WithMaxTokens(2000),
		llms.WithTemperature(0.8),
	}
	llmOptions = append(llmOptions, options...)

	return &LlmFactory{
		llms:    llmList,
		options: llmOptions,
	}, nil
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

func GetAnthropicLlmFactory(ctx context.Context, cfg config.Config) (llms.Model, error) {
	return anthropic.New(
		anthropic.WithToken(cfg.AntropicAPIKey),
		anthropic.WithModel("claude-sonnet-4-20250514"),
	)
}

func (lf *LlmFactory) GenerateFromSinglePrompt(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	allOptions := append(lf.options, options...)

	var lastErr error
	for _, model := range lf.llms {
		resp, err := llms.GenerateFromSinglePrompt(ctx, model, prompt, allOptions...)
		if err == nil {
			return resp, nil
		}
		lastErr = err
	}
	return "", fmt.Errorf("all LLMs failed: last error: %w", lastErr)
}

func (lf *LlmFactory) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {
	allOptions := append(lf.options, options...)

	var lastErr error
	for _, model := range lf.llms {
		resp, err := model.GenerateContent(ctx, messages, allOptions...)
		if err == nil {
			return resp, nil
		}
		lastErr = err
	}
	return nil, fmt.Errorf("all LLMs failed: last error: %w", lastErr)
}
