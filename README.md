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
│
├── internal/
│
│   ├── database/
│   │   ├── postgres.go
│   │   └── migrate.go
│   │
│   ├── middleware/
│   │   ├── auth.go
│   │   └── cors.go
│   │
│   ├── models/
│   │   ├── user.go
│   │   ├── movie.go
│   │   ├── studio.go
│   │   ├── studio_seat.go
│   │   ├── schedule.go
│   │   ├── booking.go
│   │   ├── booking_seat.go
│   │   └── payment.go
│   │
│   ├── repositories/
│   │   ├── user_repository.go
│   │   ├── movie_repository.go
│   │   ├── studio_repository.go
│   │   ├── schedule_repository.go
│   │   ├── booking_repository.go
│   │   └── payment_repository.go
│   │
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── movie_service.go
│   │   ├── studio_service.go
│   │   ├── schedule_service.go
│   │   ├── booking_service.go
│   │   └── payment_service.go
│   │
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── movie_handler.go
│   │   ├── studio_handler.go
│   │   ├── schedule_handler.go
│   │   ├── booking_handler.go
│   │   └── payment_handler.go
│   │
│   └── routes/
│       └── routes.go
│
├── migrations/
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
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

```
user.go
movie.go
studio.go
studio_seat.go
schedule.go
booking.go
booking_seat.go
payment.go
```

setelah selesai, kita jalankan: `go run cmd/api/main.go`

ok, setelah table berhasil di buat, selanjutanya buat file ini `routes.go` disini: `internal/routes` isi filenya seperti ini:

```
package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

}
```

kemudian isi main.go jadi seperti ini:

```
package main

import (
	"log"

	"ticket-booking/internal/database"
	"ticket-booking/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed load env")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
```

jalankan lagi: `go run cmd/api/main.go`

browser: `http://localhost:3000/health`