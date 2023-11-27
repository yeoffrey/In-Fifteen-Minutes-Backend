package main

import (
	"fmt"
	"net/http"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Handle your POST request logic here
	// This is where you can process the data received from the Vue frontend

	// Example: Log the received data
	fmt.Println("Received POST request")

	// Send a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request processed successfully"))
}

func main() {
	// Define the endpoint to handle
	http.HandleFunc("/your-endpoint", handleRequests)

	// Start the server
	port := "8080" // You can set the port according to Google Cloud Function requirements
	fmt.Printf("Server listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
