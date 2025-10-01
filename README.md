# 🏛️ Portal Budaya Kalimantan Tengah

API Backend untuk Portal Budaya Kalimantan Tengah - Platform digital untuk mengelola dan menyajikan informasi budaya, suku, dan konten warisan budaya Kalimantan Tengah.

## 📋 Deskripsi

Portal Budaya Kalteng adalah aplikasi backend RESTful API yang dibangun dengan Go (Golang) menggunakan framework Gin dan GORM sebagai ORM. Aplikasi ini menyediakan sistem manajemen konten budaya dengan fitur autentikasi JWT dan Basic Auth untuk admin.

## 🚀 Fitur Utama

- ✅ **Authentication & Authorization**

  - JWT Token untuk user authentication
  - Basic Auth untuk admin endpoints
  - User registration & login

- ✅ **Content Management**

  - CRUD operations untuk konten budaya
  - Kategori konten
  - Status draft/published
  - Upload image URL
  - Slug-based routing

- ✅ **Taxonomy Management**

  - Manajemen suku (Tribes)
  - Manajemen wilayah/region (Regions)
  - Relasi many-to-many dengan konten

- ✅ **About Page**
  - Halaman tentang aplikasi
  - Upsert functionality

## 🛠️ Tech Stack

- **Language:** Go 1.23
- **Web Framework:** Gin v1.9.1
- **ORM:** GORM v1.25.5
- **Database:** PostgreSQL
- **Authentication:** JWT (golang-jwt/jwt v5.2.0)
- **Password Hashing:** bcrypt (golang.org/x/crypto)
- **Environment Variables:** godotenv v1.5.1

## 📁 Struktur Project

```
portal-budaya-kalteng/
├── cmd/
│   ├── server/          # Main application entry point
│   │   └── main.go
│   └── migrate/         # Database migration tool
│       └── main.go
├── internal/
│   ├── config/          # Configuration management
│   │   └── config.go
│   ├── database/        # Database connection setup
│   │   └── db.go
│   ├── dto/             # Data Transfer Objects
│   │   ├── auth_dto.go
│   │   └── content_dto.go
│   ├── handlers/        # HTTP request handlers
│   │   ├── about_handler.go
│   │   ├── auth_handler.go
│   │   ├── category_handler.go
│   │   ├── content_handler.go
│   │   └── taxonomy_handler.go
│   ├── middlware/       # Middleware functions
│   │   ├── basic.go     # Basic Auth middleware
│   │   └── jwt.go       # JWT middleware
│   ├── models/          # Database models
│   │   ├── about.go
│   │   ├── category.go
│   │   ├── content.go
│   │   ├── region.go
│   │   ├── tribe.go
│   │   └── user.go
│   ├── routes/          # Route definitions
│   │   └── routes.go
│   └── util/            # Utility functions
│       ├── hash.go      # Password hashing
│       └── slug.go      # Slug generation
├── migrations/
│   └── 001_init.sql     # Initial database schema
├── bin/
│   └── server.exe       # Compiled binary
├── .env                 # Environment variables
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── migrate.ps1          # Migration script (PowerShell)
└── README.md
```

## 🔧 Instalasi

### Prerequisites

- Go 1.21 atau lebih tinggi
- PostgreSQL 12 atau lebih tinggi
- Git

### Langkah Instalasi

1. **Clone repository**

   ```bash
   git clone <repository-url>
   cd portal-budaya-kalteng
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Setup Database**

   - Buat database PostgreSQL baru:
     ```sql
     CREATE DATABASE db_portal_budaya_kalteng;
     ```

4. **Setup Environment Variables**

   Copy file `.env.example` ke `.env` atau buat file `.env` baru:

   ```env
   APP_ENV=development
   APP_PORT=8080

   # Basic Auth untuk endpoint admin
   BASIC_AUTH_USER=admin
   BASIC_AUTH_PASS=password

   # JWT Configuration
   JWT_SECRET=your-super-secret-jwt-key-here
   JWT_TTL_HOURS=24

   # Database Configuration
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASS=your-password
   DB_NAME=db_portal_budaya_kalteng
   DB_SSLMODE=disable
   ```

5. **Run Database Migration**

   ```bash
   # Menggunakan Go
   go run cmd/migrate/main.go

   # Atau menggunakan PowerShell script
   .\migrate.ps1
   ```

6. **Run Application**

   ```bash
   # Development mode
   go run cmd/server/main.go

   # Atau build dan run
   go build -o bin/server.exe ./cmd/server
   .\bin\server.exe
   ```

Server akan berjalan di `http://localhost:8080`

## 📡 API Endpoints

### Public Endpoints

#### Authentication

```http
POST /api/auth/register
POST /api/auth/login
```

#### Content (Read-Only)

```http
GET  /api/contents           # List semua konten
GET  /api/contents/:id       # Detail konten (by ID atau slug)
GET  /api/categories         # List kategori
GET  /api/tribes             # List suku
GET  /api/regions            # List wilayah
GET  /api/about              # Halaman about
```

### Admin Endpoints (Basic Auth Required)

Header: `Authorization: Basic base64(username:password)`

```http
POST /api/admin/categories   # Create kategori baru
POST /api/admin/tribes       # Create suku baru
POST /api/admin/regions      # Create region baru
POST /api/admin/about        # Upsert about page
POST /api/admin/contents     # Create konten baru
```

### Protected Endpoints (JWT Required)

Header: `Authorization: Bearer <jwt-token>`

```http
GET  /api/me/profile         # Get user profile
```

## 📝 Request/Response Examples

### Register User

**Request:**

```http
POST /api/auth/register
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "display_name": "John Doe",
  "password": "secretpassword"
}
```

**Response:**

```json
{
  "id": "uuid-here",
  "username": "johndoe"
}
```

### Login

**Request:**

```http
POST /api/auth/login
Content-Type: application/json

{
  "username_or_email": "johndoe",
  "password": "secretpassword"
}
```

**Response:**

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIs...",
  "token_type": "Bearer",
  "expires_in": 86400
}
```

### Create Content (Admin)

**Request:**

```http
POST /api/admin/contents
Authorization: Basic YWRtaW46cGFzc3dvcmQ=
Content-Type: application/json

{
  "title": "Tari Tambun dan Bungai",
  "summary": "Tarian tradisional suku Dayak",
  "body": "Deskripsi lengkap tentang tarian...",
  "category_id": "uuid-category",
  "tribe_ids": ["uuid-tribe-1"],
  "region_ids": ["uuid-region-1"],
  "status": "published",
  "image_url": "https://example.com/image.jpg"
}
```

**Response:**

```json
{
  "id": "uuid-content",
  "title": "Tari Tambun dan Bungai",
  "slug": "tari-tambun-dan-bungai",
  "status": "published",
  "created_at": "2025-10-01T12:00:00Z"
}
```

## 🗄️ Database Schema

### Users

- id (UUID, PK)
- username (unique)
- email (unique)
- display_name
- password_hash
- role (admin/member)
- bio
- created_at, updated_at

### Contents

- id (UUID, PK)
- title
- slug (unique)
- image_url
- summary
- body
- status (draft/published)
- published_at
- category_id (FK)
- author_id (FK)
- created_at, updated_at

### Categories

- id (UUID, PK)
- name (unique)
- slug (unique)
- description
- created_at, updated_at

### Tribes (Suku)

- id (UUID, PK)
- name
- slug (unique)
- description

### Regions (Wilayah)

- id (UUID, PK)
- name
- slug (unique)
- description

### About

- id (UUID, PK)
- title
- description
- updated_by (FK to users)
- created_at, updated_at

### Pivot Tables

- content_tribes (many-to-many)
- content_regions (many-to-many)

## 🔒 Security

- Password di-hash menggunakan bcrypt
- JWT untuk stateless authentication
- Basic Auth untuk admin endpoints
- Environment variables untuk sensitive data
- SQL injection prevention via GORM
- CORS configuration (dapat dikonfigurasi di middleware)

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -v ./internal/handlers
```

## 🚢 Deployment

### Build untuk Production

```bash
# Set mode production
export GIN_MODE=release

# Build binary
go build -o bin/server ./cmd/server

# Run
./bin/server
```

### Docker (Optional)

```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./server"]
```

## 📚 Dependencies

Lihat `go.mod` untuk daftar lengkap dependencies:

```go
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/golang-jwt/jwt/v5 v5.2.0
    github.com/joho/godotenv v1.5.1
    golang.org/x/crypto v0.18.0
    gorm.io/driver/postgres v1.5.4
    gorm.io/gorm v1.25.5
)
```

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

[MIT License](LICENSE)

## 👥 Authors

- **Your Name** - _Initial work_

## 🙏 Acknowledgments

- Gin Web Framework
- GORM ORM
- PostgreSQL Database
- Go Community

## 📞 Contact

- Email: your.email@example.com
- Website: https://your-website.com
- GitHub: [@yourusername](https://github.com/yourusername)

---

**Made with ❤️ for preserving Kalimantan Tengah's cultural heritage**
