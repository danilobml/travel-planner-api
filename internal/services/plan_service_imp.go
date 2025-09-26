package services

import (
	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/danilobml/travel-planner-api/internal/repositories"
	"github.com/google/uuid"
)

type PlanServiceImplementation struct {
	repository repositories.PlanRepository
}

func NewPlanService(repository repositories.PlanRepository) *PlanServiceImplementation {
	return &PlanServiceImplementation{repository: repository}
}

func (ps *PlanServiceImplementation) GeneratePlan(req dtos.CreatePlanRequestDto) dtos.CreatePlanResponseDto {
	uuid := uuid.New()

	llmResponse, err := requestLlmPlan(uuid, req.Budget, req.Season, req.Interests)
	if err != nil {
		return dtos.CreatePlanResponseDto{
			Id:        uuid,
			Completed: false,
		}
	}

	var plan Plan
	plan.Id = uuid
	plan.Suggestion = llmResponse.Response
	plan.Completed = true

	ps.repository.Create(plan)

	return dtos.CreatePlanResponseDto{
		Id:        uuid,
		Completed: true,
	}
}

func (ps *PlanServiceImplementation) ListAllPlans() ([]*Plan, error) {
	plans, err := ps.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (ps *PlanServiceImplementation) FindPlanById(id uuid.UUID) (*Plan, error) {
	plan, err := ps.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
