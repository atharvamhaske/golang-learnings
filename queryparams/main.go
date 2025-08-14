package main

import (
	"fmt"
	"net/http"
	"strings"
)

// http://api.example.com/api/v1/greet?name=Atharva

//understanding query params
func greetHandler( w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"

	}
	fmt.Fprintf(w, "Hello, %s!", name)
}


//extracting path variables
// https://api.example.com/user/123
// 1 -> User
// 2 -> UserID

func userHandler ( w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) >= 3 && pathSegments[1] == "user" {
		userID := pathSegments[2]
		name := pathSegments[3]
		fmt.Fprintf(w, "User ID: %s\n", userID )
		fmt.Fprintf(w, "Name is : %s", name )
	} else {
		http.NotFound(w, r)
	}
}

//handling both
// https://api.example.com/username/123?includeDetails=true

func userDetailsHandler ( w http.ResponseWriter, r *http.Request ) {
	pathSegments := strings.Split(r.URL.Path, "/")
	query := r.URL.Query()
	includeDetails := query.Get("includeDetails")

	if len(pathSegments) >= 3 && pathSegments[1] == "username" {
		userID := pathSegments[2]
		response := fmt.Sprintf("User ID: %s", userID)

		if includeDetails == "true" {
			response += " (Details included )"
		}
		fmt.Fprintln(w, response)
	} else {
		http.NotFound(w, r)
	}
}


func main() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/username/", userDetailsHandler)

	fmt.Println("Listening at port 8080 ....")

	// we are using default mux for demonstration

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("failed to listen at 8080", err)
	}
}

//https://api.example.com/api/v1/notion?sessionToken="5626356"&referralToken="user123" -> user123 was the one who reffered this guy to our app -> so we can increase his points in our db or reward him some amount
// https://api.example.com/api/v1/notion?sessionToken="5626356" -> user came here and signed up by his own and no ones ref is used here
