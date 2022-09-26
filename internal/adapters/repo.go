package adapters

import (
	"github.com/jinzhu/gorm"
	"github.com/jlipa/todothis/internal/entities"
)

type repo struct {
	db *gorm.DB
}

type Repo interface {
	CreateTodo(t entities.Todo) (entities.Todo, error)
	GetTodo(id uint) (entities.Todo, error)
	CreateTask(todoID uint, t entities.Task) (entities.Task, error)
	UpdateTaskSortID(todoID uint) (entities.Todo, error)
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db}
}

// skeleton
func (r *repo) GetTodo(id uint) (entities.Todo, error) {
	return entities.Todo{}, nil
}

// skeleton
func (r *repo) CreateTodo(t entities.Todo) (entities.Todo, error) {
	return t, nil
}

// skeleton
func (r *repo) CreateTask(todoID uint, t entities.Task) (entities.Task, error) {
	return t, nil
}

// skeleton
func (r *repo) EditTask(todoID uint, t entities.Task) (entities.Task, error) {
	return t, nil
}

// skeleton
func (r *repo) RemoveTask(todoID uint, taskID uint) (entities.Task, error) {
	return entities.Task{}, nil
}

// skeleton
func (r *repo) UpdateTaskSortID(todoID uint) (entities.Todo, error) {
	return entities.Todo{}, nil
}

// skeleton
func (r *repo) DecrementTaskSortIDsFrom(todoID uint, sortID uint, exclusion uint) (entities.Todo, error) {
	return entities.Todo{}, nil
}

// skeleton
func (r *repo) IncrementTaskSortIDsFrom(todoID uint, sortID uint, exclusion uint) (entities.Todo, error) {
	return entities.Todo{}, nil
}
