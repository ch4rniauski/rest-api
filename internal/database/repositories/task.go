package repositories

import (
	"database/sql"
	"rest-api/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	query := `
		SELECT *
		FROM tasks;
	`

	err := r.db.Select(&tasks, query)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepo) GetById(id uuid.UUID) (*models.Task, error) {
	var task models.Task

	query := `
		SELECT *
		FROM tasks
		WHERE id = $1;
	`

	err := r.db.Get(&task, query, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepo) Create(inputTask models.CreateTask) (*models.Task, error) {
	var task models.Task

	query := `
		INSERT INTO tasks (id, title, description, created_at, is_completed)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, description, created_at, is_completed;
	`

	created_at := time.Now()
	id := uuid.New()

	err := r.db.QueryRowx(query, id, inputTask.Title, inputTask.Description, created_at, false).StructScan(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepo) Update(id uuid.UUID, input models.UpdateTask) (*models.Task, error) {
	var task models.Task

	query := `
		UPDATE tasks
		SET title = $1, description = $2, is_completed = $3
		WHERE id = $4
		RETURNING id, title, description, created_at, is_completed;
	`

	err := r.db.QueryRowx(query, input.Title, input.Description, input.IsCompleted, id).StructScan(&task)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &task, nil
}
