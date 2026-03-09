package database

import (
	"database/sql"
	"fmt"
	"restapi-tasks/internal/models"

	"github.com/jmoiron/sqlx"
)

type TaskStore struct {
	db *sqlx.DB
}

func newTaskStore(db *sqlx.DB) *TaskStore {
	return &TaskStore{db}
}

func (s *TaskStore) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	query := `
SELECT id, title, description, completed, created_at, updated_at 
FROM tasks 
ORDER BY created_at DESC;`

	err := s.db.Select(&tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskStore) GetByID(id int) (*models.Task, error) {
	var task models.Task

	query := `
SELECT id, title, description, completed, created_at, updated_at 
FROM tasks 
where id = $1;` //SQLAttack Def
	err := s.db.Get(&task, query, id)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("task whith id: %d not found", id)
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}
