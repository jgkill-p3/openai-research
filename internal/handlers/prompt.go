package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type PromptHandler struct {
	client *openai.Client
}

func NewPromptHandler() (*PromptHandler, error) {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY not found in environment")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &PromptHandler{
		client: &client,
	}, nil
}

func (h *PromptHandler) Prompt(ctx context.Context, param openai.ChatCompletionNewParams) (openai.ChatCompletion, error) {
	stream := h.client.Chat.Completions.NewStreaming(ctx, param)

	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)

		if len(chunk.Choices) > 0 {
			print(chunk.Choices[0].Delta.Content)
		}
	}
	println()

	if stream.Err() != nil {
		panic(stream.Err())
	}

	return acc.ChatCompletion, nil
}