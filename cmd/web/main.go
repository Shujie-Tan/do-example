package main

import (
	"database/sql"
	"example/user"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=root dbname=mydb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userHandler := user.Wire(db)
	http.Handle("/user", userHandler.FetchByUsername())
	fmt.Printf("Try run server at :%d\n", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
