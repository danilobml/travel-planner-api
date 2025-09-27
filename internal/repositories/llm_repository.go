package repositories

import (
	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/google/uuid"
)

type LlmRepository interface {
	RequestLlmPlan(id uuid.UUID, place string, days int, budget int, season string, interests []string) (dtos.LlmResponseDto, error);
}
