package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klferreira/hash-challenge/product-service/pkg/product"
)

func Fetch(service product.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		products, err := service.Fetch(r.Context())
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
			return
		}

		if err := json.NewEncoder(w).Encode(products); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong"))
		}
	})
}

func GetProductHandlers(r *mux.Router, service product.Service) {
	r.Handle("/v1/products", Fetch(service)).Methods(http.MethodGet).Name("FetchProducts")
}
