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
	"path/filepath"
	"strings"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/aeberzin/chess-results-viewer/api"
)

const urlScheme = "http"

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

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
	// vueHandler := http.FileServer(Vue("web/dist/"))
	apiHandler := api.NewAPI(router.PathPrefix("/api").Subrouter(), server)

	spa := spaHandler{staticPath: "web/dist", indexPath: "/index.html"}
	router.PathPrefix("/").Handler(spa)
	router.Handle("/api/", apiHandler)
	router.Handle("/socket.io/", corsMiddleware(server))
	// router.Handle("/", vueHandler)
	// router.NotFoundHandler = spa
	// http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//   http.ServeFile(w, r, "web/dist/index.html")
	// })
	// router.PathPrefix("/dist").Handler(http.FileServer(http.Dir("web/dist/")))

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

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	log.Println("Listening on", port, "at", config.URL)
	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, handler)))
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
		URL:  "http://localhost:" + port,
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
	fmt.Println(name)
	return http.Dir(v).Open(name)
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
