package models

import (
	"github.com/google/uuid"
)

type Plan struct {
	Id        uuid.UUID `json:"id"`
	Completed bool      `json:"completed"`
	Idea      string    `json:"idea"`
}
