package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	runMigrations()
	urlExample := "postgres://postgres:postgres@localhost:5432/filerain"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Execute a query from
	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select 'Name' as name, 2 as weight;").Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)

	http.HandleFunc("GET /", signUpHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("GET /static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("GET /auth/signup", signUpHandler)
	http.HandleFunc("POST /auth/signup", signUpHandlerPost)

	log.Println("Serving at: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
