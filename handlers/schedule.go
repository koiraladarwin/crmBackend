package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koiraladarwin/crmbackend/models"
)

func (h *Handler) AddSchedule(w http.ResponseWriter, r *http.Request) {
	var s models.Schedule
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err := h.DB.AddSchedule(s.ID, s.ProcessClientID, s.ProcessAssignedEmpID, s.Schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetSchedulesByCompanyID(w http.ResponseWriter, r *http.Request) {
	companyID := mux.Vars(r)["id"]
	schedules, err := h.DB.GetSchedulesByCompanyID(companyID)
	if err != nil {
		http.Error(w, "Failed to fetch schedules", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

func (h *Handler) GetSchedulesByProcess(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID := vars["client_id"]
	employeeID := vars["employee_id"]

	schedules, err := h.DB.GetSchedulesByProcess(clientID, employeeID)
	if err != nil {
		http.Error(w, "Failed to fetch schedules", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}
