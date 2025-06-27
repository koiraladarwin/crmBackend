# 📘 CRM API Documentation
**Base URL:** `http://localhost:8080`


## 👤 Users

### ➕ `POST /adduser`
**Description:** Create a new user  
**Request Body:**
```json
{
  "name": "John Doe",
  "gmail": "john.doe@example.com",
  "phone": "9812345678"
}
```

---

### 📄 `GET /getusers`
**Description:** Get all users

---

### 📄 `GET /getuser/{id}`
**Description:** Get user by ID  
**Example:** `/getuser/550e8400-e29b-41d4-a716-446655440000`


## 🏢 Companies

### ➕ `POST /addcompany`
**Description:** Create a new company  
**Request Body:**
```json
{
  "name": "Acme Corp"
}
```

---

### 📄 `GET /getcompany/{id}`
**Description:** Get company by ID  
**Example:** `/getcompany/f47ac10b-58cc-4372-a567-0e02b2c3d479`

---

### 📄 `GET /getcompaniesbyuser/{id}`
**Description:** Get companies by user ID  
**Example:** `/getcompaniesbyuser/550e8400-e29b-41d4-a716-446655440000`


## 👥 Employees

### ➕ `POST /addemployee`
**Description:** Add an employee  
**Request Body:**
```json
{
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "company_id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
  "role": "Manager"
}
```

---

### 📄 `GET /getemployee/{id}`
**Description:** Get employee by ID  
**Example:** `/getemployee/1c6fbe1d-9f1a-4d6e-9e35-42df39b8579e`

---

### 📄 `GET /getemployeesbycompany/{id}`
**Description:** Get employees by company ID  
**Example:** `/getemployeesbycompany/f47ac10b-58cc-4372-a567-0e02b2c3d479`


## 👤 Clients

### ➕ `POST /addclient`
**Description:** Create a new client  
**Request Body:**
```json
{
  "name": "Jane Smith",
  "gmail": "jane@example.com",
  "phone": "9876543210"
}
```

---

### 📄 `GET /getclient/{id}`
**Description:** Get client by ID  
**Example:** `/getclient/9b2d0c30-4a6e-4c55-8d6e-9f11c1326e07`

---

### 📄 `GET /getclientsbycompany/{id}`
**Description:** Get clients by company ID  
**Example:** `/getclientsbycompany/f47ac10b-58cc-4372-a567-0e02b2c3d479`


## 🔄 Client Processes

### ➕ `POST /addprocess`
**Description:** Add a new process to a client  
**Request Body:**
```json
{
  "client_id": "9b2d0c30-4a6e-4c55-8d6e-9f11c1326e07",
  "assigned_employee_id": "1c6fbe1d-9f1a-4d6e-9e35-42df39b8579e",
  "expected_revenue": 50000,
  "priority": "High",
  "status": "Pending"
}
```

---

### 📄 `GET /getprocessbyclient/{id}`
**Description:** Get process by client ID  
**Example:** `/getprocessbyclient/9b2d0c30-4a6e-4c55-8d6e-9f11c1326e07`


## 📅 Schedules

### ➕ `POST /addschedule`
**Description:** Add a schedule for a process  
**Request Body:**
```json
{
  "process_client_id": "9b2d0c30-4a6e-4c55-8d6e-9f11c1326e07",
  "process_assigned_employee_id": "1c6fbe1d-9f1a-4d6e-9e35-42df39b8579e",
  "schedule": "Meeting on 2025-07-01 at 10:00 AM"
}
```

---

### 📄 `GET /getschedulesbycompany/{id}`
**Description:** Get schedules for a company  
**Example:** `/getschedulesbycompany/f47ac10b-58cc-4372-a567-0e02b2c3d479`

---

### 📄 `GET /getschedulesbyprocess/{client_id}/{employee_id}`
**Description:** Get schedules for a specific client process  
**Example:**  
`/getschedulesbyprocess/9b2d0c30-4a6e-4c55-8d6e-9f11c1326e07/1c6fbe1d-9f1a-4d6e-9e35-42df39b8579e`
