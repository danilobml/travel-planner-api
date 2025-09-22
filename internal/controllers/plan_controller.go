package controllers

import (
	"net/http"

	"github.com/danilobml/travel-planner-api/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PlanController struct {
	repository repositories.PlanRepository
}

type CreatePlanRequest struct {
	Season    string   `json:"season"`
	Interests []string `json:"interests"`
	Budget    int      `json:"budget"`
}

type CreatePlanResponse struct {
	Id        uuid.UUID `json:"id"`
	Completed bool      `json:"completed"`
}

type GetPlanResponse struct {
	Id         uuid.UUID `json:"id"`
	Completed  bool      `json:"completed"`
	Suggestion string    `json:"suggestion"`
}

func NewPlanController(repository repositories.PlanRepository) *PlanController {
	return &PlanController{repository: repository}
}

func (pc *PlanController) CreateNewPlan(c *gin.Context) {
	var req CreatePlanRequest

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, generatePlan(req))
}

func (pc *PlanController) GetAllPlans(c *gin.Context) {
	plans, _ := pc.repository.GetAll()

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

	plan, err := pc.repository.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Plan not found",
		})
		return
	}

	planResponse := GetPlanResponse{
		Id:         plan.Id,
		Completed:  plan.Completed,
		Suggestion: plan.Suggestion,
	}

	c.JSON(http.StatusOK, planResponse)
}
