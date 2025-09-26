package dtos

import "github.com/google/uuid"

type CreatePlanResponseDto struct {
	Id        uuid.UUID
	Completed bool
}

type CreatePlanRequestDto struct {
	Season    string
	Interests []string
	Budget    int
}


type GetPlanResponseDto struct {
	Id         uuid.UUID `json:"id"`
	Completed  bool      `json:"completed"`
	Suggestion string    `json:"suggestion"`
}
