package product

import (
	"ecommerse/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductsById),
		),
	)
}
