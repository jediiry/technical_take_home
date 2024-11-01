package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"technical_take_home/internal/database"
	"technical_take_home/internal/handler"
	"technical_take_home/internal/routes"
)

func main() {
	port := getPort()
	dataStore := database.NewDataStore()
	handler := handler.NewKeyValueHandler(dataStore)
	router := routes.RegisterRoutes(handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func getPort() int {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		return 8080
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}
	return port
}