package services

import (
	"errors"
	"fmt"
	"math/rand"

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

func (ps *PlanServiceImplementation) GetRevisitedPlanForSeason() (*Plan, error) {
	season := findCurrentSeason()
	
	plans, err := ps.planRepository.GetAll()
	if err != nil {
		return nil, err
	}
	if len(plans) == 0 {
		return nil, errors.New("no plans were yet created")
	}

	var seasonPlans []*Plan
	for _, plan := range plans {
		if (plan.Season == season) {
			seasonPlans = append(seasonPlans, plan)
		}
	}
	if len(seasonPlans) == 0 {
		message := fmt.Sprintf("no plans exist for the curren season (%s)", season)
		return nil, errors.New(message)
	}

	randomPlan := seasonPlans[rand.Intn(len(seasonPlans))]

	return randomPlan, nil
}

func (ps *PlanServiceImplementation) DeletePlan(id uuid.UUID) error {
	err := ps.planRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
