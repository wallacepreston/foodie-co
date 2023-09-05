package main
	
import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"foodie-co/database"
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
	dsn := "user=postgres dbname=foodie sslmode=disable"
	db, err := database.InitDatabase(dsn)
	if err != nil {
			log.Fatal(err)
	}
	// log out db
	fmt.Println(db)

	// Server setup
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
    var recipes []database.Recipe // Use the Recipe struct from the database package

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
