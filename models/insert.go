package models

import "go-queue-management-http/db"

func Insert(job Job) (id int64, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO jobs (name, result, status, data) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(sql, job.Name, job.Result, job.Status, job.Data).Scan(&id)
	return
}
