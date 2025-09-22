package routes

import (
	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func GetPlannerRouter(router *gin.Engine, pc *controllers.PlanController) *gin.Engine {
	plannerRoutes := router.Group("/plan")

	plannerRoutes.POST("/create", pc.CreateNewPlan)
	plannerRoutes.GET("/", pc.GetAllPlans)
	plannerRoutes.GET("/:id", pc.GetPlanById)

	return router
}
