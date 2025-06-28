package storage

import "github.com/koiraladarwin/crmbackend/models"

type Database interface {
	AddUser(id, password, name, gmail, phone string) error
	GetUserByID(id string) (models.User, error)
	GetUserByGmailandPassword(gmail string, password string) (models.User, error)
	GetAllUsers() ([]models.User, error)

	AddCompany(id, name string) error
	GetCompanyByID(id string) (models.Company, error)
	GetCompaniesByUserID(userID string) ([]models.Company, error)

	AddEmployee(id, userID, companyID, role string) error
	GetEmployeeByID(id string) (models.Employee, error)
	GetEmployeesByCompanyID(companyID string) ([]models.Employee, error)

	AddClient(id, companyId, name, gmail, phone string) error
	GetClientByID(id string) (models.Client, error)

	AddClientProcess(clientID, employeeID string, revenue float64, priority, status string) error
	GetClientProcessByClientID(clientID string) (models.ClientProcess, error)
	GetClientsByCompanyID(companyID string) ([]models.Client, error)

	AddSchedule(id, clientID, employeeID, schedule string) error
	GetSchedulesByProcess(clientID, employeeID string) ([]models.Schedule, error)
	GetSchedulesByCompanyID(companyID string) ([]models.Schedule, error)
}
