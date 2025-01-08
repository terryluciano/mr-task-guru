package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	model "github.com/terryluciano/mr-task-guru/model"
)

type postgres struct {
	db *pgx.Conn
}

var pgConn *postgres

func Connect(ctx context.Context, connString string) {
	db, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	} else {
		pgConn = &postgres{db}
		log.Println("Connected to Database...")
	}
	defer db.Close(ctx)

	// Example query for testing: Fetch all tasks
	rows, err := pgConn.db.Query(ctx, `SELECT * FROM tasks`)
	if err != nil {
		log.Fatal("Unable to fetch tasks")
		return
	}
	defer rows.Close()

	rowsfff, rowErr := pgx.CollectRows(rows, pgx.RowToStructByName[model.Tasks])
	if rowErr != nil {
		log.Println(rowErr)
		log.Fatal("Unable to format tasks")
		return
	}

	fmt.Println(rowsfff)
	defer rows.Close()
}
