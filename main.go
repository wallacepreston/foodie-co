package main
	
import (
	"database/sql"
	"fmt"
	"net/http"
	"encoding/json"
	_ "github.com/lib/pq"
)

type Recipe struct {
	ID          int    `json:"recipe_id"`
	Name        string `json:"name"`
	Instructions string `json:"instructions"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

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
		// Query the database for all recipes
		rows, err := db.Query("SELECT * FROM recipes")
		if err != nil {
				panic(err)
		}
    defer rows.Close()

    // Create a slice to store the results
    var recipes []Recipe

    // Iterate over the rows and scan into Recipe structs
    for rows.Next() {
			var recipe Recipe
			err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Instructions, &recipe.CreatedAt, &recipe.UpdatedAt)
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}
			recipes = append(recipes, recipe)
    }

    // Convert the results to JSON and send the response
    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(recipes)
    if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
    }
	})

	// Start the server
	fmt.Println("Server is listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
