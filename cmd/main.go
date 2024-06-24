package main

import (
	"net/http"

	// "github.com/AslanSN/CurriculumVitae/db"
	// "github.com/AslanSN/CurriculumVitae/db/models"
	handler "github.com/AslanSN/CurriculumVitae/vercel"

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
	registerStaticRoutes(e)
	registerDBRoutes(e)

	// Vercel Connection
	http.Handle("/", e)

	// Setup localhost port
	e.Logger.Fatal(e.Start(":2340"))
}
