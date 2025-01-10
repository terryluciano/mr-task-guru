package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/terryluciano/mr-task-guru/model"
)

func InsertTask(task *model.Task) error {
	if _, err := conn.db.Exec(context.Background(), `INSERT INTO tasks (name, category, minutes, current_status) VALUES ($1, $2, $3, $4)`, task.Name, task.Category, task.Minutes, task.Current_status); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Println(pgErr.Message)
			if pgErr.Code == "23503" {
				return fmt.Errorf("Category does not exist")
			}
		}
		return err
	}

	return nil
}

func SelectTask(id int) (*model.Task, error) {
	var task model.Task

	row := conn.db.QueryRow(context.Background(), `SELECT * FROM tasks WHERE id=$1`, id)
	switch err := row.Scan(&task.ID, &task.Name, &task.Category, &task.Minutes, &task.Current_status, &task.Deleted, &task.Timestamp); err {
	case sql.ErrNoRows:
		log.Println(err.Error())
		return nil, fmt.Errorf("Task does not exist")
	case nil:
		return &task, err
	default:
		log.Println(err.Error())
		return nil, err
	}

}

func SelectAllTasks() ([]model.Task, error) {
	rows, err := conn.db.Query(context.Background(), `SELECT * FROM tasks WHERE deleted = false`)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	outputRows, rowsErr := pgx.CollectRows(rows, pgx.RowToStructByName[model.Task])
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	return outputRows, nil
}

func UpdateTask(id int, task model.Task) error {
	_, err := conn.db.Exec(context.Background(), `UPDATE tasks SET name = $1, category = $2, minutes = $3, current_status = $4 WHERE id = $5`, task.Name, task.Category, task.Minutes, task.Current_status, id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Println(pgErr.Message)
			if pgErr.Code == "23503" {
				return fmt.Errorf("Category does not exist")
			}
		}
		return err
	}

	return nil
}

func RemoveTask(id int) error {
	_, err := conn.db.Exec(context.Background(), `UPDATE tasks SET deleted = true WHERE id = $1`, id)
	return err
}
