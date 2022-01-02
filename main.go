package main

import (
	"encoding/json"
	"github.com/hid8/logger"
	"net/http"
)

const (
	VERSION = "1.0.0"
	PROJECT_KEY = "hid8.ntf"
)

var log = logger.NewChannel("main")

type Response struct {
	Method string `json:"method"`
	Host string `json:"host"`
	Path string `json:"path"`
	Headers map[string][]string `json:"headers"`
}

type Handler struct {}

func (h Handler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Method:  r.Method,
		Host:    r.Host,
		Path:    r.URL.Path,
		Headers: r.Header,
	}

	bytes, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {
	log.InfoF("Starting %s v%s", PROJECT_KEY, VERSION)

	log.Info("Starting server...")
	http.ListenAndServe(":8080", Handler{})
}