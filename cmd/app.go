package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/AntonPates/mock-http-server/internal/respconf"
)

const (
	generalExitCode = 1
	successExitCode = 0
)

func app() int {
	confFilePath := flag.String("config", "config.json", "path to config file")
	addr := flag.String("addr", ":8080", "address to listen")
	flag.Parse()

	conf, err := respconf.ReadConfig(*confFilePath)
	if err != nil {
		return generalExitCode
	}
	router := http.NewServeMux()
	for _, v := range conf {
		fmt.Println("status code:", v.StatusCode)
		router.HandleFunc(v.Path, CreateHttpHandlerFunc(v))
	}
	err = http.ListenAndServe(*addr, router)
	if err != nil {
		return generalExitCode
	}
	return successExitCode
}

func CreateHttpHandlerFunc(config respconf.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for key, value := range config.Headers {
			w.Header().Set(key, value)
		}
		if body, ok := config.Body.(string); ok {
			w.WriteHeader(config.StatusCode)
			w.Write([]byte(body))
		} else {
			w.WriteHeader(config.StatusCode)
			json.NewEncoder(w).Encode(config.Body)
		}
	}
}
