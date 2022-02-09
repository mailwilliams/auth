package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mailwilliams/auth/src/database"
	"github.com/mailwilliams/auth/src/handlers"
)

var (
	app = &fiber.App{}
	ctx = context.Background()
)

func main() {

	//	on startup, create new connection to GORM to begin automatic migrations
	migrationsDB, err := database.NewGORMConnection()
	if err != nil {
		panic(err)
	}

	//	run the migrate function to sync database with most recent model.
	//	To see existing models, go to /models directory
	if err := database.AutoMigrate(migrationsDB); err != nil {
		panic(err)
	}

	//	connect to MySQL database
	db, err := database.NewSQLConnection(ctx)
	if err != nil {
		panic(err)
	}

	//	connect to Redis cache
	cache := database.ConnectCache(ctx)

	//	ping database and cache, assign app to handler
	handler, dbErr, cacheErr := handlers.NewHandler(ctx, db, cache, app)
	if dbErr != nil {
		panic(dbErr)
	}
	if cacheErr != nil {
		panic(cacheErr)
	}

	//	initializes new app, with optional config parameters
	//	leaving fiber.Config{} empty because we may need to add something
	handler.CreateApp(fiber.Config{})
	handler.CreateRouter(cors.New(cors.Config{AllowCredentials: true}))

	//	map endpoints to their intended methods
	handler.ConfigureRoutes()

	//	begin to listen for requests
	if err := handler.App.Listen(":8000"); err != nil {
		panic(err)
	}
}
