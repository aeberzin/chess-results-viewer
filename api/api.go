package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

type API struct {
	cors string
	*mux.Router
	Tournament *Tournament
}

func NewAPI(subrouter *mux.Router) *API {
	router := subrouter
	tournament := NewTournament("535239")
	api := &API{
		"",
		router,
		tournament,
	}
	api.HandleFunc("/pairs", api.handleGetPairs()).Methods(http.MethodGet)
	api.HandleFunc("/competitors", api.handleGetCompetitors()).Methods(http.MethodGet)
	api.HandleFunc("/round", api.handleSetRound()).Methods(http.MethodPost)
	api.HandleFunc("/tournament", api.handleSetTournament()).Methods(http.MethodPost)
	return api
}

func (api *API) handleGetPairs() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		res, err := http.Get("http://chess-results.com/tnr" + api.Tournament.id + ".aspx?lan=11&art=2&rd=" + api.Tournament.round)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".CRs1").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			pairs := NewPairs(s.Find("tr").First())
			s.Find("tr").Each(func(i int, tr *goquery.Selection) {
				if i != 0 {
					pairs.AddPair(tr)
				}
			})
			res, _ := json.Marshal(pairs.Pairs)
			resp.Write(res)
		})
	}
}

func (api *API) handleGetCompetitors() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		res, err := http.Get("http://chess-results.com/tnr" + api.Tournament.id + ".aspx?lan=11&art=1&rd=" + api.Tournament.round)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find(".CRs1").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			competitors := NewCompetitors(s.Find("tr").First())
			lastResult := ""
			s.Find("tr").Each(func(i int, tr *goquery.Selection) {
				if i != 0 {
					competitors.AddCompetitor(tr)
					competitor := &competitors.Competitors[len(competitors.Competitors)-1]
					if competitor.Finish != "" {
						lastResult = competitor.Finish
					} else {
						competitor.Finish = lastResult
					}
				}
			})
			res, _ := json.Marshal(competitors.Competitors)
			resp.Write(res)
		})
	}
}

func (api *API) handleSetRound() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Round string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)

		api.Tournament.SetRound(post.Round)
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleSetTournament() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Tournament string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)

		api.Tournament.SetID(post.Tournament)
		resp.WriteHeader(http.StatusOK)
	}
}
