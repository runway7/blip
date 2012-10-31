package blip

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	m := map[string]string{
		"latLong": r.Header.Get("X-AppEngine-CityLatLong"),
		"city":    r.Header.Get("X-AppEngine-City"),
		"region":  r.Header.Get("X-AppEngine-Region"),
		"country": r.Header.Get("X-AppEngine-Country"),
	}

	js, err := json.Marshal(m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(js))

}
