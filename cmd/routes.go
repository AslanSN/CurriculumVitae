package main

import (
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/AslanSN/CurriculumVitae/internal/handlers"
	"github.com/labstack/echo/v4"
)

func registerDBRoutes(e *echo.Echo) {
	homeHandler := handlers.HomeHandler{}

	data := e.Group("/api/v1/data")

	data.GET(helpers.Api("/info"), homeHandler.HandleShowInfo)
	data.GET(helpers.Api("/aboutMe"), homeHandler.HandleShowAboutMe)
	data.GET(helpers.Api("/experience"), homeHandler.HandleShowExperience)
	data.GET(helpers.Api("/skills"), homeHandler.HandleShowSkills)
}

func registerStaticRoutes(e *echo.Echo) {
	e.Static("/static", "assets")
	e.Static("/icons", "assets/icons")
	e.Static("/images", "assets/images")
	e.Static("/assets/js", "assets/js")
}
