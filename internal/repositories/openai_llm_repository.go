package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/openai/openai-go/v2"
)

type OpenaiLlmRepository struct {
	llmClient *openai.Client
}

func NewOpenaiLlmRepository(llmClient *openai.Client) *OpenaiLlmRepository {
	return &OpenaiLlmRepository{llmClient: llmClient}
}

func (lr *OpenaiLlmRepository) RequestLlmPlan(req dtos.LlmRequestDto) (*dtos.LlmResponseDto, error) {

	log.Printf("Generating new plan with id: %s.", req.Id)

	placeText := "Suggest a location"
	if req.Place != "" {
		placeText = fmt.Sprintf("It must be in %s", req.Place)
	}

	daysText := fmt.Sprintf("%d days", req.Days)
	if req.Days == 0 {
		daysText = "an open-ended number of days"
	}

	prompt := fmt.Sprintf(
		"Generate a travel plan suggestion for the %s, with a budget of $%d, for %s. %s. The focus should be on: %v",
		req.Season, req.Budget, daysText, placeText, req.Interests,
	)

	chatCompletion, err := lr.llmClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		return nil, err
	}

	llmResponse := chatCompletion.Choices[0].Message.Content

	return &dtos.LlmResponseDto{
		Response: llmResponse,
	}, nil
}
