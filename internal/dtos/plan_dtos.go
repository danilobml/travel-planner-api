package dtos

import "github.com/google/uuid"

type CreatePlanResponseDto struct {
	Id        uuid.UUID `json:"id"`
	Completed bool      `json:"completed"`
}

type CreatePlanRequestDto struct {
	Season    string   `json:"season"`
	Interests []string `json:"interests"`
	Budget    int      `json:"budget"`
}

type GetPlanResponseDto struct {
	Id         uuid.UUID `json:"id"`
	Completed  bool      `json:"completed"`
	Suggestion string    `json:"suggestion"`
}
