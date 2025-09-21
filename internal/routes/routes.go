package routes

import (
	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func GetPlannerRouter(router *gin.Engine) *gin.Engine {
	plannerRoutes := router.Group("/plan")

	plannerRoutes.POST("/create", controllers.CreatePlan)

	return router
}
