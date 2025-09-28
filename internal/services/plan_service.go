package services

import (
	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/danilobml/travel-planner-api/internal/models"
	"github.com/google/uuid"
)

type Plan = models.Plan

type PlanService interface {
	GeneratePlan(req dtos.CreatePlanRequestDto) (dtos.CreatePlanResponseDto, error)
	ListAllPlans() ([]*Plan, error)
	FindPlanById(id uuid.UUID) (*Plan, error)
	GetRevisitedPlanForSeason() (*Plan, error)
	DeletePlan(id uuid.UUID) error
}
