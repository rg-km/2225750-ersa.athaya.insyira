package api

import (
	"encoding/json"

	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	products, err := api.productsRepo.SelectAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder := json.NewEncoder(w)
		encoder.Encode(ProductListErrorResponse{Error: err.Error()})
		return
	}

	// looping products dan ubah kedalam bentuk json
	var productsJson []Product
	for _, product := range products {
		productsJson = append(productsJson, Product{
			Name:     product.ProductName,
			Price:    product.Price,
			Category: product.Category,
		})
	}

	// kirim response
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(ProductListSuccessResponse{Products: productsJson})

	// fmt.Println(products)

	// encoder.Encode(ProductListSuccessResponse{Products: []Product{}}) // TODO: replace this
}
