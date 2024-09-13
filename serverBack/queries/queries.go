package queries

import (
	"database/sql"
	"fmt"
)

type Listing struct {
	Email       string `json:"email"`
	ListID      int    `json:"listID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Active      bool   `json:"active"`
	Category    string `json:"category"`
	Rarity      string `json:"rarity"`
}

func Listings(db *sql.DB) []Listing {
	fmt.Println("ITS QUERING TIME BABYYYYYYY!!!")
	rows, err := db.Query("SELECT * FROM listings")
	if err != nil {
		fmt.Println("Query failed at Listings:", err)
		return nil
	}
	defer rows.Close()

	var listings []Listing

	for rows.Next() {
		var currListing Listing
		if err := rows.Scan(&currListing.Email, &currListing.ListID,
			&currListing.Title, &currListing.Description, &currListing.Img,
			&currListing.Active, &currListing.Category, &currListing.Rarity); err != nil {
			fmt.Println("Row scan failed:", err)
			return listings
		}
		listings = append(listings, currListing)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Rows error:", err)
		return listings
	}

	return listings
}
