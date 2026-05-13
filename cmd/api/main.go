package main

import (
	"log"
	"rest-api/internal/database"
	"rest-api/internal/database/repositories"
	"rest-api/internal/handlers"
	"rest-api/internal/middleware"

	"github.com/labstack/echo/v5"
)

const dbUrl = "postgres://postgres:5432@postgres:5432/RestApi?sslmode=disable"

func main() {
	db, err := database.Connect(dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	taskRepo := repositories.NewTaskRepo(db)

	e := echo.New()
	e.Use(middleware.Logging)

	handlers.RegisterRoutes(e, taskRepo)

	e.Start(":8080")
}
