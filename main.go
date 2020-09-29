package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/aeberzin/chess-results-viewer/api"
)

const urlScheme = "http"

func main() {
	config := loadConfig()
	router := mux.NewRouter()

	// Endpoints for the API and Vue client
	vueHandler := http.FileServer(Vue("web/dist/"))
	apiHandler := api.NewAPI(router.PathPrefix("/api").Subrouter())

	router.Handle("/", vueHandler)
	router.Handle("/api/", apiHandler)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			switch arr := strings.Split(scanner.Text(), " "); arr[0] {
			case "r":
				apiHandler.Tournament.SetRound(arr[1])
			case "t":
				apiHandler.Tournament.SetID(arr[1])
			}
		}
	}()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{config.URL, config.DevURL},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("Listening on", config.Port, "at", config.URL)
	log.Fatal(http.ListenAndServe(config.Port, handlers.LoggingHandler(os.Stdout, handler)))
}

type Config struct {
	Name   string
	Port   string
	URL    string
	DevURL string
}

func loadConfig() *Config {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [input]\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nInput:\n  Files or directories, leave blank to use stdin\n")
	}
	configFilename := flag.String("c", "config.json", "Specify the configuration file.")
	flag.Parse()

	configFile, err := os.Open(*configFilename)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	config := &Config{
		Port: ":3000",
		URL:  "http://localhost:8080",
	}
	if err := json.NewDecoder(configFile).Decode(config); err != nil {
		panic("parsing config: " + err.Error())
	}

	return config
}

type Vue string

func (v Vue) Open(name string) (http.File, error) {
	if ext := path.Ext(name); name != "/" && (ext == "" || ext == ".html") {
		name = "index.html"
	}
	return http.Dir(v).Open(name)
}
