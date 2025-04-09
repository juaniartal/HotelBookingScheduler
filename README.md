# Booking Scheduler

This is a personal project built with the main goal of learning as much as possible about **Golang**. It’s a hotel booking scheduler app with a backend built in Go and a simple frontend using React (JSX + NPM), mainly to test the backend functionality. The backend connects to a **PostgreSQL** database to manage booking information.

---

## 🚀 Project Goal

This project was created as a personal challenge to:

- Learn how to build and structure a professional backend using Go.
- Build and expose REST APIs using Golang.
- Integrate a PostgreSQL relational database.
- Quickly test backend functionality with a minimal React frontend.

---

## ⚙️ Technologies Used

### Backend
- 🐹 **Golang**
- 📦 Libraries such as `net/http`, `mux`, and `sqlx`
- 🐘 **PostgreSQL** database
- ✅ Unit testing

### Frontend
- ⚛️ **React** (JSX)
- 📦 NPM
- 🧪 Simple forms to test API endpoints

---

## 📁 Project Structure

```bash
HotelBookingScheduler/
├── backend/                # Golang backend source code
├── vite-project/           # React frontend (Vite-based)
├── docker-compose.yml      # Docker Compose configuration
├── .gitignore              # Git ignore rules
├── LICENSE                 # Project license
└── README.md               # This file
```

## SETUP

### Prerequisites

- Go installed on your system.
- A running and accessible PostgreSQL instance.

### Steps

- Clone the repository:
  
 ```bash
git clone https://github.com/juaniartal/HotelBookingScheduler.git
cd HotelBookingScheduler/backend
```

- Install dependencies:
 
```bash
go moD download
```

- Create a .env file:

```bash
DB_HOST=localhost
DB_PORT=your_port
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database_name
```

## CI-CD🔄

The project includes a GitHub Actions pipeline that runs on every push to the repository.
The pipeline is currently set to run backend tests, but it's still a work in progress and needs to be fully functional.

