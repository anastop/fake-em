package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Metrics struct {
	Metric1 float64
	Metric2 int64
	Metric3 string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := &Metrics{rand.Float64(), rand.Int63n(100), strconv.FormatFloat(rand.Float64(), 'f', 2, 32)}
	res, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("failed to marshal %v: %v\n", m, err)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(res); err != nil {
		log.Fatalf("failed to write data %+v: %v", res, err)
	}
}

func main() {
	port := flag.String("port", "9999", "server port")
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/data", Handler)
	srv := &http.Server{Addr: ":" + *port, Handler: mux}
	srv.ListenAndServe()
}
