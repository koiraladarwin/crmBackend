package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/koiraladarwin/crmbackend/models"
)

func (h *Handler) AddCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Wrong format"})
		return
	}
  userId := company.ID; //change this later as this is weird 
	newComanyId := uuid.New().String()
	if err := h.DB.AddCompany(newComanyId, company.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newEmployeeid := uuid.New().String()
	if err := h.DB.AddEmployee(newEmployeeid,userId , newComanyId, "Owner"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetCompany(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	company, err := h.DB.GetCompanyByID(id)
	if err != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(company)
}

func (h *Handler) GetCompaniesByUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	companies, err := h.DB.GetCompaniesByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to fetch companies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}
