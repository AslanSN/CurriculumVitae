// main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AslanSN/CurriculumVitae/handlers"
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestStaticFiles(t *testing.T) {
	e := setupRouter()

	testCases := []struct {
		url          string
		expectedCode int
	}{
		{"/static", http.StatusOK},
		{"/icons", http.StatusOK},
		{"/images", http.StatusOK},
		{"/js", http.StatusOK},
		{"/css", http.StatusOK},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, tc.url, nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, tc.expectedCode, rec.Code)
	}
}

func TestApiRoutes(t *testing.T) {
	e := echo.New()
	homeHandler := handlers.HomeHandler{}
	data := e.Group("/api/v1/data")
	data.GET(helpers.Api("/info"), homeHandler.HandleShowInfo)
	data.GET(helpers.Api("/aboutMe"), homeHandler.HandleShowAboutMe)
	data.GET(helpers.Api("/experience"), homeHandler.HandleShowExperience)
	data.GET(helpers.Api("/skills"), homeHandler.HandleShowSkills)

	testCases := []struct {
		url          string
		expectedCode int
	}{
		{"/api/v1/data/info", http.StatusOK},
		{"/api/v1/data/aboutMe", http.StatusOK},
		{"/api/v1/data/experience", http.StatusOK},
		{"/api/v1/data/skills", http.StatusOK},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, tc.url, nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, tc.expectedCode, rec.Code)
	}
}
