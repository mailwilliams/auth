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
)

func main() {

	ctx := context.Background()
	//	baseHandler is connecting to databases and assigning reference to app
	db, err := database.NewSQLConnection(ctx)
	if err != nil {
		panic(err)
	}

	cache := database.ConnectCache(ctx)

	baseHandler, dbErr, cacheErr := handlers.NewHandler(ctx, db, cache, app)
	if dbErr != nil {
		panic(dbErr.Error())
	}
	if cacheErr != nil {
		panic(cacheErr.Error())
	}

	//	initializes new app, with optional config parameters
	//	leaving fiber.Config{} empty because we may need to add something
	baseHandler.CreateApp(fiber.Config{})
	baseHandler.CreateRouter(cors.New(cors.Config{AllowCredentials: true}))
	baseHandler.ConfigureRoutes()
	if err := baseHandler.App.Listen(":8000"); err != nil {
		panic(err)
	}
}
