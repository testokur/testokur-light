package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/testokur/testokur-light/api/config"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Healthy!")
}

func Metrics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	promhttp.Handler().ServeHttp(w, r)
}

func main() {
	router := httprouter.New()
	router.GET("/hc", HealthCheck)
	router.GET("/metrics", Metrics)

	log.Println(fmt.Sprintf("Listening on %s...", config.GetPort()))
	log.Fatal(http.ListenAndServe(config.GetPort(), router))
}
