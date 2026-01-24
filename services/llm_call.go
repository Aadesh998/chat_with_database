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

	client := openai.NewClient(option.WithAPIKey(envloader.AppConfig.OpenAIAPIKey))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	response, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(
				"You are an expert SQL generator. Use the following schema to generate ONLY valid Postgres SQL. " +
					"No explanations, no markdown. Schema:\n" + schema,
			),
			openai.UserMessage(userQuery),
		},
		Temperature: openai.Float(0),
	})

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
