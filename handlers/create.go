package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davimelovasc/api-postgresql/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error ao fazer decode do json: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := models.Insert(todo)

	var resp = make(map[string]any)

	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": fmt.Sprintf("Ocorreu um erro ao tentar inserir %v", err),
		}
	} else {
		resp["error"] = false
		resp["message"] = fmt.Sprintf("Todo inserido com sucesso! ID %v", id)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
