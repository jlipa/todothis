package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jlipa/todothis/internal/entities"
	"github.com/labstack/echo"
)

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to TODOTHIS!")
}

// Fetch all TODO
func TodoGetAll(c echo.Context) error {
	output, err := service.GetTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, output)
	}

	return c.JSON(http.StatusOK, output)
}

// Fetch TODO and tasks
func TodoGet(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	output, err := service.GetTodo(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, output)
	}

	return c.JSON(http.StatusOK, output)
}

// Create a TODO item
func TodoNew(c echo.Context) error {
	var todo entities.Todo
	inputTodo := c.FormValue("todo")
	json.Unmarshal([]byte(inputTodo), &todo)

	if todo.Title != "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	output, err := service.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, output)
	}

	return c.JSON(http.StatusOK, output)
}

// Create a TODO task
func TodoAddTask(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var task entities.Task
	inputTask := c.FormValue("task")
	json.Unmarshal([]byte(inputTask), &task)

	if task.Content != "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	output, err := service.AddTodoTask(uint(id), task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, output)
	}

	return c.JSON(http.StatusOK, output)
}

// Edit a TODO task
func TodoEditTask(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	taskID, err := strconv.ParseUint(c.Param("task_id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var task entities.Task
	inputTask := c.FormValue("task")
	json.Unmarshal([]byte(inputTask), &task)

	output, err := service.EditTodoTask(uint(id), uint(taskID), task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, output)
	}

	return c.JSON(http.StatusOK, output)
}

// Delete a TODO task
func TodoDeleteTask(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	taskID, err := strconv.ParseUint(c.Param("task_id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	output, err := service.RemoveTodoTask(uint(id), uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, output)
	}

	return c.JSON(http.StatusOK, output)
}

// Reorder/reposition a TODO task
func TodoReorderTask(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	taskID, err := strconv.ParseUint(c.Param("task_id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	newSortID, err := strconv.ParseUint(c.FormValue("sort_id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	output, err := service.ReorderTodoTask(uint(id), uint(taskID), uint(newSortID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, output)
	}

	return c.JSON(http.StatusOK, output)
}
