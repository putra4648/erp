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

## ✨ Features

- **Inventory Management**:
  - **Stock Levels**: Real-time monitoring of product quantities across warehouses.
  - **Stock Movements**: Track inbound, outbound, and internal transfers.
  - **Stock Adjustments**: Manual inventory corrections with supervisor approval workflow.
  - **Stock Transactions**: Detailed audit trail of all inventory changes.
- **Master Data Management** (Admin only):
  - Manage Products, Categories, UOMs, Suppliers, and Warehouses.
- **Role-Based Access Control (RBAC)**:
  - Dynamic navigation menu based on Keycloak groups (`admin`, `staff`, `inventory`).
  - Secure API routes and frontend middleware protection.
- **Modern UI/UX**:
  - Unified layout with responsive sidebar and mobile drawer.
  - Dark/Light mode support.
  - Reusable data-driven components (Badges, Tables, Forms).

---

## 🚀 Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (version 1.25 or higher)
- [Node.js](https://nodejs.org/) (LTS recommended)
- [pnpm](https://pnpm.io/installation)
- [PostgreSQL](https://www.postgresql.org/)
- [Keycloak](https://www.keycloak.org/)

### 🔐 Keycloak Setup

This project uses Keycloak for Authentication and Authorization. You need to configure a Realm and Client:

1. **Realm**: Create a realm named `erp`.
2. **Client**:
   - Create a client named `erp`.
   - Client Protocol: `openid-connect`.
   - Access Type: `confidential` (or `public` depending on your flow, but `confidential` is recommended for SSR).
   - Valid Redirect URIs: `http://localhost:3000/*`.
   - Web Origins: `http://localhost:3000`.
3. **Roles & Groups**:
   - Create groups: `admin`, `staff`, `inventory`.
   - Add your user to these groups to test access control.

### ⚙️ Environment Variables

#### Backend (`backend/.env`)

```env
KEYCLOAK_URL=http://localhost:8081
KEYCLOAK_REALM_NAME=erp
KEYCLOAK_CLIENT_ID=erp
KEYCLOAK_CLIENT_SECRET=your-client-secret
DB_DSN=host=localhost user=postgres password=password dbname=erp port=5432 sslmode=disable
PORT=8080
```

#### Frontend (`frontend/.env`)

```env
KEYCLOAK_URL=http://localhost:8081
KEYCLOAK_ISSUER=http://localhost:8081/realms/erp
KEYCLOAK_CLIENT_ID=erp
KEYCLOAK_CLIENT_SECRET=your-client-secret
AUTH_SECRET=your-nuxt-auth-secret
```

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

```text
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
