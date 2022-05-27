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
	CategoryID  int
	Category    Category `gorm:"foreignKey:CategoryID"`
}

// A Category describes a group of Products.
type Category struct {
	ID          int
	Name        string
	Description string
}

// getProducts is the HTTP handler for GET /products.
func Handler(w http.ResponseWriter, r *http.Request) {
	// Connect to PlanetScale database using DSN environment variable.
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}

	var products []Product
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var product Product
		products = append(products, product)
	}
	// if err = rows.Err(); err != nil {
	// 	return products, err
	// }
	// return products, nil
	js, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}
