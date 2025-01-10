package model

import (
	"time"
)

type Task struct {
	ID             int       `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	Category       *int      `db:"category" json:"category"`
	Minutes        *int      `db:"minutes" json:"minutes"`
	Current_status *string   `db:"current_status" json:"current_status"`
	Deleted        *bool     `db:"deleted" json:"deleted"`
	Timestamp      time.Time `db:"timestamp" json:"timestamp"`
}

type Category struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Deleted   *bool     `db:"deleted" json:"deleted"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
}
