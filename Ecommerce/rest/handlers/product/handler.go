package product

import (
	repo "ecommerse/repo"
	"ecommerse/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middlewares *middlewares.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: repo.NewProductRepo(),
	}
}
