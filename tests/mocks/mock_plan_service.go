package mocks

import (
	"errors"

	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/danilobml/travel-planner-api/internal/models"
	"github.com/google/uuid"
)

type MockPlanService struct{}

func (m *MockPlanService) GeneratePlan(req dtos.CreatePlanRequestDto) (dtos.CreatePlanResponseDto, error) {
	return dtos.CreatePlanResponseDto{
		Id:        uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
		Completed: true,
	}, nil
}

func (m *MockPlanService) ListAllPlans() ([]*models.Plan, error) {
	return []*models.Plan{
		{
			Id:         uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			Suggestion: "Munich beer tour",
			Completed:  true,
			Season:     "summer",
		},
		{
			Id:         uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"),
			Suggestion: "Revisited plan suggestion",
			Completed:  false,
			Season:     "winter",
		},
	}, nil
}

func (m *MockPlanService) FindPlanById(id uuid.UUID) (*models.Plan, error) {
	if id.String() == "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa" {
		return &models.Plan{
			Id:         id,
			Suggestion: "Berlin city trip",
			Completed:  true,
			Season:     "spring",
		}, nil
	}
	return nil, errors.New("plan not found")
}

func (m *MockPlanService) GetRevisitedPlanForSeason() (*models.Plan, error) {
	return &models.Plan{
		Id:         uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"),
		Suggestion: "Revisited plan suggestion",
		Completed:  true,
		Season:     "winter",
	}, nil
}
