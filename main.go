package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/Shreyas-Prabhu/GoPostgreSQLMux/docs"
	"github.com/Shreyas-Prabhu/GoPostgreSQLMux/router"
	"github.com/joho/godotenv"
)

// @title EmployeeAssetManagement
// @version 1.0
// @description This is a Employee Asset management intended to use internally only. Not a public facing api.
// @termsOfService http://swagger.io/terms/
// @host localhost:4000
// @BasePath /

func main() {
	route := router.CreateRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	log.Println("Starting at", port, "...")
	log.Fatal(http.ListenAndServe(port, route))
}
