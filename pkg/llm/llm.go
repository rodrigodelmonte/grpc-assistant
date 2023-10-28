package llm

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// Ask asks LLM for a response to a user input.
func Ask(user_input string) (string, error) {

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	chatRequest := openai.ChatCompletionRequest{
		Temperature: 0,
		Model:       os.Getenv("OPENAI_MODEL_NAME"),
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: user_input,
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), chatRequest)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
