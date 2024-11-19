package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sfdbhw/entity"
)

type TaskStorage struct {
	db Client
}

func NewTaskStorage(db *pgxpool.Pool) TaskStorage {
	return TaskStorage{db: db}
}

type Tasker interface {
	NewTask(ctx context.Context, task entity.Task) (int, error)
	Tasks(ctx context.Context) ([]entity.Task, error)
	TasksByAuthor(ctx context.Context, author int) ([]entity.Task, error)
	TasksByLabel(ctx context.Context, label int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

func (s *TaskStorage) NewTask(ctx context.Context, task entity.Task) (int, error) {
	var id int
	err := s.db.QueryRow(ctx, `
		INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, task.Opened, task.Closed, task.AuthorID, task.AssignedID, task.Title, task.Content).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}

	for _, label := range task.Labels {
		_, err = s.db.Exec(ctx, `
		INSERT INTO tasks_labels (task_id, label_id)
		VALUES ($1, $2)
	`, id, label.ID)
		if err != nil {
			return 0, fmt.Errorf("failed to add default label to task: %w", err)
		}
	}
	return id, nil
}

func (s *TaskStorage) Tasks(ctx context.Context) ([]entity.Task, error) {
	rows, err := s.db.Query(ctx, `
		SELECT t.id, t.opened, t.closed, t.author_id, t.assigned_id, t.title, t.content, l.id, l.name
		FROM tasks t
		LEFT JOIN tasks_labels tl ON t.id = tl.task_id
		LEFT JOIN labels l ON tl.label_id = l.id
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}
	defer rows.Close()

	taskMap := make(map[int]*entity.Task)
	for rows.Next() {
		var task entity.Task
		var label entity.Label
		err := rows.Scan(&task.ID, &task.Opened, &task.Closed, &task.AuthorID, &task.AssignedID, &task.Title, &task.Content, &label.ID, &label.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		if _, ok := taskMap[task.ID]; !ok {
			taskMap[task.ID] = &task
		}
		if label.ID != 0 {
			taskMap[task.ID].Labels = append(taskMap[task.ID].Labels, label)
		}
	}

	var tasks []entity.Task
	for _, task := range taskMap {
		tasks = append(tasks, *task)
	}

	return tasks, nil
}

func (s *TaskStorage) TasksByAuthor(ctx context.Context, authorID int) ([]entity.Task, error) {
	rows, err := s.db.Query(ctx, `
		SELECT t.id, t.opened, t.closed, t.author_id, t.assigned_id, t.title, t.content, l.id, l.name
		FROM tasks t
		LEFT JOIN tasks_labels tl ON t.id = tl.task_id
		LEFT JOIN labels l ON tl.label_id = l.id
		WHERE t.author_id = $1
	`, authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by author: %w", err)
	}
	defer rows.Close()

	taskMap := make(map[int]*entity.Task)
	for rows.Next() {
		var task entity.Task
		var label entity.Label
		err := rows.Scan(&task.ID, &task.Opened, &task.Closed, &task.AuthorID, &task.AssignedID, &task.Title, &task.Content, &label.ID, &label.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		if _, ok := taskMap[task.ID]; !ok {
			taskMap[task.ID] = &task
		}
		if label.ID != 0 {
			taskMap[task.ID].Labels = append(taskMap[task.ID].Labels, label)
		}
	}

	var tasks []entity.Task
	for _, task := range taskMap {
		tasks = append(tasks, *task)
	}

	return tasks, nil
}

func (s *TaskStorage) TasksByLabel(ctx context.Context, labelID int) ([]entity.Task, error) {
	rows, err := s.db.Query(ctx, `
		SELECT t.id, t.opened, t.closed, t.author_id, t.assigned_id, t.title, t.content, l.id, l.name
		FROM tasks t
		JOIN tasks_labels tl ON t.id = tl.task_id
		JOIN labels l ON tl.label_id = l.id
		WHERE tl.label_id = $1
	`, labelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by label: %w", err)
	}
	defer rows.Close()

	taskMap := make(map[int]*entity.Task)
	for rows.Next() {
		var task entity.Task
		var label entity.Label
		err := rows.Scan(&task.ID, &task.Opened, &task.Closed, &task.AuthorID, &task.AssignedID, &task.Title, &task.Content, &label.ID, &label.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		if _, ok := taskMap[task.ID]; !ok {
			taskMap[task.ID] = &task
		}
		if label.ID != 0 {
			taskMap[task.ID].Labels = append(taskMap[task.ID].Labels, label)
		}
	}

	var tasks []entity.Task
	for _, task := range taskMap {
		tasks = append(tasks, *task)
	}

	return tasks, nil
}

func (s *TaskStorage) UpdateTask(ctx context.Context, task entity.Task) error {
	_, err := s.db.Exec(ctx, `
		UPDATE tasks
		SET opened = $1, closed = $2, author_id = $3, assigned_id = $4, title = $5, content = $6
		WHERE id = $7
	`, task.Opened, task.Closed, task.AuthorID, task.AssignedID, task.Title, task.Content, task.ID)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	for _, label := range task.Labels {
		_, err = s.db.Exec(ctx, `
		UPDATE tasks_labels
		SET label_id = $1
		WHERE task_id = $2
	`, label.ID, task.ID)
		if err != nil {
			return fmt.Errorf("failed to update task: %w", err)
		}
	}

	return nil
}

func (s *TaskStorage) DeleteTask(ctx context.Context, id int) error {
	_, err := s.db.Exec(ctx, `
		DELETE FROM tasks
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
