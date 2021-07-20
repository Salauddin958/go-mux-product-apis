package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Salauddin958/go-mux-product-apis/driver"
	ph "github.com/Salauddin958/go-mux-product-apis/handler/http"
	route "github.com/Salauddin958/go-mux-product-apis/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	serverPort := os.Getenv("SERVER_PORT")
	fmt.Println("dbName ::", dbName)
	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	handler := ph.NewProductHandler(connection)
	r := route.SetUpRoutes(handler)
	fmt.Println("Server listening at :", serverPort)
	http.ListenAndServe(":"+serverPort, r)
}
