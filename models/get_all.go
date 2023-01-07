package models

import (
	"go-queue-management-http/db"
)

func GetAll() (jobs []Job, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM jobs`)
	if err != nil {
		return
	}

	for rows.Next() {
		var job Job
		//var data map[interface{}]interface{}
		//json.Unmarshal([]byte(job.Data), &data)

		err = rows.Scan(&job.ID, &job.Name, &job.Result, &job.Status, &job.Data)
		if err != nil {
			continue
		}

		jobs = append(jobs, job)
	}

	return
}

func GetByStatus(status string) (jobs []Job, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM jobs WHERE status = $1`, status)
	if err != nil {
		return
	}

	for rows.Next() {
		var job Job
		//var data map[interface{}]interface{}
		//json.Unmarshal([]byte(job.Data), &data)

		err = rows.Scan(&job.ID, &job.Name, &job.Result, &job.Status, &job.Data)
		if err != nil {
			continue
		}

		jobs = append(jobs, job)
	}

	return
}

func GetById(id int64) (job Job, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM jobs WHERE id=$1`, id)
	err = row.Scan(&job.ID, &job.Name, &job.Result, &job.Status, &job.Data)

	return
}
