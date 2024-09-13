package main

import (
	"backend/routes"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

func main() {
	// Connect to database
	fmt.Println("Try to connect to DB")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Agora_Project")

	if err != nil {
		fmt.Println("Error connecting")
		return
	} else {
		fmt.Println("I AM IN!!!! THE MAINFRAME")
	}
	defer db.Close()

	// Initialize router and handlers
	router := http.NewServeMux()
	router.HandleFunc("/allList", routes.HandleQueryAll(db))

	// Setup CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow only frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	// Start server
	server := http.Server{
		Addr:    ":8080",
		Handler: corsHandler,
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
