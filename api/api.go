package api

import (
	"encoding/json"
	"net/http"

	socketio "github.com/googollee/go-socket.io"

	"github.com/gorilla/mux"
)

type API struct {
	cors string
	*mux.Router
	socket     *Socket
	Tournament *Tournament
}

func NewAPI(subrouter *mux.Router, server *socketio.Server) *API {
	tournament := NewTournament("548741")
	socket := NewSocket(server, tournament)
	api := &API{
		"",
		subrouter,
		socket,
		tournament,
	}
	api.HandleFunc("/info", api.handleGetInfo()).Methods(http.MethodGet)
	api.HandleFunc("/pairs", api.handleGetPairs()).Methods(http.MethodGet)
	api.HandleFunc("/competitors", api.handleGetCompetitors()).Methods(http.MethodGet)
	api.HandleFunc("/round", api.handleSetRound()).Methods(http.MethodPost)
	api.HandleFunc("/tournament", api.handleSetTournament()).Methods(http.MethodPost)
	api.HandleFunc("/result", api.handleSetResults()).Methods(http.MethodPost)
	return api
}

func (api *API) handleGetPairs() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		pairs := api.Tournament.GetPairs()
		res, _ := json.Marshal(pairs.Items)
		resp.Write(res)
	}
}

func (api *API) handleGetCompetitors() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		competitors := api.Tournament.GetCompetitors()
		res, _ := json.Marshal(competitors.Items)
		resp.Write(res)
	}
}

func (api *API) handleGetInfo() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		
		result := struct {
			Round      string
			Tournament string
		}{
			api.Tournament.round,
			api.Tournament.id,
		}
		res, _ := json.Marshal(result)
		resp.Write(res)
	}
}

func (api *API) handleSetRound() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Round string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)

		api.Tournament.SetRound(post.Round)
		api.socket.SendNewRound()
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
		api.socket.SendStartList()
		resp.WriteHeader(http.StatusOK)
	}
}

func (api *API) handleSetResults() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post struct {
			Round string
		}
		_ = json.NewDecoder(req.Body).Decode(&post)
		api.Tournament.SetRound(post.Round)
		api.socket.SendResults()
		resp.WriteHeader(http.StatusOK)
	}
}
