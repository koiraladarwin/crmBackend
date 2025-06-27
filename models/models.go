package models

type User struct {
	ID    string
	Name  string
	Gmail string
	Phone string
}

type Company struct {
	ID   string
	Name string
}

type Employee struct {
	ID        string
	UserID    string
	CompanyID string
	Role      string
}

type Client struct {
	ID    string
	Name  string
	Gmail string
	Phone string
}

type ClientProcess struct {
	ClientID           string
	AssignedEmployeeID string
	ExpectedRevenue    float64
	Priority           string
	Status             string
}

type Schedule struct {
	ID                   string
	ProcessClientID      string
	ProcessAssignedEmpID string
	Schedule             string
}
