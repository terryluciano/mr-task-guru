package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var conn *postgres
var ctx = context.Background()

func Connect(connString string) {

	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	} else {
		conn = &postgres{db}
		log.Println("Connected to Database...")
	}
}
