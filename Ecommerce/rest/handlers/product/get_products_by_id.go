package product

import (
	"ecommerse/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductsById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		util.SendError(w, 404, "Product Not Found")
		return
	}

	util.SendData(w, product, 200)
}
