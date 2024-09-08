package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func main() {
	fmt.Println("Try to connect to DB")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/agora")

	if err != nil {
		fmt.Println("Error connecting")
		return
	} else {
		fmt.Println("I AM IN!!!! THE MAINFRAME")
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println("The Connection died")
		return
	} else {
		fmt.Println("STILL IN BABY")
	}

} // end of main
