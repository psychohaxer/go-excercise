package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// SortBooksRequest represents the JSON request body for sorting books
type SortBooksRequest struct {
	Input string `json:"input"`
}

// SortBooksResponse represents the JSON response body for sorted books
type SortBooksResponse struct {
	SortedBooks string `json:"sorted_books"`
}

func sortBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request
	var req SortBooksRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Sort the books
	sortedBooks := SortTheBooks(req.Input)

	// Create JSON response
	res := SortBooksResponse{
		SortedBooks: sortedBooks,
	}

	// Convert response to JSON
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	w.Write(jsonRes)
}

func main() {
	// Define the HTTP route
	http.HandleFunc("/sort-books", sortBooksHandler)

	// Start the web server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
