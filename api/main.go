package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/testokur/testokur-light/api/config"
)

func main() {
	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Healthy!")
	})

	log.Println(fmt.Sprintf("Listening on %s...", config.GetPort()))
	log.Fatal(http.ListenAndServe(config.GetPort(), nil))
}
