package handlers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

type MapItem struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type MapItemJson struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, exists := pathsToUrls[r.URL.Path]; exists {
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(yamlRead []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var list []MapItem
	if err := yaml.Unmarshal(yamlRead, &list); err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for _, item := range list {
		m[item.Path] = item.Url
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if url, exists := m[r.URL.Path]; exists {
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}, nil
}

func JSONHandler(jsonRead []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var list []MapItemJson
	if err := json.Unmarshal(jsonRead, &list); err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for _, item := range list {
		m[item.Path] = item.Url
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if url, exists := m[r.URL.Path]; exists {
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}, nil
}
