package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	middlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Use(middlewares ...Middleware) {  
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *Manager) WrapMux(mux *http.ServeMux) http.Handler {
	return m.With(mux)
}

func (m *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	all := append([]Middleware{}, m.middlewares...)
	all = append(all, middlewares...)

	for i := len(all) - 1; i >= 0; i-- {
		handler = all[i](handler)
	}

	return handler
}
