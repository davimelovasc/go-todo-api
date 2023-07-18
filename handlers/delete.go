package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/davimelovasc/api-postgresql/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error ao fazer parse do id: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error excluir registro id: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: Foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"error":   false,
		"message": "Registro removido com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
