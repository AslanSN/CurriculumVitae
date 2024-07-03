package main

import (
	"net/http"

	// "github.com/AslanSN/CurriculumVitae/db"
	handler "github.com/AslanSN/CurriculumVitae/api"
	"github.com/AslanSN/CurriculumVitae/handlers"
	"github.com/AslanSN/CurriculumVitae/helpers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// DB
	// db.DBconnection()

	// Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Vercel connection
	e.GET("/", echo.WrapHandler(http.HandlerFunc(handler.Handler)))

	// Routes
	// registerStaticRoutes(e)
	// registerDBRoutes(e)

	/**
	* ? Unmodulated routes 'cause error:
	* * cmd\main.go:29:2: undefined: registerStaticRoutes
	* * cmd\main.go:30:2: undefined: registerDBRoutes
	 */
	homeHandler := handlers.HomeHandler{}

	data := e.Group("/api/v1/data")

	data.GET(helpers.Api("/info"), homeHandler.HandleShowInfo)
	data.GET(helpers.Api("/aboutMe"), homeHandler.HandleShowAboutMe)
	data.GET(helpers.Api("/experience"), homeHandler.HandleShowExperience)
	data.GET(helpers.Api("/skills"), homeHandler.HandleShowSkills)

	e.Static("/static", "assets")
	e.Static("/icons", "assets/icons")
	e.Static("/images", "assets/images")
	e.Static("/assets/js", "assets/js")

	// Vercel Connection
	http.Handle("/", e)

	// Setup localhost port
	e.Logger.Fatal(e.Start(":2340"))
}
