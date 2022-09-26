package engines

import (
	"github.com/jlipa/todothis/internal/adapters"
	"github.com/jlipa/todothis/internal/entities"
)

type engine struct {
	db *adapters.Repo
}

type Engine interface {
	GetTodos() ([]entities.Todo, error)
	GetTodo(id uint) (entities.Todo, error)
	CreateTodo(entities.Todo) (entities.Todo, error)
	AddTodoTask(todoID uint, task entities.Task) (entities.Todo, error)
	EditTodoTask(todoID uint, taskID uint, task entities.Task) (entities.Todo, error)
	RemoveTodoTask(todoID uint, taskID uint) (entities.Todo, error)
	ReorderTodoTask(todoID uint, taskID uint, sortID uint) (entities.Todo, error)
}

func NewEngine(r *adapters.Repo) Engine {
	return &engine{db: r}
}

// skeleton
func (e *engine) GetTodos() ([]entities.Todo, error) {
	// database call: Select * from todo
	// map output into []TODO struct
	return []entities.Todo{}, nil
}

// skeleton
func (e *engine) GetTodo(id uint) (entities.Todo, error) {
	// database call: Select * from todo where id = :id
	// map output into TODO struct
	return entities.Todo{}, nil
}

// skeleton
func (e *engine) CreateTodo(entities.Todo) (entities.Todo, error) {
	// database call: insert into table TODO with passed todo values
	// map output into TODO struct
	return entities.Todo{}, nil
}

// skeleton
func (e *engine) AddTodoTask(todoID uint, task entities.Task) (entities.Todo, error) {
	// database call: insert into table TODO with passed todo values
	// database call: Select * from todo where id = :id
	// database call:
	// 		select * from todo_tasks tt
	// 			inner join tasks ta ON tt.task_id = ta.id
	// 		sort by tt.sort_id asc
	// Map output into TODO struct
	return entities.Todo{}, nil
}

// skeleton
func (e *engine) EditTodoTask(todoID uint, taskID uint, task entities.Task) (entities.Todo, error) {
	// database call: update table TASK with passed task values
	// database call: Select * from todo where id = :id
	// database call:
	// 		select * from todo_tasks tt
	// 			inner join tasks ta ON tt.task_id = ta.id
	// 		sort by tt.sort_id asc
	// Map output into TODO struct
	return entities.Todo{}, nil
}

// skeleton
func (e *engine) RemoveTodoTask(todoID uint, taskID uint) (entities.Todo, error) {
	// Get task ID current sort ID
	// database call: DELETE from TASK and TODO_TASK tables
	// database call: update TODO_TASK
	// 					set sort_id = sort_id - 1
	//					where todo_id = todoID
	// 					and sort_id > deleted_task_sort_id
	// database call: Select * from todo where id = :id
	// database call:
	// 		select * from todo_tasks tt
	// 			inner join tasks ta ON tt.task_id = ta.id
	// 		sort by tt.sort_id asc
	// Map output into TODO struct
	return entities.Todo{}, nil
}

// skeleton
func (e *engine) ReorderTodoTask(todoID uint, taskID uint, sortID uint) (entities.Todo, error) {
	// database call: select * from TODO_TASK
	// Validations: if sort_index is within range
	// store task that was moved as moved_task_id and move_sort_id
	// database call: update TODO_TASK set sort_id = :sortID where task_id = :taskID
	// database call: update TODO_TASK
	//				  	set sort_id = sort_id + 1
	// 					where todo_id = todoID
	// 					and sort_id > sortID
	//					and task_id != taskID

	// database call: Select * from todo where id = :id
	// database call:
	// 		select * from todo_tasks tt
	// 			inner join tasks ta ON tt.task_id = ta.id
	// 		sort by tt.sort_id asc
	// Map output into TODO struct
	return entities.Todo{}, nil
}
