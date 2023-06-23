package postgres

import "github.com/jmoiron/sqlx"

type Task struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) *Task {
	return &Task{
		db: db,
	}
}
