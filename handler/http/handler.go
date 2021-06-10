package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Salauddin958/go-mux-product-apis/driver"
	models "github.com/Salauddin958/go-mux-product-apis/models"
	repository "github.com/Salauddin958/go-mux-product-apis/repository"
	product "github.com/Salauddin958/go-mux-product-apis/repository/product"
	mux "github.com/gorilla/mux"
)

// NewProductHandler ....
func NewProductHandler(db *driver.DB) *Product {
	return &Product{
		repo: product.NewSQLProductRepo(db.SQL),
	}
}

// Product ....
type Product struct {
	repo repository.ProductRepo
}

// Fetch all Product data
func (b *Product) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, err := b.repo.Fetch(r.Context(), 5)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, payload)
}

// Create new Product
func (b *Product) Create(w http.ResponseWriter, r *http.Request) {
	book := models.Product{}
	json.NewDecoder(r.Body).Decode(&book)

	newID, err := b.repo.Create(r.Context(), &book)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update Product
func (b *Product) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusMisdirectedRequest, err.Error())
		return
	}
	data := models.Product{ID: int(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := b.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a product details
func (b *Product) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	payload, err := b.repo.GetByID(r.Context(), int64(id))
	if err != nil {
		respondWithError(w, http.StatusNoContent, err.Error())
		return
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete Product
func (b *Product) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = b.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
