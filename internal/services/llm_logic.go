package services

import (
	"log"

	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/google/uuid"
)

func requestLlmPlan(id uuid.UUID, budget int, season string, interests []string) (dtos.LlmResponseDto, error) {
	// TODO - implement actual llm logic

	log.Printf("Generating new plan with id: %s, budget: %d, season: %s, interests: %s", id, budget, season, interests)

	return dtos.LlmResponseDto{
		Response: "Test response",
	}, nil
}
