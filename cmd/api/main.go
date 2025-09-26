package main

import (
	"log"

	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/danilobml/travel-planner-api/internal/repositories"
	"github.com/danilobml/travel-planner-api/internal/routes"
	"github.com/danilobml/travel-planner-api/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	planRepository := repositories.NewInMemoryPlanRepository()
	planService := services.NewPlanService(planRepository)
	planController := controllers.NewPlanController(planService)

	routes.GetPlannerRouter(r, planController)

	err := r.Run()
	if err != nil {
		log.Panic("Api initialization failed: ", err)
	}
}
