package dtos

import "github.com/google/uuid"

type LlmResponseDto struct {
	Response string
}

type LlmRequestDto struct {
	Id        uuid.UUID
	Place     string
	Days      int
	Season    string
	Interests []string
	Budget    int
}
