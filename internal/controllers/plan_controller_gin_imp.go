package controllers

import (
	"net/http"

	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/danilobml/travel-planner-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PlanControllerImplementation struct {
	service services.PlanService
}

func NewPlanControllerImplementation(service services.PlanService) *PlanControllerImplementation {
	return &PlanControllerImplementation{service: service}
}

func (pc *PlanControllerImplementation) CreateNewPlan(c *gin.Context) {
	var req dtos.CreatePlanRequestDto

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"errors":  errors.Error(),
		})
		return
	}

	planResponse, err := pc.service.GeneratePlan(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": planResponse, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"response": planResponse})
}

func (pc *PlanControllerImplementation) GetAllPlans(c *gin.Context) {
	plans, _ := pc.service.ListAllPlans()

	c.JSON(http.StatusOK, plans)
}

func (pc *PlanControllerImplementation) GetPlanById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	plan, err := pc.service.FindPlanById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Plan not found",
		})
		return
	}

	planResponse := dtos.GetPlanResponseDto{
		Id:         plan.Id,
		Completed:  plan.Completed,
		Suggestion: plan.Suggestion,
	}

	c.JSON(http.StatusOK, planResponse)
}

func (pc *PlanControllerImplementation) Revisit(c *gin.Context) {
	revisitedPlan, err := pc.service.GetRevisitedPlanForSeason()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, revisitedPlan)
}
