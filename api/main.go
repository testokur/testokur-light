package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/testokur/testokur-light/api/config"
)

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Healthy!")
}

func prometheusMetrics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	promhttp.Handler().ServeHTTP(w, r)
}

type district struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type city struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Districts []district `json:"districts"`
}

func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

var memoryCache = cache.New(99999999, 99999999)

func cities(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	const cacheKey = "cities"
	cities, found := memoryCache.Get(cacheKey)

	if found {
		writeOKResponse(w, cities)
		return
	}

	dat, err := ioutil.ReadFile(filepath.Join("data", "cities.json"))
	if err != nil {
		panic(err)
	}
	json.Unmarshal(dat, &cities)
	memoryCache.Set(cacheKey, cities, cache.NoExpiration)
	writeOKResponse(w, cities)
}

func main() {
	router := httprouter.New()
	router.GET("/hc", healthCheck)
	router.GET("/cities", cities)
	router.GET("/metrics", prometheusMetrics)

	log.Println(fmt.Sprintf("Listening on %s...", config.GetPort()))
	log.Fatal(http.ListenAndServe(config.GetPort(), router))
}
