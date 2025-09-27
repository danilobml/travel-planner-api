package repositories

import (
	"github.com/danilobml/travel-planner-api/internal/dtos"
)

type LlmRepository interface {
	RequestLlmPlan(req dtos.LlmRequestDto) (*dtos.LlmResponseDto, error)
}
