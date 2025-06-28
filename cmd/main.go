package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koiraladarwin/crmbackend/handlers"
	postgres "github.com/koiraladarwin/crmbackend/storage/postgress"
)

func main() {
	connStr := "postgres://postgres:mysecretpassword@localhost:5432/crm?sslmode=disable"

	db, err := postgres.NewPostGres("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	h := &handlers.Handler{DB: db}

	r := mux.NewRouter()
	setupRoutes(r, h)

	log.Fatal(http.ListenAndServe(":8080", handlers.EnableCORS(r)))
}

func setupRoutes(r *mux.Router, h *handlers.Handler) {
	r.HandleFunc("/adduser", h.AddUser).Methods("POST")
	r.HandleFunc("/getusers", h.GetAllUsers).Methods("GET")
	r.HandleFunc("/getuser/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/getjwt", h.GetJwt).Methods("POST")

	r.HandleFunc("/addcompany", h.AddCompany).Methods("POST")
	r.HandleFunc("/getcompany/{id}", h.GetCompany).Methods("GET")
	r.HandleFunc("/getcompaniesbyuser/{id}", h.GetCompaniesByUserID).Methods("GET")

	r.HandleFunc("/addemployee", h.AddEmployee).Methods("POST")
	r.HandleFunc("/getemployeesbycompany/{id}", h.GetEmployeesByCompanyID).Methods("GET")
	r.HandleFunc("/getemployee/{id}", h.GetEmployee).Methods("GET")

	r.HandleFunc("/addclient", h.AddClient).Methods("POST")
	r.HandleFunc("/getclient/{id}", h.GetClient).Methods("GET")
	r.HandleFunc("/getclientsbycompany/{id}", h.GetClientsByCompanyID).Methods("GET")

	r.HandleFunc("/addprocess", h.AddClientProcess).Methods("POST")
	r.HandleFunc("/getprocessbyclient/{id}", h.GetClientProcessByClientID).Methods("GET")

	r.HandleFunc("/addschedule", h.AddSchedule).Methods("POST")
	r.HandleFunc("/getschedulesbycompany/{id}", h.GetSchedulesByCompanyID).Methods("GET")
	r.HandleFunc("/getschedulesbyprocess/{client_id}/{employee_id}", h.GetSchedulesByProcess).Methods("GET")
}
