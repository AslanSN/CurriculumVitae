package main

import (
	"net/http"

	"github.com/AslanSN/CurriculumVitae/handlers"
	"github.com/AslanSN/CurriculumVitae/helpers"
	handler "github.com/AslanSN/CurriculumVitae/vercel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	homeHandler := handlers.HomeHandler{}
	// homeHandler.CreateHTMLFiles()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", echo.WrapHandler(http.HandlerFunc(handler.Handler)))

	e.GET(helpers.Api("/info"), homeHandler.HandleShowInfo)
	e.GET(helpers.Api("/aboutMe"), homeHandler.HandleShowAboutMe)
	e.GET(helpers.Api("/experience"), homeHandler.HandleShowExperience)
	e.GET(helpers.Api("/skills"), homeHandler.HandleShowSkills)

	e.Static("/static", "assets")
	e.Static("/icons", "assets/icons")
	e.Static("/images", "assets/images")
	e.Static("/assets/js", "assets/js")

	http.Handle("/", e)

	e.Logger.Fatal(e.Start(":2340"))
}
