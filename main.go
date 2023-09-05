package main

import (
	"encoding/json"
	"fmt"
	"foodie-co/database"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Recipe struct {
	ID           int    `json:"recipe_id"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func main() {
	// Database connection
	dsn := "user=postgres dbname=foodie sslmode=disable"
	db, err := database.InitDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}
	// log out db
	fmt.Println(db)

	// Server

	// GET /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// GET /recipes
	http.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		var recipes []database.Recipe

		// Query the "recipes" table using GORM
		database.DB.Find(&recipes)

		// Convert the results to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(recipes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the server
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
