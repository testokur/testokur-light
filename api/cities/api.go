package cities

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/patrickmn/go-cache"
	"github.com/testokur/testokur-light/utils"
)

var memoryCache = cache.New(99999999, 99999999)

//Get returns cities from cities.json file
func Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	const cacheKey = "cities"
	cities, found := memoryCache.Get(cacheKey)

	if found {
		utils.WriteOKResponse(w, cities)
		return
	}

	dat, err := ioutil.ReadFile(filepath.Join("data", "cities.json"))
	if err != nil {
		panic(err)
	}
	json.Unmarshal(dat, &cities)
	memoryCache.Set(cacheKey, cities, cache.NoExpiration)
	utils.WriteOKResponse(w, cities)
}
