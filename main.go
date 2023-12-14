package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	// Register the handler function for the root ("/") route
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
