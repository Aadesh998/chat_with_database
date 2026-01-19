package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Aadesh-lab/envloader"
	"github.com/Aadesh-lab/utils"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func LLMCall(schema, userQuery string) (string, error) {
	fmt.Println("LLM Calling Service")

	prompt := utils.GetSQLPrompt(schema, userQuery)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := openai.NewClient(
		option.WithAPIKey(envloader.AppConfig.OpenAIAPIKey),
	)

	response, err := client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model: openai.ChatModelGPT4o,
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(
					"You generate only valid SQL queries. No explanations. No markdown.",
				),
				openai.UserMessage(prompt),
			},
			Temperature: openai.Float(0),
			MaxTokens:   openai.Int(300),
		},
	)

	if err != nil {
		log.Printf("LLM call failed: %v", err)
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", errors.New("LLM returned no choices")
	}

	rawSQL := response.Choices[0].Message.Content
	if rawSQL == "" {
		return "", errors.New("LLM returned empty SQL")
	}

	cleanSQL := utils.CleanSQLResponse(rawSQL)

	fmt.Println("SQL Query received from LLM:")
	fmt.Println(cleanSQL)

	return cleanSQL, nil
}
