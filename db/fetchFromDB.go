package db

import "fmt"

// These can be called inside services
func FetchFromDB() {
	fmt.Println("Query DB to fetch data")
}