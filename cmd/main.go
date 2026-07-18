package main

import (
	"fmt"
	"net/http"

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

/* main is the entry point of the application where Echo instance is created, middlewares are added, routes are defined, and the server is started on port ":2340".
 */
func main() {
	e := setupRouter()
	// DB
	// db.DBconnection()

	// Vercel Connection
	http.Handle("/", e)

	// Setup localhost port
	e.Logger.Fatal(e.Start(":2340"))
}

func setupRouter() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Custom Middleware to log routes
	e.Use(routeLogger)

	// Routes
	registerStaticRoutes(e)
	registerDBRoutes(e)
	registerPageRoutes(e)

	return e
}

// registerPageRoutes wires each locale's page URL to the Vercel handler, which
// resolves the locale from the path/cookie/Accept-Language. Production routes
// every path here via vercel.json; locally we list them explicitly.
func registerPageRoutes(e *echo.Echo) {
	page := echo.WrapHandler(http.HandlerFunc(handler.Handler))
	e.GET("/", page)
	e.GET("/es", page)
	e.GET("/fr", page)
	e.GET("/sitemap.xml", page)
	e.GET("/robots.txt", page)
}

func registerDBRoutes(e *echo.Echo) {
	homeHandler := handlers.HomeHandler{}

	data := e.Group("/api/v1/data")

	data.GET(helpers.Api("/info"), homeHandler.HandleShowInfo)
	data.GET(helpers.Api("/aboutMe"), homeHandler.HandleShowAboutMe)
	data.GET(helpers.Api("/experience"), homeHandler.HandleShowExperience)
	data.GET(helpers.Api("/skills"), homeHandler.HandleShowSkills)
}

func registerStaticRoutes(e *echo.Echo) {
	// Serve everything under /assets from disk (local dev). In production the
	// Vercel handler serves the same paths from an embedded FS. Reading from
	// disk locally means a recompiled output.css shows up on refresh.
	e.Static("/assets", "assets")
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
