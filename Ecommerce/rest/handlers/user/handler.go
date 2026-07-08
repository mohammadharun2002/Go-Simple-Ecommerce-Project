package user

import userrepo "ecommerse/repo"

type Handler struct {
	userRepo userrepo.UserRepo
}

func NewHandler() *Handler {
	return &Handler{
		userRepo: userrepo.NewUserRepo(),
	}
}
