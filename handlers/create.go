package handlers

import (
	"encoding/json"
	"fmt"
	"go-queue-management-http/models"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	accessToken := hasAuth(r)

	if accessToken == false {
		log.Printf("Unauthorized")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var job models.Job

	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(job)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	}

	job, err = models.GetById(id)

	resp = map[string]any{
		"Error": false,
		"Name":  &job.Name,
		"Data":  &job.Data,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func hasAuth(r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	if auth != "allow" {
		return false
	}
	return true
}
