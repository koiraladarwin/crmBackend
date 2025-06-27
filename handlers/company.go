package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koiraladarwin/crmbackend/models"
)

func (h *Handler) AddCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.DB.AddCompany(company.ID, company.Name); err != nil {
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
