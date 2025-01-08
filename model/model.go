package model

import (
	"database/sql"
	"time"
)

type Tasks struct {
	ID             int            `db:"id"`
	Name           string         `db:"name"`
	Category       sql.NullInt64  `db:"category"`
	Minutes        sql.NullInt64  `db:"minutes"`
	Current_status sql.NullString `db:"current_status"`
	Deleted        sql.NullBool   `db:"deleted"`
	Timestamp      time.Time      `db:"timestamp"`
}
