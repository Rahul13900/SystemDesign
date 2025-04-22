package main

import (
	"fmt"
	"net/http"
)

// HelloHandler is the basic HTTP handler that responds with a message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, user!")
}

// LoggingMiddleware is a decorator that logs each request before it is processed.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and path
		fmt.Println("[LOG] Incoming request:", r.Method, r.URL.Path)

		// Call the actual handler (next in the chain)
		next(w, r)
	}
}

// AuthMiddleware is a decorator that checks for authentication before processing the request.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the Authorization header from the request
		token := r.Header.Get("Authorization")

		// Check if the token matches the expected value
		if token != "Bearer secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authorized, call the actual handler
		next(w, r)
	}
}

var requestCount = 0

// RateLimitMiddleware is a decorator that limits the number of requests per minute.
func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Increment the request count (simplified rate-limiting logic)
		requestCount++

		// If more than 5 requests are made, reject further requests
		if requestCount > 5 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Call the actual handler if within the limit
		next(w, r)
	}
}

func main() {
	// Compose the handler by stacking decorators
	finalHandler := LoggingMiddleware(AuthMiddleware(RateLimitMiddleware(HelloHandler)))

	// Set up the server to handle requests at /hello
	http.HandleFunc("/hello", finalHandler)

	// Start the server on port 8080
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
