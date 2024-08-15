package main

import (
	"fmt"
	"net/http"

	// "github.com/AslanSN/CurriculumVitae/internal/db"
	"github.com/AslanSN/CurriculumVitae/helpers"
	"github.com/AslanSN/CurriculumVitae/internal/handlers"
	handler "github.com/AslanSN/CurriculumVitae/vercel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

// main is the entry point of the application where Echo instance is created, middlewares are added, routes are defined, and the server is started on port ":2340".
//
// No parameters.
// No return.
func main() {

	// DB
	// db.DBconnection()

	// Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Vercel connection
	e.GET("", echo.WrapHandler(http.HandlerFunc(handler.Handler)))

	// Routes
	// registerStaticRoutes(e)
	// registerDBRoutes(e)

	// /**
	// * ? Unmodulated routes 'cause error:
	// * * cmd\main.go:29:2: undefined: registerStaticRoutes
	// * * cmd\main.go:30:2: undefined: registerDBRoutes
	//  */
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
	e.Static("/css", "assets/css")

	// Vercel Connection
	http.Handle("/", e)

	// Setup localhost port
	e.Logger.Fatal(e.Start(":2340"))
}

func routeLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		methodColor := getMethodColor(c.Request().Method)
		fmt.Printf("%sMethod: %s%s%s, Route: %s%s%s\n", reset, methodColor, c.Request().Method, reset, blue, c.Request().URL.Path, reset)
		return next(c)
	}
}

func getMethodColor(method string) string {
	switch method {
	case http.MethodGet:
		return green
	case http.MethodPost:
		return yellow
	case http.MethodPut:
		return blue
	case http.MethodDelete:
		return red
	default:
		return white
	}
}
