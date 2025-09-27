package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

type LangchainGoogleLlmRepository struct {
	llmClient *googleai.GoogleAI
}

func NewLangchainGoogleLlmRepository(llmClient *googleai.GoogleAI) *LangchainGoogleLlmRepository {
	return &LangchainGoogleLlmRepository{llmClient: llmClient}
}

func (lr *LangchainGoogleLlmRepository) RequestLlmPlan(req dtos.LlmRequestDto) (*dtos.LlmResponseDto, error) {

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

	answer, err := lr.llmClient.Call(context.TODO(), prompt, llms.WithMaxTokens(3000))
	if err != nil {
		return nil, err
	}

	log.Println(answer)

	return &dtos.LlmResponseDto{
		Response: answer,
	}, nil
}
