package pg

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
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
}
