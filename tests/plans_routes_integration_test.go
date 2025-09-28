package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danilobml/travel-planner-api/internal/controllers"
	"github.com/danilobml/travel-planner-api/internal/dtos"
	"github.com/danilobml/travel-planner-api/mocks"
	"github.com/danilobml/travel-planner-api/internal/models"
	"github.com/danilobml/travel-planner-api/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var mockService = &mocks.MockPlanService{}

func setupServer() *httptest.Server {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	pc := controllers.NewPlanControllerGinImplementation(mockService)
	router := routes.GetPlannerRouter(r, pc)

	return httptest.NewTLSServer(router)
}

func Test_PlanRoutes_HealthCheck(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	res, err := client.Get(srv.URL + "/health")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_PlanRoutes_CreateNewPlan(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	payload := map[string]any{
		"place":     "Berlin",
		"days":      3,
		"budget":    200,
		"season":    "summer",
		"interests": []string{"history", "food"},
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(http.MethodPost, srv.URL+"/plans/create", bytes.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, "201 Created", res.Status)

	var resp struct {
		Response dtos.CreatePlanResponseDto `json:"response"`
	}
	require.NoError(t, json.NewDecoder(res.Body).Decode(&resp))

	require.True(t, resp.Response.Completed)
	require.NotEmpty(t, resp.Response.Id)
}

func Test_PlanRoutes_CreateNewPlan_InvalidJSON(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	req, err := http.NewRequest(http.MethodPost, srv.URL+"/plans/create", bytes.NewBuffer([]byte("{invalid")))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func Test_PlanRoutes_CreateNewPlan_EmptyBody(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	req, err := http.NewRequest(http.MethodPost, srv.URL+"/plans/create", bytes.NewBuffer(nil))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func Test_PlanRoutes_GetAllPlans(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	res, err := client.Get(srv.URL + "/plans")
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, "200 OK", res.Status)

	var got []models.Plan
	require.NoError(t, json.NewDecoder(res.Body).Decode(&got))

	require.Len(t, got, 2)

	require.Equal(t, "Munich beer tour", got[0].Suggestion)
	require.Equal(t, true, got[0].Completed)

	require.Equal(t, "Revisited plan suggestion", got[1].Suggestion)
	require.Equal(t, false, got[1].Completed)

	// Omitted in json
	require.Equal(t, "", got[0].Season)
}

func Test_PlanRoutes_GetOnePlan(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	res, err := client.Get(srv.URL + "/plans/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	require.NoError(t, err)
	require.Equal(t, "200 OK", res.Status)

	var resp models.Plan
	require.NoError(t, json.NewDecoder(res.Body).Decode(&resp))

	require.Equal(t, "", resp.Season)
	require.Equal(t, "Berlin city trip", resp.Suggestion)
	require.Equal(t, true, resp.Completed)
}

func Test_PlanRoutes_GetOnePlan_InvalidUUID(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	res, err := client.Get(srv.URL + "/plans/not-a-uuid")
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func Test_PlanRoutes_GetOnePlan_NotFound(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	res, err := client.Get(srv.URL + "/plans/ffffffff-ffff-ffff-ffff-ffffffffffff")
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_PlanRoutes_Revisit(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	res, err := client.Get(srv.URL + "/plans/revisit")
	require.NoError(t, err)
	require.Equal(t, "200 OK", res.Status)

	var resp models.Plan
	require.NoError(t, json.NewDecoder(res.Body).Decode(&resp))

	require.Equal(t, "", resp.Season)
	require.Equal(t, "Revisited plan suggestion", resp.Suggestion)
	require.Equal(t, true, resp.Completed)
}

func Test_PlanRoutes_DeletePlan(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	req, err := http.NewRequest("DELETE", srv.URL + "/plans/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", nil)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, "204 No Content", res.Status)
}

func Test_PlanRoutes_DeletePlan_WrongId(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	client := srv.Client()

	req, err := http.NewRequest("DELETE", srv.URL + "/plans/cccccccc-cccc-cccc-cccc-cccccccccccc", nil)
	require.NoError(t, err)

	res, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, res.StatusCode)
}
