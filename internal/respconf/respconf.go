package respconf

import (
	"encoding/json"
	"os"
)

type Config struct {
	Path       string            `json:"path"`
	StatusCode int               `json:"status_code"`
	Body       interface{}       `json:"body"`
	Headers    map[string]string `json:"headers"`
	Method     string            `json:"method"`
}

func ReadConfig(fpath string) ([]Config, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var conf []Config
	err = dec.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
