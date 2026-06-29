package cmd

import (
	"ecommerse/config"
	"ecommerse/rest"
	"ecommerse/rest/handlers/product"
	"ecommerse/rest/handlers/review"
	"ecommerse/rest/handlers/user"
	"ecommerse/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middleware := middlewares.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middleware)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(productHandler, userHandler, reviewHandler)
	server.Start(*cnf)
}
