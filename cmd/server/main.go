package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	db "github.com/sq1754/user-age-api/db/sqlc"
	"github.com/sq1754/user-age-api/internal/handler"
	"github.com/sq1754/user-age-api/internal/logger"
	"github.com/sq1754/user-age-api/internal/middleware"
	"github.com/sq1754/user-age-api/internal/repository"
	"github.com/sq1754/user-age-api/internal/routes"
)

func main() {
	// DB DSN
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/userdb?sslmode=disable"
	}

	// Connect database (database/sql)
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		log.Fatal(err)
	}

	// SQLC queries
	queries := db.New(dbConn)

	// Logger
	logg := logger.New()
	defer logg.Sync()

	// Fiber app
	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(logg))

	// Repository & Handler
	userRepo := repository.NewUserRepository(queries)
	userHandler := handler.NewUserHandler(userRepo)

	// Routes
	routes.RegisterUserRoutes(app, userHandler)

	logg.Info("server started on :3000")
	log.Fatal(app.Listen(":3000"))
}
