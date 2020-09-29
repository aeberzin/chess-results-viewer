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
	"time"

	"github.com/aeberzin/chess-results-viewer/api"
)

const urlScheme = "http"

func main() {
	config := loadConfig()

	s := &http.Server{
		Addr:           config.Port,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	// Endpoints for the API and Vue client
	vueHandler := http.FileServer(Vue("web/dist/"))
	apiHandler := api.NewAPI()

	http.Handle("/api/", apiHandler)
	http.Handle("/", vueHandler)

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

	log.Println("Listening on", config.Port, "at", config.URL)
	log.Fatal(s.ListenAndServe())
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
