package cities

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/patrickmn/go-cache"

	"github.com/julienschmidt/httprouter"
)

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

//Api is endpoint for cities
func Api(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
