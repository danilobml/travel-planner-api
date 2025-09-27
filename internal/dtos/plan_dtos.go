package dtos

import (
	"github.com/google/uuid"
)

type CreatePlanResponseDto struct {
	Id        uuid.UUID `json:"id"`
	Completed bool      `json:"completed"`
}

type CreatePlanRequestDto struct {
	Place     string   `json:"place"`
	Days      int      `json:"days"`
	Season    string   `json:"season" validate:"required,oneof='winter' 'summer' 'fall' 'spring'"`
	Interests []string `json:"interests"`
	Budget    int      `json:"budget" validate:"required"`
}

type GetPlanResponseDto struct {
	Id         uuid.UUID `json:"id"`
	Completed  bool      `json:"completed"`
	Suggestion string    `json:"suggestion"`
}
