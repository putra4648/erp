# ERP System

A modern ERP (Enterprise Resource Planning) system built with high-performance technologies, focused on Inventory Management.

## 🛠 Tech Stack

### Backend

- **Language**: [Go](https://go.dev/) (1.25+)
- **Framework**: [Fiber v2](https://gofiber.io/)
- **Database ORM**: [GORM](https://gorm.io/) (PostgreSQL)
- **Authentication/Authorization**:
  - [Auth0](https://auth0.com/) (OIDC)
  - [Goidc](https://github.com/coreos/go-oidc)
- **Dependency Injection**: [Dig](https://github.com/uber-go/dig)
- **Logging**: [Zap](https://github.com/uber-go/zap)
- **Dev Tooling**: [Air](https://github.com/air-verse/air) (Live reload)

### Frontend

- **Framework**: [Nuxt 4](https://nuxt.com/) (Vue 3)
- **UI Library**: [@nuxt/ui](https://ui.nuxt.com/) (Tailwind CSS 4)
- **Authentication**: [nuxt-auth-utils](https://github.com/atinux/nuxt-auth-utils)
- **Package Manager**: [pnpm](https://pnpm.io/)

---

## ✨ Features

- **Inventory Management**:
  - **Stock Levels**: Real-time monitoring of product quantities across warehouses.
  - **Stock Movements**: Track inbound, outbound, and internal transfers.
  - **Stock Adjustments**: Manual inventory corrections with reason tracking.
  - **Low Stock Alerts**: Intelligent dashboard alerts for replenishment.
- **Master Data Management**:
  - Manage **Products**, **Categories**, **UOMs**, **Suppliers**, and **Warehouses**.
- **Modern UI/UX**:
  - Full-featured Dashboard with quick statistics and recent activity.
  - Responsive design with dark/light mode support.
  - Data-driven components based on `@nuxt/ui`.
- **RBAC & Security**:
  - Secure API routes with OIDC-based middleware.
  - Role-based navigation and action control.

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (version 1.25 or higher)
- [Node.js](https://nodejs.org/) (LTS recommended)
- [pnpm](https://pnpm.io/installation)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/) (Optional, for containerized setup)

### 🔐 Auth0 Setup

This project uses Auth0 for Authentication. You need to configure an Application and an API:

1. **Regular Web Application**:
   - Create a Regular Web Application in Auth0.
   - Allowed Callback URLs: `http://localhost:3000/auth/auth0`.
   - Allowed Logout URLs: `http://localhost:3000`.
2. **API**:
   - Create an API in Auth0.
   - Identifier (Audience): e.g., `http://localhost:8080`.

### ⚙️ Environment Variables

#### Backend (`backend/.env`)

```env
AUTH0_CLIENT_ID=your-auth0-client-id
AUTH0_CLIENT_SECRET=your-auth0-client-secret
AUTH0_DOMAIN=your-auth0-domain
AUTH0_AUDIENCE=http://localhost:8080
DB_DSN=host=localhost user=postgres password=password dbname=erp port=5432 sslmode=disable
PORT=8080
```

#### Frontend (`frontend/.env`)

```env
NUXT_OAUTH_AUTH0_CLIENT_ID=your-auth0-client-id
NUXT_OAUTH_AUTH0_CLIENT_SECRET=your-auth0-client-secret
NUXT_OAUTH_AUTH0_DOMAIN=your-auth0-domain
NUXT_OAUTH_AUTH0_AUDIENCE=http://localhost:8080
NUXT_OAUTH_AUTH0_REDIRECT_URL=http://localhost:3000/auth/auth0
NUXT_SESSION_PASSWORD=at-least-32-character-long-password
NUXT_SERVER_URL=http://localhost:8080
```

### Local Development

#### Backend Setup

1. Navigate to the backend directory:

   ```bash
   cd backend
   ```

2. Install Go dependencies:

   ```bash
   go mod download
   ```

3. Run with hot reload (using Air):

   ```bash
   go tool air
   ```

#### Frontend Setup

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

---

## 🐳 Docker Setup

You can run the entire system using Docker Compose:

1. Create a `.env` file in the root directory with the corresponding environment variables.
2. Run:

   ```bash
   docker-compose up -d --build
   ```

---

## 📂 Project Structure

```text
.
├── backend/            # Go Backend API (Fiber + GORM)
│   ├── cmd/            # Entry point
│   ├── internal/       # Core business logic (Hexagonal-lite)
│   ├── routes/         # API Route definitions
│   └── configs/        # Middleware, DI, and Auth configs
├── frontend/           # Nuxt 4 Frontend
│   ├── app/            # Vue components, pages, and layouts
│   ├── server/         # Nuxt server utilities
│   └── public/         # Static assets
└── migrations/         # Database migrations
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
