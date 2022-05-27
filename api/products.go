// Command golang-example demonstrates how to connect to PlanetScale from a Go
// application.
package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// A Product contains metadata about a product for sale.
type Product struct {
	ID          int
	Name        string
	Description string
	Image       string
}

// getProducts is the HTTP handler for GET /products.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Connect to PlanetScale database using DSN environment variable.
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}

	var products []Product
	rows, err := db.Query("SELECT ID, Name, Description, Image FROM products")
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)
}
