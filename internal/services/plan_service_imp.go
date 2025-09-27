package services

import (
	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/danilobml/travel-planner-api/internal/repositories"
	"github.com/google/uuid"
)

type PlanServiceImplementation struct {
	planRepository repositories.PlanRepository
	llmRepository  repositories.LlmRepository
}

func NewPlanService(planRepository repositories.PlanRepository, llmRepository repositories.LlmRepository) *PlanServiceImplementation {
	return &PlanServiceImplementation{planRepository: planRepository, llmRepository: llmRepository}
}

func (ps *PlanServiceImplementation) GeneratePlan(req dtos.CreatePlanRequestDto) (dtos.CreatePlanResponseDto, error) {
	uuid := uuid.New()

	llmrequest := dtos.LlmRequestDto{
		Id: uuid,
		Place: req.Place,
		Days: req.Days,
		Season: req.Season,
		Interests: req.Interests,
		Budget: req.Budget,
	}

	llmResponse, err := ps.llmRepository.RequestLlmPlan(llmrequest)
	if err != nil {
		return dtos.CreatePlanResponseDto{
			Id:        uuid,
			Completed: false,
		}, err
	}

	var plan Plan
	plan.Id = uuid
	plan.Season = req.Season
	plan.Suggestion = llmResponse.Response
	plan.Completed = true

	ps.planRepository.Create(plan)

	return dtos.CreatePlanResponseDto{
		Id:        uuid,
		Completed: true,
	}, nil
}

func (ps *PlanServiceImplementation) ListAllPlans() ([]*Plan, error) {
	plans, err := ps.planRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return plans, nil
}

func (ps *PlanServiceImplementation) FindPlanById(id uuid.UUID) (*Plan, error) {
	plan, err := ps.planRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
