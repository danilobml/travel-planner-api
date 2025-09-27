package routes

import (
	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/danilobml/travel-planner-api/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func GetPlannerRouter(router *gin.Engine, pc controllers.PlanControllerGin) *gin.Engine {
	corsConfig := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000", "http://t-planner.com"}
	corsConfig.AllowAllOrigins = true

	secureConfig := middleware.DefaultSecureConfig()

	router.Use(secure.New(secureConfig))
	router.Use(cors.New(corsConfig))

	router.GET("/health", heatlthCheck)

	plannerRoutes := router.Group("/plans")

	plannerRoutes.POST("/create", pc.CreateNewPlan)
	plannerRoutes.GET("/", pc.GetAllPlans)
	plannerRoutes.GET("/:id", pc.GetPlanById)
	plannerRoutes.GET("/revisit", pc.Revisit)

	return router
}

func heatlthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"health-check": "OK!"})
}
