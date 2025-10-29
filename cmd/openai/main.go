package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/jgkill-p3/openai-research/internal/handlers"
	"github.com/openai/openai-go/v3"
)

func main() {
	handler, err := handlers.NewPromptHandler()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	param := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT5Nano2025_08_07,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("How many 'R's are in Strawberry?"),
		},
		Seed: openai.Int(0),
	}

	completion, err := handler.Prompt(ctx, param)
	if err != nil {
		log.Fatal(err)
	}

	param.Messages = append(param.Messages, completion.Choices[0].Message.ToParam())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	param.Messages = append(param.Messages, openai.UserMessage(input))

	completion, err = handler.Prompt(ctx, param)
	if err != nil {
		log.Fatal(err)
	}
}