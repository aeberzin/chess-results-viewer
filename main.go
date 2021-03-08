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

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/aeberzin/chess-results-viewer/api"
)

const urlScheme = "http"

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With"

		// w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "false")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}

func main() {
	config := loadConfig()
	router := mux.NewRouter()

	server, _ := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Join("all")
		fmt.Println("connected:", s.ID())
		return nil
	})
	go server.Serve()
	defer server.Close()

	// Endpoints for the API and Vue client
	vueHandler := http.FileServer(Vue("web/dist/"))
	apiHandler := api.NewAPI(router.PathPrefix("/api").Subrouter(), server)

	router.Handle("/", vueHandler)
	router.Handle("/api/", apiHandler)
	router.Handle("/socket.io/", corsMiddleware(server))

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
		AllowedOrigins:   []string{"*"},
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

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	config := &Config{
		Port: port,
		URL:  "http://localhost:"+port,
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
