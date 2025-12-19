# User Age API

A RESTful API built with Go to manage users with their **name** and **date of birth (dob)**, calculating their **age** dynamically.

---

## 🚀 Setup & Run

### 1. Clone the repository

```bash
git clone [<repo-url>](https://github.com/Sq1754/Go-backend-task.git)
```

### 2. Docker Setup

The project uses Docker to run both the API and PostgreSQL.

- **Build and start containers:**

```bash
docker compose down -v   # optional: remove existing volumes for fresh init
docker compose up --build
```

- The API runs at: `http://localhost:3000`  
- PostgreSQL runs at: `localhost:5432`  

> The database is automatically initialized, and the `users` table is created from migrations.

### 3. API Environment Variables

The API reads the database URL from `DB_DSN`. Default:

```
postgres://postgres:postgres@db:5432/userdb?sslmode=disable
```

### 4. Testing the API

You can test endpoints using **curl**, **Postman**, or PowerShell's `Invoke-RestMethod`.  

Example to create a user:

```powershell
Invoke-RestMethod `
  -Method POST `
  -Uri http://localhost:3000/users `
  -Headers @{ "Content-Type" = "application/json" } `
  -Body '{"name":"Alice","dob":"1990-05-10"}'
```

---




