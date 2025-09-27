package main

import (
	"context"
	"log"

	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/danilobml/travel-planner-api/internal/db"
	"github.com/danilobml/travel-planner-api/internal/repositories"
	"github.com/danilobml/travel-planner-api/internal/routes"
	"github.com/danilobml/travel-planner-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/openai/openai-go/v2"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initializations:

	// // Open Ai
	// llmClient := openai.NewClient()
	// llmRepository := repositories.NewOpenaiLlmRepository(&llmClient)

	// Langchain/GoogleGemini
	llmClient, err := googleai.New(
		context.Background(),
		googleai.WithDefaultModel("gemini-2.5-flash-lite"),
	)
	if err != nil {
		log.Panic("Llm initialization failed")
	}
	llmRepository := repositories.NewLangchainGoogleLlmRepository(llmClient)

	// In Memory (non-persistent) repo:
	// planRepository := repositories.NewInMemoryPlanRepository()

	// Gorm/Postgres
	DB := db.Init()

	planRepository := repositories.NewPgPlanRepository(DB)
	
	planService := services.NewPlanService(planRepository, llmRepository)
	planController := controllers.NewPlanControllerImplementation(planService)

	// Http server:
	r := gin.Default()

	routes.GetPlannerRouter(r, planController)

	err = r.Run()
	if err != nil {
		log.Panic("Api initialization failed: ", err)
	}
}
