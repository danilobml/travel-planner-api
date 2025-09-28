package repositories

import (
	"github.com/danilobml/travel-planner-api/internal/models"
	"github.com/google/uuid"
)

type Plan = models.Plan

type PlanRepository interface {
	GetAll() ([]*Plan, error)
	GetById(id uuid.UUID) (*Plan, error)
	Create(p Plan) error
	Delete(id uuid.UUID) error
}
