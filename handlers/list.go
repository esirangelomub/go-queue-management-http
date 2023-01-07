package handlers

import (
	"encoding/json"
	"go-queue-management-http/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, request *http.Request) {
	status := request.URL.Query().Get("status")
	var jobs []models.Job
	var err error
	if status == "" {
		jobs, err = models.GetAll()
	} else {
		jobs, err = models.GetByStatus(status)
	}

	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
