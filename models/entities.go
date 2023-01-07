package models

import "database/sql"

type Job struct {
	ID     int64          `json:"id"`
	Name   string         `json:"name"`
	Result sql.NullString `json:"result"`
	Status string         `json:"status"`
	Data   string         `json:"data"`
}
