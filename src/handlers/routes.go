package handlers

import "github.com/mailwilliams/auth/src/middlewares"

func (handler *Handler) ConfigureRoutes() {
	app := handler.App
	app.Get("/", handler.Hello)

	//	creating groups for prefixes
	//	/api
	api := app.Group("api")
	api.Post("register", handler.Register)
	//api.Post("login", controllers.Login)

	//	authenticated API paths
	_ = api.Use(middlewares.IsAuthenticated)

	//	/api/admin
	_ = api.Group("admin")

}

//func SetUp(app *fiber.App) {
//
//	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
//	adminAuthenticated.Get("user", controllers.User)
//	adminAuthenticated.Post("logout", controllers.Logout)
//	adminAuthenticated.Put("users/info", controllers.UpdateInfo)
//	adminAuthenticated.Put("users/password", controllers.UpdatePassword)
//	adminAuthenticated.Get("ambassadors", controllers.Ambassadors)
//	adminAuthenticated.Get("products", controllers.Products)
//	adminAuthenticated.Post("products", controllers.CreateProduct)
//	adminAuthenticated.Get("products/:id", controllers.GetProduct)
//	adminAuthenticated.Put("products/:id", controllers.UpdateProduct)
//	adminAuthenticated.Delete("products/:id", controllers.DeleteProduct)
//	adminAuthenticated.Get("users/:id/links", controllers.Link)
//	adminAuthenticated.Get("orders", controllers.Orders)
//
//	ambassador := api.Group("ambassador")
//	ambassador.Post("register", controllers.Register)
//	ambassador.Post("login", controllers.Login)
//	ambassador.Get("products/frontend", controllers.ProductsFrontend)
//	ambassador.Get("products/backend", controllers.ProductsBackend)
//
//	ambassadorAuthenticated := admin.Use(middlewares.IsAuthenticated)
//	ambassadorAuthenticated.Get("user", controllers.User)
//	ambassadorAuthenticated.Post("logout", controllers.Logout)
//	ambassadorAuthenticated.Put("users/info", controllers.UpdateInfo)
//	ambassadorAuthenticated.Put("users/password", controllers.UpdatePassword)
//	ambassadorAuthenticated.Post("links", controllers.CreateLink)
//	ambassadorAuthenticated.Get("stats", controllers.Stats)
//	ambassadorAuthenticated.Get("rankings", controllers.Rankings)
//
//}
