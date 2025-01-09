package pg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/terryluciano/mr-task-guru/model"
)

func InsertTask(task *model.Task) error {
	if _, err := pgConn.db.Exec(context.Background(), `INSERT INTO tasks (name, category, minutes, current_status) VALUES ($1, $2, $3, $4)`, task.Name, task.Category, task.Minutes, task.Current_status); err != nil {
		return err
	}

	return nil
}

func SelectAllTasks() ([]model.Task, error) {
	// Example query for testing: Fetch all tasks
	rows, err := pgConn.db.Query(context.Background(), `SELECT * FROM tasks`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	outputRows, rowsErr := pgx.CollectRows(rows, pgx.RowToStructByName[model.Task])
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	return outputRows, nil
}
