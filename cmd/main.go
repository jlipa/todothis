package main

import (
	"github.com/jlipa/todothis/config"
	"github.com/jlipa/todothis/internal/adapters"
	"github.com/jlipa/todothis/internal/services"
	"github.com/labstack/echo"
)

var service services.Service

func main() {
	// Configs and DB
	config.ReadConfig()
	db := config.NewDB()
	db.LogMode(true)
	defer db.Close()

	repo := adapters.NewRepo(db)
	service = services.NewService(&repo)

	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", welcome)

	e.GET("/todos", TodoGetAll)  // Get All TODO
	e.GET("/todos/:id", TodoGet) // Get specific TODO and tasks list

	e.POST("/todos/new", TodoNew)                             // Create TODO
	e.POST("/todos/:id/add-task", TodoAddTask)                // Add TODO task
	e.POST("/todos/:id/edit-task/:task_id", TodoEditTask)     // Edit TODO task
	e.POST("/todos/:id/delete-task/:task_id", TodoDeleteTask) // Delete TODO task
	e.POST("/todos/:id/reorder-task/:task_id", welcome)       // Delete TODO task

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
