package main

import (
	"database/sql"
	"example/domain"
	"example/user"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/samber/do"
)

func main() {
	injector := do.New()
	connStr := "user=root dbname=mydb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	do.ProvideNamed[*sql.DB](injector, "user", func(i *do.Injector) (*sql.DB, error) {
		return db, nil
	})

	do.Provide(injector, user.NewRepository)
	do.Provide(injector, user.NewService)
	do.Provide(injector, user.NewHandler)

	userHandler := do.MustInvoke[domain.UserHandler](injector)
	http.Handle("/user", userHandler.FetchByUsername())
	fmt.Printf("Try run server at :%d\n", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
