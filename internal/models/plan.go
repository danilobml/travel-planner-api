package models

import (
	"github.com/google/uuid"
)

type Plan struct {
	Id         uuid.UUID `json:"id" gorm:"primaryKey"`
	Completed  bool      `json:"completed"`
	Suggestion string    `json:"suggestion"`
	Season     string    `json:"-"`
}
