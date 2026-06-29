package product

import (
	"ecommerse/database"
	"ecommerse/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProd database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProd)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "PLease give me a valid json", 400)
		return
	}
	prod := database.Store(newProd)
	util.SendData(w, prod, 201)
}
