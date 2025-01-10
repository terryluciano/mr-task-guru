package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type postgres struct {
	db *pgx.Conn
}

var conn *postgres

func Connect(ctx context.Context, connString string) {
	db, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	} else {
		conn = &postgres{db}
		log.Println("Connected to Database...")
	}
}
