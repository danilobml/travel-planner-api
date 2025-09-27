package controllers

import (
	"net/http"

	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/danilobml/travel-planner-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PlanController struct {
	service services.PlanService
}

func NewPlanController(service services.PlanService) *PlanController {
	return &PlanController{service: service}
}

func (pc *PlanController) CreateNewPlan(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"response": planResponse})
}

func (pc *PlanController) GetAllPlans(c *gin.Context) {
	plans, _ := pc.service.ListAllPlans()

	c.JSON(http.StatusOK, plans)
}

func (pc *PlanController) GetPlanById(c *gin.Context) {
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
