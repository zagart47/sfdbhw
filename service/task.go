package service

import (
	"context"
	"sfdbhw/entity"
)

type Tasker interface {
	NewTask(ctx context.Context, task entity.Task) (int, error)
	AllTasks(ctx context.Context) ([]entity.Task, error)
	TaskByAuthor(ctx context.Context, author int) ([]entity.Task, error)
	TaskByLabel(ctx context.Context, label int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task entity.Task) error
	DeleteTask(ctx context.Context, id int64) error
}

func (s Services) NewTask(ctx context.Context, task entity.Task) (int, error) {
	return s.taskStorage.Tasker.NewTask(ctx, task)
}

func (s Services) AllTasks(ctx context.Context) ([]entity.Task, error) {
	return s.taskStorage.Tasker.Tasks(ctx)
}

func (s Services) TaskByAuthor(ctx context.Context, author int) ([]entity.Task, error) {
	return s.taskStorage.Tasker.TasksByAuthor(ctx, author)
}

func (s Services) TaskByLabel(ctx context.Context, label int) ([]entity.Task, error) {
	return s.taskStorage.Tasker.TasksByLabel(ctx, label)
}

func (s Services) UpdateTask(ctx context.Context, task entity.Task) error {
	return s.taskStorage.Tasker.UpdateTask(ctx, task)
}

func (s Services) DeleteTask(ctx context.Context, id int) error {
	return s.taskStorage.Tasker.DeleteTask(ctx, id)
}
