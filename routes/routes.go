package routes

import (
	handler "github.com/Salauddin958/go-mux-product-apis/handler/http"
	mux "github.com/gorilla/mux"
)

func SetUpRoutes(h *handler.Product) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/products", h.Fetch).Methods("GET")
	r.HandleFunc("/product", h.Create).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", h.GetByID).Methods("GET")
	r.HandleFunc("/product/{id:[0-9]+}", h.Update).Methods("PUT")
	r.HandleFunc("/product/{id:[0-9]+}", h.Delete).Methods("DELETE")
	return r
}
