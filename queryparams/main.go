package main

import (
	"fmt"
	"net/http"
)

// http://api.example.com/api/v1/greet?name=Atharva
func greetHandler( w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"

	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/greet", greetHandler)

	fmt.Println("Listening at port 8080 ....")

	// we are using default mux for demonstration

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("failed to listen at 8080", err)
	}
}

//https://api.example.com/api/v1/notion?sessionToken="5626356"&referralToken="user123" -> user123 was the one who reffered this guy to our app -> so we can increase his points in our db or reward him some amount
// https://api.example.com/api/v1/notion?sessionToken="5626356" -> user came here and signed up by his own and no ones ref is used here
