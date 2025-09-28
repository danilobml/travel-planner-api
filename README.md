# Travel Planner API

A RESTful API built with Go and Gin to generate, store, and revisit travel plans, using LLMs. This service demonstrates clean architecture with controllers, services, repositories, DTOs, routes, and tests.

## Table of Contents
- Features
- Tech Stack
- Project Structure
- Getting Started
  - Prerequisites
  - Environment Variables
  - Run Locally
  - Run with Docker Compose
- Makefile Workflow
- Testing
- Configuration
- API Reference
  - Health Check
  - Create New Plan
  - Get All Plans
  - Get One Plan
  - Revisit Plan

## Features
- Create a new travel plan with place, days, budget, season, and interests
- Retrieve all stored plans
- Retrieve a plan by ID
- Revisit a previous plan suggestion
- JSON response DTOs with validation
- Unit and integration tests with testify and httptest

## Tech Stack
- Go (Gin framework)
- Gorm ORM
- Langchain for LLM calls
- Testify for testing
- Playground validator
- Air for hot reload in dev mode
- PostgresDb in Docker container

## Project Structure
- travel-planner-api/
  - cmd/api/ (main entrypoint)
  - internal/
    - controllers/ (Gin handlers)
    - services/ (business logic)
    - repositories/ (persistence of plans and LLM calls)
    - dtos/ (request/response DTOs)
    - models/ (domain models)
    - routes/ (route registration)
    - middleware
  - tests/
    - mocks/ (mock service)
    - *.go (integration tests)

## Getting Started

### Prerequisites
- Go 1.22+
- Docker and docker-compose (for Postgres DB)

### Environment Variables
Create a .env (or export in your shell) and adjust values as needed:

PORT=8080
ENVIRONMENT=development
POSTGRES_URL="postgres://pg:pass@localhost:5432/plans"
OPENAI_API_KEY=abc...
GOOGLE_API_KEY=abc...

(Only the api key from the chosen model is required)

## Choosing LLM api and model
at ./cmd/api/main.go

By default Gemini 2.5 pro is selected. Uncomment the OpenAi section (and comment out the Gemini one), to select it. You can also change the model (refer to `www.openai.com` or `www.gemini.com`):

```go
func main() {
    ...

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
    
    ...
}
```

### Run Locally
1) Clone the repo: git clone https://github.com/danilobml/travel-planner-api.git
2) Change directory: cd travel-planner-api
3) run in dev: `make dev`
4) or build and run in production mode `make run`

## Makefile Workflow
Targets (example):
- dev: starts docker services, waits for DB, runs air for live reload
- up: docker-compose up -d database
- down: docker-compose down
- wait-db: loop until postgres is ready using pg_isready
- build: go build ./cmd/api -o ./api
- run: sets ENVIRONMENT=production and runs the API
- test: go test ./... -v

## Testing
- Run all tests:
```bash
make test
```

## API Reference

### Health Check
- Method and path: GET /health
- Success response body: { "ok": true }

### Create New Plan
- Method and path: POST /plans/create
- Required header: Content-Type: application/json
- Example body fields: place (string), days (number), budget (number), season (string), interests (array of strings)
- Example success response body fields: response.id (uuid), response.completed (bool)

- Possible errors:
  - 400 Bad Request (invalid JSON or validation error)
  - 415 Unsupported Media Type (wrong Content-Type)

### Get All Plans
- Method and path: GET /plans
- Example success response body: array of plan objects with id, suggestion, completed

### Get One Plan
- Method and path: GET /plans/{id}
- Example success response body: plan object with id, suggestion, completed
- Possible errors:
  - 400 Bad Request (invalid UUID format)
  - 404 Not Found (plan does not exist)

### Revisit Plan
- Method and path: GET /plans/revisit
- Example success response body: plan object with id, suggestion, completed

