package main
	
import (
		"database/sql"
		"fmt"
		"net/http"
	_ "github.com/lib/pq"
)

func main() {

	// Database connection setup
	db, err := sql.Open("postgres", "user=postgres dbname=foodie sslmode=disable")
	if err != nil {
			panic(err)
	}
	defer db.Close()
	
	// Check the connection
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
			panic(err)
	}
	fmt.Println("PostgreSQL version:", version)

	// Server setup
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Here are the recipes")
	})

	// Start the server
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
