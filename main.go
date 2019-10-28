package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/testokur/testokur-light/api/cities"
	"github.com/testokur/testokur-light/api/licensetypes"
	"github.com/testokur/testokur-light/api/localization"
	"github.com/testokur/testokur-light/config"
)

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Healthy!")
}

func prometheusMetrics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	router := httprouter.New()
	router.GET("/hc", healthCheck)
	router.GET("/api/v1/cities", cities.Get)
	router.GET("/api/v1/license-types", licensetypes.Get)
	router.GET("/api/v1/localization", localization.Get)
	router.GET("/metrics", prometheusMetrics)

	log.Println(fmt.Sprintf("Listening on %s...", config.GetPort()))
	log.Fatal(http.ListenAndServe(config.GetPort(), router))
}
