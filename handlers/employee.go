package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/koiraladarwin/crmbackend/models"
)

func (h *Handler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
  newId := uuid.New().String()

	err := h.DB.AddEmployee(newId, emp.UserID, emp.CompanyID, emp.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	emp, err := h.DB.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(emp)
}

func (h *Handler) GetEmployeesByCompanyID(w http.ResponseWriter, r *http.Request) {
	companyID := mux.Vars(r)["id"]

	employees, err := h.DB.GetEmployeesByCompanyID(companyID)
	if err != nil {
		http.Error(w, "Failed to fetch employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
