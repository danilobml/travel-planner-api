package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

type LangchainGoogleLlmRepository struct {
	llmClient *googleai.GoogleAI
}

func NewLangchainGoogleLlmRepository(llmClient *googleai.GoogleAI) *LangchainGoogleLlmRepository {
	return &LangchainGoogleLlmRepository{llmClient: llmClient}
}

func (lr *LangchainGoogleLlmRepository) RequestLlmPlan(id uuid.UUID, place string, days int, budget int, season string, interests []string) (dtos.LlmResponseDto, error) {

	log.Printf("Generating new plan with id: %s.", id)

	placeText := "Suggest a location"
	if place != "" {
		placeText = fmt.Sprintf("It must be in %s", place)
	}

	daysText := fmt.Sprintf("%d days", days)
	if days == 0 {
		daysText = "an open-ended number of days"
	}

	prompt := fmt.Sprintf(
		"Generate a travel plan suggestion for the %s, with a budget of $%d, for %s. %s. The focus should be on: %v",
		season, budget, daysText, placeText, interests,
	)

	answer, err := lr.llmClient.Call(context.TODO(), prompt, llms.WithMaxTokens(3000))
	if err != nil {
		return dtos.LlmResponseDto{}, err
	}

	log.Println(answer)

	return dtos.LlmResponseDto{
		Response: answer,
	}, nil
}
