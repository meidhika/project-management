# 🚀 Project Management API (Go + Fiber)

A simple Project Management REST API built using **Golang** and **Fiber**.

## 📦 Tech Stack

- Golang
- Fiber (Web Framework)
- GORM / SQL (optional)
- PostgreSQL / MySQL (optional)
- JWT Authentication (optional)

---

## 📁 Project Structure


├── cmd/ # Entry point aplikasi
├── config/ # Konfigurasi (DB, env, dll)
├── controllers/ # Handler / controller
├── services/ # Business logic
├── repositories/ # Database access layer
├── models/ # Struct model
├── routes/ # Routing
├── middleware/ # Middleware (auth, logging, dll)
├── utils/ # Helper functions
├── .env
├── go.mod
└── main.go


---

## ⚙️ Installation

### 1. Clone Repository

```bash
git clone https://github.com/your-username/project-management.git
cd project-management
2. Install Dependencies
go mod tidy
3. Setup Environment

Buat file .env:

APP_PORT=3000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=
DB_NAME=project_management

JWT_SECRET=your_secret_key

▶️ Run Application
go run main.go

atau jika pakai air (hot reload):

air
🔐 Authentication

Menggunakan JWT Token:

Authorization: Bearer <token>
🧪 Testing
go test ./...
📦 Build
go build -o app
./app
📝 Notes
Pastikan database sudah berjalan sebelum menjalankan aplikasi
Gunakan migration tool jika diperlukan (golang-migrate / gorm auto migrate)