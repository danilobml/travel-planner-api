package main

import (
	"log"

	"github.com/danilobml/travel-planner-api/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.GetPlannerRouter(r)

	err := r.Run()
	if err != nil {
		log.Panic("Api initialization failed: ", err)
	}
}
