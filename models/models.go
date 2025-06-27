package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Gmail string `json:"gmail"`
	Phone string `json:"phone"`
}

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	Role      string `json:"role"`
}


type Client struct {
	ID        string `json:"id"`
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Gmail     string `json:"gmail"`
	Phone     string `json:"phone"`
}


type ClientProcess struct {
	ClientID           string  `json:"client_id"`
	AssignedEmployeeID string  `json:"assigned_employee_id"`
	ExpectedRevenue    float64 `json:"expected_revenue"`
	Priority           string  `json:"priority"`
	Status             string  `json:"status"`
}

type Schedule struct {
	ID                   string `json:"id"`
	ProcessClientID      string `json:"process_client_id"`
	ProcessAssignedEmpID string `json:"process_assigned_employee_id"`
	Schedule             string `json:"schedule"`
}

