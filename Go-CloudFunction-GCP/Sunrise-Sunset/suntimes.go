package sundemo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/schema"
	"github.com/nathan-osman/go-sunrise"
)

/*Request is the request struct*/
type Request struct {
	Lat  float64   `schema:"lat,required"`
	Lon  float64   `schema:"lon,required"`
	Date time.Time `schema:"date,required"`
}

/*Response is the response struct*/
type Response struct {
	Sunset  time.Time `json:"sunset"`
	Sunrise time.Time `json:"sunrise"`
}

/*SunTimes is the function for tracking sunrise and sunset time by location and date.*/
func SunTimes(w http.ResponseWriter, r *http.Request) {
	var decoder = schema.NewDecoder()
	decoder.RegisterConverter(time.Time{}, dateConverter)

	// Parse the request from query string
	var req Request
	if err := decoder.Decode(&req, r.URL.Query()); err != nil {
		// Report any parsing errors
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	// Perform sunrise/sunset calculation
	sunrise, sunset := sunrise.SunriseSunset(
		req.Lat, req.Lon,
		req.Date.Year(), req.Date.Month(), req.Date.Day(),
	)

	// Send response back to client as JSON
	w.WriteHeader(http.StatusOK)
	response := Response{sunset, sunrise}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		panic(err)
	}
}

func dateConverter(value string) reflect.Value {
	s, _ := time.Parse("2006-01-_2", value)
	return reflect.ValueOf(s)
}
