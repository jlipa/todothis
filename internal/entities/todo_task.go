package entities

import "time"

type TodoTask struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	TodoID    uint      `json:"todo_id"`
	TaskID    uint      `json:"task_id"`
	SortID    uint      `json:"sort_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
