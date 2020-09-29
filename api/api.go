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
	api.HandleFunc("/pairs", api.GetPairs())
	return api
}

func (api *API) GetPairs() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.Println("get pairs")
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
