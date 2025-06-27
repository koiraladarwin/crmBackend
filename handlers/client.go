package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koiraladarwin/crmbackend/models"
)

func (h *Handler) AddClient(w http.ResponseWriter, r *http.Request) {
	var c models.Client
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err := h.DB.AddClient(c.ID, c.Name, c.Gmail, c.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetClient(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	client, err := h.DB.GetClientByID(id)
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(client)
}

func (h *Handler) GetClientsByCompanyID(w http.ResponseWriter, r *http.Request) {
	companyID := mux.Vars(r)["id"]

	clients, err := h.DB.GetClientsByCompanyID(companyID)
	if err != nil {
		http.Error(w, "Failed to fetch clients", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}
