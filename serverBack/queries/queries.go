package main 

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)


func Listings(db *sql.DB) []string{
	var result []string
	query, err := db.Query("select * from listings")
	if err != nil{
		fmt.Println("query failed")
	}else{
		result = append(result,query)
	}
	return result

}
func MyListings(db *sql.DB, email string) []string{
	var result  []string
	query, err := db.Query(fmt.Sprintf("select * from listings where email = %s",email))
	if err != nil{
		fmt.Println("query failed")
	}else{
		result = append(result,query)
	}
	return result


}
func Chat(db *sql.DB, email string) []string{
	var result  []string
	query, err := db.Query(fmt.Sprintf("select * from chats where users like \" % %s % \"",email))
	if err != nil{
		fmt.Println("query failed")
	}else{
		or rows.Next() {
			var 
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		    return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
		for _, response := range query {
			result = append(result,string(response))
		}
	}
	return result


}
func Messages(db *sql.DB, cID string) []string{
	var result  []string
	query, err := db.Query(fmt.Sprintf("select * from messages where cID = %s",cID))
	if err != nil{
		fmt.Println("query failed")
	}else{
		result = append(result,query)
	}
	return result
}
