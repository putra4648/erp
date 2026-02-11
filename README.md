# ERP System

A modern ERP (Enterprise Resource Planning) system built with high-performance technologies.

## 🛠 Tech Stack

### Backend
- **Language**: [Go](https://go.dev/) (1.25+)
- **Framework**: [Fiber v2](https://gofiber.io/)
- **Database ORM**: [GORM](https://gorm.io/) (PostgreSQL)
- **Authentication/Authorization**: 
  - [Casbin](https://casbin.org/) (RBAC)
  - [Gocloak](https://github.com/Nerzal/gocloak) / [OIDC](https://github.com/coreos/go-oidc)
- **Dependency Injection**: [Dig](https://github.com/uber-go/dig)
- **Logging**: [Zap](https://github.com/uber-go/zap)
- **Dev Tooling**: [Air](https://github.com/air-verse/air) (Live reload)

### Frontend
- **Framework**: [Nuxt 4](https://nuxt.com/) (Vue 3)
- **Styling**: [Tailwind CSS](https://tailwindcss.com/)
- **Authentication**: [Nuxt Auth](https://sidebase.io/nuxt-auth)
- **Package Manager**: [pnpm](https://pnpm.io/)

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:
- [Go](https://go.dev/dl/) (version 1.25 or higher)
- [Node.js](https://nodejs.org/) (LTS recommended)
- [pnpm](https://pnpm.io/installation)
- [PostgreSQL](https://www.postgresql.org/)

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables:
   Copy the example environment file and update valid configurations (Database credentials, etc.).
   ```bash
   cp .env.example .env
   ```

4. Run the development server (with hot reload):
   ```bash
   go tool air
   ```
   Or run normally:
   ```bash
   go run cmd/main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   pnpm install
   ```

3. Run the development server:
   ```bash
   pnpm dev
   ```
   The application will be accessible at `http://localhost:3000`.

---

## 📂 Project Structure

```
.
├── backend/            # Go Backend API
│   ├── cmd/            # Entry point
│   ├── configs/        # Configuration files
│   ├── internal/       # Internal application logic (Handlers, Services, Repositories)
│   ├── routes/         # API Routing definition
│   └── ...
├── frontend/           # Nuxt Frontend Application
│   ├── app/            # App components and pages
│   ├── server/         # Server-side routes
│   └── ...
└── README.md           # Project documentation
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request