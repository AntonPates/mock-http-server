package main

import (
	"encoding/json"
	"flag"
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
	m := make(map[string]map[string]respconf.Config)
	for _, v := range conf {
		if _, ok := m[v.Path]; !ok {
			m[v.Path] = make(map[string]respconf.Config)
		}
		m[v.Path][v.Method] = v
	}
	router := http.NewServeMux()
	for path, v := range m {
		router.HandleFunc(path, CreateHttpHandlerFunc(v))
	}
	err = http.ListenAndServe(*addr, router)
	if err != nil {
		return generalExitCode
	}
	return successExitCode
}

func CreateHttpHandlerFunc(config map[string]respconf.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, v := range config {
			if v.Method != r.Method {
				continue
			}
			for key, value := range v.Headers {
				w.Header().Set(key, value)
			}
			if body, ok := v.Body.(string); ok {
				w.WriteHeader(v.StatusCode)
				w.Write([]byte(body))
			} else {
				w.WriteHeader(v.StatusCode)
				json.NewEncoder(w).Encode(v.Body)
			}
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
