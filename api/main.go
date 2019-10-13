package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Healthy!")
	})

	fmt.Println("Listening on :8066")
	log.Fatal(http.ListenAndServe(":8066", nil))
}
