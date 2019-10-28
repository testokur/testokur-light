package localization

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

//Get returns license types
func Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	const cacheKey = "localization"
	cities, found := memoryCache.Get(cacheKey)

	if found {
		utils.WriteOKResponse(w, cities)
		return
	}

	dat, err := ioutil.ReadFile(filepath.Join("data", "local-strings.json"))
	if err != nil {
		panic(err)
	}
	json.Unmarshal(dat, &cities)
	memoryCache.Set(cacheKey, cities, cache.NoExpiration)
	utils.WriteOKResponse(w, cities)
}
