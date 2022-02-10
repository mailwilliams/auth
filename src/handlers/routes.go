package handlers

//	ConfigureRoutes is the method that is used to direct incoming api requests
//	to the appropriate endpoint and its associated method
func (handler *Handler) ConfigureRoutes() {
	app := handler.App
	app.Get("/", handler.Hello)

	//	creating groups for prefixes
	//	/api
	api := app.Group("api")
	api.Post("register", handler.Register)
	api.Post("login", handler.Login)
	api.Delete("logout", handler.Logout)

	//	authenticated API paths
	authenticatedUser := api.Use(handler.IsAuthenticated)
	authenticatedUser.Get("/users", handler.ListUsers)
	authenticatedUser.Put("/users/me", handler.UpdateInfo)
	authenticatedUser.Put("/users/me/password", handler.UpdatePassword)
	authenticatedUser.Put("/users/me/wallet", handler.UpdateWalletAddress)

	//	/api/admin
	_ = api.Group("admin")

}
