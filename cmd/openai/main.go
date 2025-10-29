package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func main() {
	_ = godotenv.Load()

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-5",
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String("Say this is a test"),
		},
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp.OutputText())
}