package controllers

import "github.com/gin-gonic/gin"

type PlanControllerGin interface {
	CreateNewPlan(c *gin.Context)
	GetAllPlans(c *gin.Context)
	GetPlanById(c *gin.Context)
	Revisit(c *gin.Context)
	DeletePlan(c *gin.Context)
}
