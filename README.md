# 📅 Go REST API v2

A RESTful API built with **Go**, using **Gin** and **SQLite**, featuring JWT-based authentication, user management, and event CRUD operations with registration functionality.

---

## 🚀 Features

- Secure **User Signup** and **Login** with bcrypt password hashing
- **JWT Authentication** middleware to protect routes
- Full **Event Management**: Create, Read, Update, Delete events
- User **Event Registration** and cancellation endpoints
- Clear modular MVC-style project organization
- SQLite for lightweight local persistence

---

## 📁 Project Structure

```
go-rest-api-v2/
│
├── main.go # Application entry point
├── db/ # Database setup and schema
│ └── db.go
├── models/ # Business logic and database models
│ ├── user.go
│ └── event.go
├── routes/ # HTTP route handlers
│ ├── users.go
│ ├── events.go
│ ├── routes.go
│ └── register.go
├── middlewares/ # Middleware for authentication
│ └── auth.go
├── utils/ # Utilities (e.g., password hashing, JWT tokens)
│ ├── hash.go
│ └── jwt.go
├── api.db # SQLite database file
├── go.mod
├── go.sum
└── .gitignore
```


## 🧪 API Endpoints

### Authentication

| Method | Endpoint | Description                   |
|--------|----------|-------------------------------|
| POST   | /signup  | Register a new user            |
| POST   | /login   | Authenticate user & receive JWT |

#### Request body for signup/login:

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}

```

### 📆 Events

| Method | Endpoint              | Description                     | Auth Required |
| ------ | --------------------- | ------------------------------- | ------------- |
| GET    | /events               | Get all events                  | No            |
| GET    | /events/\:id          | Get event by ID                 | No            |
| POST   | /events               | Create new event                | Yes           |
| PUT    | /events/\:id          | Update existing event           | Yes           |
| DELETE | /events/\:id          | Delete event                    | Yes           |
| POST   | /events/\:id/register | Register current user for event | Yes           |
| DELETE | /events/\:id/register | Cancel event registration       | Yes           |

#### Sample JSON for Creating/Updating an Event:

```json
{
  "name": "Go Meetup",
  "description": "A meetup to discuss Go projects",
  "location": "Remote",
  "datetime": "2025-07-25T18:00:00Z"
}
```

## 🛠️ Getting Started

### Prerequisites

* Go 1.18 or later

### Clone and Run

```bash
git clone https://github.com/Niranjini-Kathiravan/go-rest-api.git
cd go-rest-api-v2
go mod tidy
go run main.go
```

Server runs at: `http://localhost:8080`

---


## 🤝 Acknowledgements

* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [SQLite Driver (mattn/go-sqlite3)](https://github.com/mattn/go-sqlite3)
* [bcrypt for Go](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

✅ This project is free and open-source. Contributions and feedback are welcome!
