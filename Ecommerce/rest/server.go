package rest

import (
	"ecommerse/config"
	"ecommerse/rest/handlers/product"
	"ecommerse/rest/handlers/review"
	"ecommerse/rest/handlers/user"
	"ecommerse/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler *review.Handler) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server running on :", config.GetConfig().HttpPort)

	addr := ":" + strconv.FormatInt(cnf.HttpPort, 10)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
