package product

import (
	productrepo "ecommerse/repo"
	"ecommerse/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProd productrepo.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProd)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "PLease give me a valid json", 400)
		return
	}
	prod, err := h.productRepo.Store(newProd)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, prod, 201)
}
