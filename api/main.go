package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/testokur/testokur-light/api/config"
)

func main() {
	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Healthy!")
	})
	http.Handle("/metrics", promhttp.Handler())

	log.Println(fmt.Sprintf("Listening on %s...", config.GetPort()))
	log.Fatal(http.ListenAndServe(config.GetPort(), nil))
}
