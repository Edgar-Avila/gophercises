package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
    "url-shortener/handlers"
)


func main() {
    // Flags
    yamlFlag := flag.String("yaml", "redirects.yaml", "Path to a yaml file to read redirects from")
    jsonFlag := flag.String("json", "redirects.json", "Path to a json file to read redirects from")
    flag.Parse()
    pathToYaml := *yamlFlag
    pathToJson := *jsonFlag

    // Start message
	fmt.Println("Started server at port 8080")

    // Default handler
	server := mux.NewRouter()
    server.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello")
    })

    // Map of paths and their long versions
	paths := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

    // Read yaml file
    textYaml, err := os.ReadFile(pathToYaml)
    if err != nil {
        log.Fatal(err)
    }

    // Read json file
    textJson, err := os.ReadFile(pathToJson)
    if err != nil {
        log.Fatal(err)
    }

    // Create map handler, yaml and handler
	mapHandler := handlers.MapHandler(paths, server)
	yamlHandler, err := handlers.YAMLHandler(textYaml, mapHandler)
	if err != nil {
		log.Fatal(err)
	}
    jsonHandler, err := handlers.JSONHandler(textJson, yamlHandler)
	if err != nil {
		log.Fatal(err)
	}

    // Serve app
	if err := http.ListenAndServe(":8080", jsonHandler); err != nil {
		log.Fatal("Could not initialize server")
	}
}
