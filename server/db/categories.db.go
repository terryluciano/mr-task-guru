package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/terryluciano/mr-task-guru/server/model"
)

func InsertCategory(category string) (*model.Category, error) {
	var newCategory model.Category

	row := conn.db.QueryRow(context.Background(), `INSERT INTO categories (name) VALUES ($1) RETURNING *`, category)

	if err := row.Scan(&newCategory.ID, &newCategory.Name, &newCategory.Deleted, &newCategory.Timestamp); err != nil {
		return nil, err
	}
	return &newCategory, nil
}

func RemoveCategory(id int) error {
	_, err := conn.db.Exec(context.Background(), `UPDATE categories SET deleted = true WHERE id = $1`, id)
	return err

}

func SelectCategory(id int) (*model.Category, error) {
	var category model.Category

	row := conn.db.QueryRow(context.Background(), `SELECT * FROM categories WHERE id=$1`, id)

	switch err := row.Scan(&category.ID, &category.Name, &category.Deleted, &category.Timestamp); err {
	case sql.ErrNoRows:
		log.Println(err.Error())
		return nil, fmt.Errorf("Category does not exist")
	case nil:
		return &category, err
	default:
		log.Println(err.Error())
		return nil, err
	}

}

func SelectAllCategories() ([]model.Category, error) {
	rows, err := conn.db.Query(context.Background(), `SELECT * FROM categories WHERE deleted = false`)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	outputRows, rowsErr := pgx.CollectRows(rows, pgx.RowToStructByName[model.Category])
	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	return outputRows, nil
}
