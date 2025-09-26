package routes

import (
	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func GetPlannerRouter(router *gin.Engine, pc *controllers.PlanController) *gin.Engine {
	router.GET("/health", heatlthCheck)

	plannerRoutes := router.Group("/plans")

	plannerRoutes.POST("/create", pc.CreateNewPlan)
	plannerRoutes.GET("/", pc.GetAllPlans)
	plannerRoutes.GET("/:id", pc.GetPlanById)

	return router
}

func heatlthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"health-check": "OK!"})
}
