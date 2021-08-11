package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type placeHandler struct {
	places *places
}

func (h *placeHandler) handleGetPlace(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	cord := q.Get("cord")
	cords := strings.Split(cord, ",")
	if len(cords) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing lat/long"))
		return
	}
	latStr := cords[0]
	longStr := cords[1]
	lat, err1 := strconv.ParseFloat(latStr, 64)
	long, err2 := strconv.ParseFloat(longStr, 64)
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Missing lat/long, %s %s", err1, err2)))
		return
	}
	place := h.places.nearestPlace(lat, long)
	res, err := json.Marshal(place)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *placeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet:
		h.handleGetPlace(w, r)
	default:
		notFound(w, r)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("Can not %s %s ", r.Method, r.RequestURI)))
}

func CreateServer(address string, places *places) {
	mux := http.NewServeMux()

	placeHandler := &placeHandler{
		places: places,
	}
	mux.Handle("/api/place", placeHandler)
	mux.Handle("/api/place/", placeHandler)
	fmt.Printf("Listening on %s\n", address)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatalf("Failed listening on %s, %s", address, err)
	}
}
