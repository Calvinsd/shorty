package main

import (
	"fmt"
	"net/http"

	"github.com/Calvinsd/shorty"
)

func main() {

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/PhUWOmnPhw4": "https://www.youtube.com/watch?v=PhUWOmnPhw4",
		"/Ek2VkdkOenw": "https://www.youtube.com/watch?v=Ek2VkdkOenw",
	}

	mapHandler := shorty.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", mapHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to shorty")
}
