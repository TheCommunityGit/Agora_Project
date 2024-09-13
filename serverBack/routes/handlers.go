package routes

import (
	"backend/queries"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleQueryAll accepts a database pointer and returns an http.HandlerFunc
func HandleQueryAll(db *sql.DB) http.HandlerFunc {
	fmt.Println("HANDLER TIME??")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("FETCHING TIME??")
		// Fetch listings from the database using queries.Listings
		listings := queries.Listings(db)
		if listings == nil {
			http.Error(w, "Failed to fetch listings", http.StatusInternalServerError)
			return
		}

		// Set response header and return listings as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(listings)
	}
}
