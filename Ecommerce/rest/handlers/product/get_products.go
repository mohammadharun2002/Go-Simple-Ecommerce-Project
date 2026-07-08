package product

import (
	"ecommerse/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, h.productRepo.List(), 200)
}
