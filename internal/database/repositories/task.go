package repositories

import (
	"database/sql"
	"rest-api/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (s *TaskRepo) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	query := `
		SELECT *
		FROM tasks;
	`

	err := s.db.Select(&tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskRepo) GetById(id uuid.UUID) (*models.Task, error) {
	var task models.Task

	query := `
		SELECT *
		FROM tasks
		WHERE id = $1
	`

	err := s.db.Get(&task, query, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *TaskRepo) Create(task models.CreateTask) (*models.Task, error) {

}
