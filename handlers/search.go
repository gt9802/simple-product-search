package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type SearchHandler struct {
	Products []string
}

func NewSearchHandler(products []string) *SearchHandler {
	return &SearchHandler{Products: products}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing Query Parameter", http.StatusBadRequest)
	}
	query = strings.ToLower(query)
	var results []string
	for _, product := range h.Products {
		if strings.HasPrefix(strings.ToLower(product), query) {
			results = append(results, product)
		}
		if len(results) >= 5 {
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}
