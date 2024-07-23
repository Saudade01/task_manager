package services

import (
	"context"
	"database/sql"
	"errors"
	"task_manager/config"
	"task_manager/models"
)

func CreateTask(ctx context.Context, task models.Task) (int64, error) {
	tx, err := config.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := "INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, task.Title, task.Description, task.Completed)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return id, nil
}

func GetTask(ctx context.Context, id int) (*models.Task, error) {
	task := &models.Task{}
	query := "SELECT id, title, description, completed FROM tasks WHERE id = ?"
	err := config.DB.QueryRowContext(ctx, query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return task, nil
}

func ListTasks(ctx context.Context) ([]models.Task, error) {
	rows, err := config.DB.QueryContext(ctx, "SELECT id, title, description, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func UpdateTask(ctx context.Context, task models.Task) error {
	tx, err := config.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, query, task.Title, task.Description, task.Completed, task.ID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func DeleteTask(ctx context.Context, id int) error {
	tx, err := config.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "DELETE FROM tasks WHERE id = ?"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
