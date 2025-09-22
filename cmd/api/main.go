package main

import (
	"log"

	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/danilobml/travel-planner-api/internal/repositories"
	"github.com/danilobml/travel-planner-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	planRepository := repositories.NewInMemoryPlanRepository()
	planController := controllers.NewPlanController(planRepository)

	routes.GetPlannerRouter(r, planController)

	err := r.Run()
	if err != nil {
		log.Panic("Api initialization failed: ", err)
	}
}
