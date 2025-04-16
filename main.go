package main

import (
	"log"
	"net/http"

	"github.com/gt9802/simple-product-search/handlers"
)

func main() {
	products := []string{
		"Phone",
		"Photo Frame",
		"Laptop",
		"Headphones",
		"Tablet",
		"Charger",
		"Camera",
		"Speaker",
		"Monitor",
		"Keyboard",
	}

	handler := handlers.NewSearchHandler(products)
	http.HandleFunc("/products/search/", func(w http.ResponseWriter, r *http.Request) {
		handler.Search(w, r)
	})

	log.Println("Server running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
