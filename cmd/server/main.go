package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	fiberzap "github.com/gofiber/contrib/fiberzap/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gitams17/ainyx-assignment/config"
	db "github.com/gitams17/ainyx-assignment/db/sqlc"
	"github.com/gitams17/ainyx-assignment/internal/handler"
	"github.com/gitams17/ainyx-assignment/internal/logger"
	"github.com/gitams17/ainyx-assignment/internal/service"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Initialize Zap Logger
	logger.InitLogger()
	defer logger.Log.Sync()

	// 3. Connect to Database (PGX Pool)
	connPool, err := pgxpool.New(context.Background(), cfg.DBUrl)
	if err!= nil {
		log.Fatal("Cannot connect to db:", err)
	}
	defer connPool.Close()

	// 4. Initialize Layers
	queries := db.New(connPool)
	// Note: We pass *db.Queries directly because sqlc's generated code 
	// satisfies the interface requirements of the service layer if structured correctly.
	// For stricter interface compliance, an adapter might be used.
	userService := service.NewUserService(queries) 
	userHandler := handler.NewUserHandler(userService)

	// 5. Setup Fiber
	app := fiber.New(fiber.Config{
		AppName: "User Age API",
	})

	// 6. Middleware
	app.Use(requestid.New()) // Injects X-Request-ID
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.Log,
	}))

	// 7. Routes
	api := app.Group("/users")
	api.Post("/", userHandler.CreateUser)
	api.Get("/", userHandler.ListUsers)
	api.Get("/:id", userHandler.GetUser)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)

	// 8. Start Server
	log.Fatal(app.Listen(cfg.Port))
}