package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

type listing struct {
	email       string
	listID      int
	title       string
	description string
	img         string
	active      bool
	category    string
	rarity      string
}

func Listings(db *sql.DB) []listing {
	rows, err := db.Query("select * from listings")
	if err != nil {
		fmt.Println("query failed at Listings")
		return nil
	}
	defer rows.Close()
	var listings []listing

	for rows.Next() {
		var currListing listing

		if err := rows.Scan(&currListing.email, &currListing.listID,
			&currListing.title, &currListing.description,
			&currListing.img, &currListing.active,
			&currListing.category, &currListing.rarity); err != nil {
			return listings
		}
		listings = append(listings, currListing)

	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return listings
	}
	return listings
}

func UserListings(db *sql.DB, email string) []listing {
	rows, err := db.Query("select * from listings where email = ?", email)
	if err != nil {
		fmt.Println("query failed at UserListings")
		return nil
	}
	defer rows.Close()
	var listings []listing

	for rows.Next() {
		var currListing listing

		if err := rows.Scan(&currListing.email, &currListing.listID,
			&currListing.title, &currListing.description,
			&currListing.img, &currListing.active,
			&currListing.category, &currListing.rarity); err != nil {
			return listings
		}
		listings = append(listings, currListing)

	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return listings
	}
	return listings
}
