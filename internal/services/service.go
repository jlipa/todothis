package services

import (
	"github.com/jlipa/todothis/internal/adapters"
	"github.com/jlipa/todothis/internal/engines"
	"github.com/jlipa/todothis/internal/entities"
)

type service struct {
	e *engines.Engine
}

func NewService(repo *adapters.Repo) Service {
	en := engines.NewEngine(repo)
	return &service{e: &en}
}

type Service interface {
	GetTodos() ([]entities.Todo, error)
	GetTodo(id uint) (entities.Todo, error)
	CreateTodo(t entities.Todo) (entities.Todo, error)
	AddTodoTask(id uint, t entities.Task) (entities.Todo, error)
	EditTodoTask(id uint, taskID uint, t entities.Task) (entities.Todo, error)
	RemoveTodoTask(id uint, taskID uint) (entities.Todo, error)
	ReorderTodoTask(id uint, taskID uint, newSortIndex uint) (entities.Todo, error)
}

func (s *service) GetTodos() ([]entities.Todo, error) {
	// call to engine
	return []entities.Todo{}, nil
}

func (s *service) GetTodo(id uint) (entities.Todo, error) {
	// call to engine
	return entities.Todo{}, nil
}

func (s *service) CreateTodo(t entities.Todo) (entities.Todo, error) {
	// call to engine
	return entities.Todo{}, nil
}

func (s *service) AddTodoTask(id uint, t entities.Task) (entities.Todo, error) {
	// call to engine
	return entities.Todo{}, nil
}

func (s *service) EditTodoTask(id uint, taskID uint, t entities.Task) (entities.Todo, error) {
	// call to engine
	return entities.Todo{}, nil
}

func (s *service) RemoveTodoTask(id uint, taskID uint) (entities.Todo, error) {
	// call to engine
	return entities.Todo{}, nil
}

func (s *service) ReorderTodoTask(id uint, taskID uint, newSortIndex uint) (entities.Todo, error) {
	// call to engine
	return entities.Todo{}, nil
}

// Basic business rules
// The user should be able to list all tasks in the TODO list (GetTodo)
// The user should be able to add a task to the TODO list (AddTodoTask)
// The user should be able to update the details of a task in the TODO list (EditTodoTask)
// The user should be able to remove a task from the TODO list (RemoveTodoTask)
// The user should be able to reorder the tasks in the TODO list (ReorderTodoTask)
// A task in the TODO list should be able to handle being moved more than 50 times (ReorderTodoTask)
// A task in the TODO list should be able to handle being moved to more than one task away from its current position (ReorderTodoTask)
