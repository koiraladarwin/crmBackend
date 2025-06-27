package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/koiraladarwin/crmbackend/models"
)

func (h *Handler) AddClientProcess(w http.ResponseWriter, r *http.Request) {
	var cp models.ClientProcess
	if err := json.NewDecoder(r.Body).Decode(&cp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err := h.DB.AddClientProcess(cp.ClientID, cp.AssignedEmployeeID, cp.ExpectedRevenue, cp.Priority, cp.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetClientProcessByClientID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	
	process, err := h.DB.GetClientProcessByClientID(id)
	if err != nil {
		http.Error(w, "Client process not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(process); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
