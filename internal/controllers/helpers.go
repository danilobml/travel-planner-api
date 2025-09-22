package controllers

import "github.com/google/uuid"

func generatePlan(req CreatePlanRequest) CreatePlanResponse {
	uuid := uuid.New()
		
	// TODO - Implement LLM logic using req

	newPlanResponse := CreatePlanResponse{
		Id: uuid,
		Completed: true,
	}

	return newPlanResponse 
}
