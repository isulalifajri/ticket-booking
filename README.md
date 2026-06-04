# membuat sistem booking tiket

tech stack: `golang Version go 1.25.1` and db: `postgreSQL`

---

## Start Project

**Structur Folder Planning**

```
ticket-booking/

├── cmd/
│   └── api/
│       └── main.go
│
├── configs/
│   └── config.go
│
├── internal/
│
│   ├── database/
│   │   └── postgres.go
│   │
│   ├── middleware/
│   │   ├── auth.go
│   │   └── cors.go
│   │
│   ├── models/
│   │   ├── user.go
│   │   ├── movie.go
│   │   ├── schedule.go
│   │   ├── seat.go
│   │   ├── booking.go
│   │   └── payment.go
│   │
│   ├── repositories/
│   │   ├── user_repository.go
│   │   ├── movie_repository.go
│   │   └── booking_repository.go
│   │
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── booking_service.go
│   │   ├── payment_service.go
│   │   └── seat_service.go
│   │
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── movie_handler.go
│   │   ├── booking_handler.go
│   │   └── payment_handler.go
│   │
│   └── routes/
│       └── routes.go
│
├── migrations/
│
├── .env
├── docker-compose.yml
├── go.mod
└── README.md
```

---

jalankan : `go mod init ticket-booking` -> membuat folder go


**step 2: buat struktur folder dasar**

```
ticket-booking/

├── cmd/
│   └── api/
│       └── main.go
│
├── configs/
│
├── internal/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   └── services/
│
├── migrations/
│
├── .env
├── .gitignore
├── go.mod
└── README.md
```

**step 3: install dependency awal**

```
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

**step 4: setup env**
```
APP_PORT=3000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ticket_booking
DB_SSLMODE=disable
```

**step 4: buat file model untuk strucktur database**