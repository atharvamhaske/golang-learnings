package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func headerMiddleware( next http.Handler ) http.Handler {
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {
		//implement logic here

		// auth here ! X-API-key
		w.Header().Set("X-Custom-Header", "Pokemon")

		// end of Middleware
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc( func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		// Calculate duration and log with nanosecond precision
		duration := time.Since(start)
		log.Printf("%s %s - Duration: %v (%d ns)", 
			r.Method, 
			r.RequestURI, 
			duration, 
			duration.Nanoseconds())
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to home page")
}

func aboutHandler ( w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to the about page")
}
func main() {
	mux := http.NewServeMux()

	mux.Handle("/", loggingMiddleware(headerMiddleware(http.HandlerFunc(homeHandler))))
	mux.Handle("/about", loggingMiddleware(headerMiddleware(http.HandlerFunc(aboutHandler))))

    log.Println("Starting port on 8080....")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		// the logic that should be executed in case the listen and serve returns error
		log.Fatal("Server Failed", err)
	}
}