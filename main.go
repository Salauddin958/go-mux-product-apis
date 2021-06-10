package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Salauddin958/go-mux-product-apis/driver"
	ph "github.com/Salauddin958/go-mux-product-apis/handler/http"
	mux "github.com/gorilla/mux"
)

func main() {

	dbName := "productdb"
	dbPass := "root"
	dbHost := "localhost"
	dbPort := "3306"

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := mux.NewRouter()
	handler := ph.NewProductHandler(connection)
	r.HandleFunc("/products", handler.Fetch).Methods("GET")
	r.HandleFunc("/product", handler.Create).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", handler.GetByID).Methods("GET")
	r.HandleFunc("/product/{id:[0-9]+}", handler.Update).Methods("PUT")
	r.HandleFunc("/product/{id:[0-9]+}", handler.Delete).Methods("DELETE")

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}
