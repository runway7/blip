package blip

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	latLon := r.Header.Get("X-AppEngine-CityLatLong")
	parts := strings.Split(latLon, ",")
	lat := strings.TrimSpace(parts[0])
	lon := strings.TrimSpace(parts[1])

	m := map[string]string{
		"latLong": latLon,
		"city":    r.Header.Get("X-AppEngine-City"),
		"region":  r.Header.Get("X-AppEngine-Region"),
		"country": r.Header.Get("X-AppEngine-Country"),
		"lat":     lat,
		"lon":     lon,
	}

	js, err := json.Marshal(m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(js))

}
